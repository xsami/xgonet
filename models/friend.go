package models

import "log"

// SearchTreshHold contain the depth amount of search
// that will use the application to complete the FindTwoUserRelationShip
// as this can enter into a circular recursion that may be a bit isssue
const SearchTreshHold = 10

// Friend struct contains the friend relationship
type Friend struct {
	ID         int  `json:"id" bson:"_id"`
	UserIDFrom int  `json:"from_id" bson:"from_id"`
	UserIDTo   int  `json:"to_id" bson:"to_id"`
	Accepted   bool `json:"accepted" bson:"accepted"`
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

		friendUser, err := FindUserByID(GetUserList(), friendID)
		if err != nil {
			log.Fatal("FindUserFriends failed.", err) // As couldn't find the user, or the user isn't active
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
func FindTwoUserRelationShip(friends []Friend, userA User, userB User, counter int) ([]User, int) {

	if counter >= SearchTreshHold {
		return []User{}, -1
	}

	if ValidateFriendShip(friends, userA.ID, userB.ID) {
		return []User{}, counter // TODO: hold the users where the relationship begins
	}

	userAFriends := FindUserFriends(friends, userA)
	userBFriends := FindUserFriends(friends, userB)

	for _, uA := range userAFriends {
		for _, uB := range userBFriends {
			return FindTwoUserRelationShip(friends, uA, uB, counter+1)
		}
	}

	return []User{}, 0
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

// GetFriendList return a slice with all the friends
// from the mock data file
func GetFriendList() []Friend {

	return []Friend{}
}
