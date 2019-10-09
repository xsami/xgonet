package models

import "testing"

var userList = []User{}

func TestFindUserByID(t *testing.T) {

	dummyUser := User{
		ID:        1,
		Username:  "xsami",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "xsami@xgonet.com",
		Gender:    "M",
		Active:    true}

	user, err := FindUserByID(userList, dummyUser.ID)

	if err != nil {
		t.Error(err)
	}

	if user != dummyUser {
		t.Errorf("User %+v dismatch with found user: %+v", dummyUser, user)
	}
}

func BenchmarkFindUserByID(b *testing.B) {

}

func TestFindUserByUsername(t *testing.T) {

	dummyUser := User{
		ID:        1,
		Username:  "xsami",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "xsami@xgonet.com",
		Gender:    "M",
		Active:    true}

	user, err := FindUserByUsername(userList, dummyUser.Username)

	if err != nil {
		t.Error(err)
	}

	if user != dummyUser {
		t.Errorf("User %+v dismatch with found user: %+v", dummyUser, user)
	}
}

func BenchmarkFindUserByUsername(b *testing.B) {

}

func TestFindUserByEmail(t *testing.T) {

	dummyUser := User{
		ID:        1,
		Username:  "xsami",
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "xsami@xgonet.com",
		Gender:    "M",
		Active:    true}

	user, err := FindUserByEmail(userList, dummyUser.Email)

	if err != nil {
		t.Error(err)
	}

	if user != dummyUser {
		t.Errorf("User %+v dismatch with found user: %+v", dummyUser, user)
	}
}

func BenchmarkFindUserByEmail(b *testing.B) {

}
