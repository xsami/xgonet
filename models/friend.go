package models

import (
	"log"

	logger "github.com/xsami/xgonet/logger"
)

// Friend struct contains the friend relationship
type Friend struct {
	ID         int  `json:"id"`
	UserIDFrom int  `json:"from_id"`
	UserIDTo   int  `json:"to_id"`
	Accepted   bool `json:"accepted"`
}

type RelateFriend struct {
	UserIDA int
	UserIDB int
}

func BuildFriendMap(friends []Friend) (resultMap map[RelateFriend]bool) {

	acceptedFriend := FilterFriends(friends, func(f Friend) bool {
		return true
	})
	resultMap = make(map[RelateFriend]bool, len(acceptedFriend))

	for _, value := range acceptedFriend {
		rFriend := NewRelatedFriend(value.UserIDFrom, value.UserIDTo)
		resultMap[rFriend] = value.Accepted
	}
	return resultMap
}

func NewRelatedFriend(u1, u2 int) RelateFriend {
	uA, uB := u1, u2
	if u1 < u2 {
		uA, uB = u2, u1
	}
	return RelateFriend{
		UserIDA: uA,
		UserIDB: uB}
}

// FindUserFriends return a slice of type User with the friends info
func FindUserFriends(friends []Friend, user User) []User {
	var (
		resultFriends []Friend
		resultUsers   []User
		friendID      int
		userID        = user.ID
	)
	resultFriends = FilterFriends(friends, func(f Friend) bool {
		return (f.UserIDFrom == userID || f.UserIDTo == userID)
	})

	resultUsers = []User{}
	for _, friend := range resultFriends {

		// Check the ID whose different to the current user.ID
		if friend.UserIDFrom != userID {
			friendID = friend.UserIDFrom
		} else {
			friendID = friend.UserIDTo
		}

		friendUser, err := FindUserByID(UserList, friendID)
		if err != nil {
			log.Fatal("FindUserFriends failed. ", err) // As couldn't find the user, or the user isn't active
		} else {
			resultUsers = append(resultUsers, friendUser)
		}
	}

	return resultUsers
}

// ValidateFriendShip return a boolean (true) value whenever 2 users are friends
func ValidateFriendShip(friends []Friend, userIDA int, userIDB int) bool {

	resultFriends := FilterFriends(friends, func(f Friend) bool {
		return ((f.UserIDFrom == userIDA && f.UserIDTo == userIDB) ||
			(f.UserIDFrom == userIDB && f.UserIDTo == userIDA))
	})

	if len(resultFriends) > 0 {
		return true
	}

	return false
}

// FindTwoUserRelationShip check how can two users be releated
//
// counter parameter must always start passing 0 by parameter
func FindTwoUserRelationShip(friends map[RelateFriend]bool, userA User, userB User, counter int, treshold uint) (resultArray []User, resultCounter int) {

	if counter == 0 {
		logger.Log("FindTwoUserRelationShip", FriendMap)
	}

	if counter >= int(treshold) && treshold != 0 {
		return []User{}, -1
	}

	// TODO: the function should receibe a list with the friends that should be evaluated
	// that are friend of a x person.

	// TODO: Make a for loop to iterate over the list of friend
	// and add this block as a functio
	// TODO: return empty if the list of comparation has been already comparated
	// and in every comparation check that you're not repeating yourself
	// if you start repeating yourself then. there's no way userA is friend of userB
	tmpRS := NewRelatedFriend(userA.ID, userB.ID)

	if FriendMap[tmpRS] {
		if resultArray != nil {
			resultArray = append(resultArray, userA)
		} else {
			resultArray = []User{userA}
		}
		return resultArray, counter
	}
	// TODO: end of the for loop

	// TODO: Call a function which give you the ID's of the friends of the friends
	// in the already passed function so you should call the function again
	// This must be unique friends and not evaluated comparision already
	// for k1, v1 := range FriendMap {

	// }

	return resultArray, 0
}

// FilterFriends return a slice of type Friend.
// Given a slice of type Friend, this function evaluate the condition
// that is passed as parameter and return a new slice of Friends
// where the condition meet the evaluator and if the friend request was accepted
func FilterFriends(friends []Friend, evaluator func(f Friend) bool) []Friend {

	resultArray := []Friend{}
	for _, friend := range friends {
		if evaluator(friend) && friend.Accepted {
			resultArray = append(resultArray, friend)
		}
	}

	return resultArray
}
