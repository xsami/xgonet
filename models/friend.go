package models

import "log"

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
			log.Fatal("FindUserFriends failed.", err) // As couldn't find the user
		} else {
			if friendUser.Active { // Only add active users
				resultUsers = append(resultUsers, friendUser)
			}
		}
	}

	return resultUsers
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
