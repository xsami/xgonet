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

// RelateFriend is a struct used to relate user friendship
type RelateFriend struct {
	UserIDA int
	UserIDB int
}

// BuildFriendMap receive a list of type Friend
// then it return a map of type RelateFriend
// with the unique friend relationships
func BuildFriendMap(friends []Friend) (resultMap map[RelateFriend]bool) {

	acceptedFriend := FilterFriends(friends, func(f Friend) bool { // Filter only by users that are active
		return true
	})

	resultMap = make(map[RelateFriend]bool, len(acceptedFriend))

	for _, value := range acceptedFriend {
		rFriend := NewRelatedFriend(value.UserIDFrom, value.UserIDTo)

		resultMap[rFriend] = value.Accepted
	}

	return resultMap
}

// NewRelatedFriend creates a new RelateFriend relationship
// and this must be used to create a RelateFriend object
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

// GetFriendsID return a list of IDs
// with the friends of a given userID
func GetFriendsID(id int) []int {

	result := []int{}
	for _, v := range FriendList {
		if v.Accepted {
			if v.UserIDFrom == id {
				result = append(result, v.UserIDTo)
			} else if v.UserIDTo == id {
				result = append(result, v.UserIDFrom)
			}
		}

	}

	return result
}

// FindTwoUserRelationShip check how can two users be releated
//
// counter parameter must always start passing 0 by parameter
func FindTwoUserRelationShip(friends map[RelateFriend]bool, evaluatedFriendShipt map[RelateFriend]bool, user User, userIDList []int, counter int, treshold uint) (resultArray []User, resultCounter int) {

	if counter == 0 {
		logger.Log("FindTwoUserRelationShip", FriendMap)
	}

	if counter > int(treshold) && treshold != 0 {
		return []User{}, -1
	}

	flagRepetitions := false
	possibleFriends := []int{}
	newPossibleFriends := []int{}

	for _, uid := range userIDList {

		tmpRS := NewRelatedFriend(user.ID, uid)

		if !evaluatedFriendShipt[tmpRS] {

			possibleFriends = append(possibleFriends, uid)
			evaluatedFriendShipt[tmpRS] = true
			flagRepetitions = true

			if FriendMap[tmpRS] {
				user, err := FindUserByID(UserList, uid)
				if err != nil {
					logger.Log("FindTwoUserRelationShip", err)
				}
				return []User{user}, counter
			}
		}
	}

	// Return empty if the list of comparation has already been compared
	if !flagRepetitions {
		logger.Log("FindTwoUserRelationShip", "Circular Comparation. Must Break of tryng to find")
		return []User{}, -1
	}

	for _, v := range possibleFriends {
		newPossibleFriends = append(newPossibleFriends, GetFriendsID(v)...)
	}

	return FindTwoUserRelationShip(friends, evaluatedFriendShipt, user, newPossibleFriends, counter+1, treshold)

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
