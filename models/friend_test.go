package models

import "testing"

var friendList = []Friend{}

func TestFindUserFriends(t *testing.T) {

	dummyUser := User{
		ID:        1,
		Username:  "xsami",
		FirstName: "first_name",
		LastName:  "last_name",
		Gender:    "M",
		Active:    true}

	result := FindUserFriends(friendList, dummyUser)

	if len(result) < 1 {
		t.Errorf("Failed finding the amount of friend for the user %v. It found: %v", dummyUser, result)
	}
}

func BenchmarkFindUserFriends(b *testing.B) {

}
