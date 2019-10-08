package models

import "testing"

var userList = []User{}

func TestFindUserByID(t *testing.T) {

	dummyUser := User{
		ID:        1,
		Username:  "xsami",
		FirstName: "first_name",
		LastName:  "last_name",
		Gender:    "M",
		Active:    true}

	user, err := FindUserByID(userList, dummyUser.ID)

	if err != nil {
		t.Error(err)
	}

	if user != dummyUser {
		t.Errorf("User %v dismatch with found user: %v", dummyUser, user)
	}
}

func BenchmarkFindUserByID(b *testing.B) {

}
