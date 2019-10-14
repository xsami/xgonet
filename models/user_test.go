package models

import "testing"

// Global testUserList for testing userlist functionality
var testUserList = []User{
	{
		ID:        1,
		Username:  "mhallihan0",
		FirstName: "Marla",
		LastName:  "Hallihan",
		Email:     "mhallihan0@google.co.jp",
		Gender:    "F",
		Active:    true},
	{
		ID:        2,
		Username:  "mbrereton1",
		FirstName: "Mozes",
		LastName:  "Brereton",
		Email:     "mbrereton1@icq.com",
		Gender:    "M",
		Active:    true},
	{
		ID:        3,
		Username:  "klaverack2",
		FirstName: "Kenny",
		LastName:  "Laverack",
		Email:     "klaverack2@pinterest.com",
		Gender:    "M",
		Active:    true},
	{
		ID:        4,
		Username:  "mbreit3",
		FirstName: "Marla",
		LastName:  "Breit",
		Email:     "mbreit3@huffingtonpost.com",
		Gender:    "F",
		Active:    true},
	{
		ID:        5,
		Username:  "lquirke4",
		FirstName: "Lorant",
		LastName:  "Quirke",
		Email:     "lquirke4@bluehost.com",
		Gender:    "M",
		Active:    true}}

func TestFindUserByID(t *testing.T) {

	dummyUser := User{
		ID:        1,
		Username:  "mhallihan0",
		FirstName: "Marla",
		LastName:  "Hallihan",
		Email:     "mhallihan0@google.co.jp",
		Gender:    "F",
		Active:    true}

	user, err := FindUserByID(testUserList, dummyUser.ID)

	if err != nil {
		t.Error(err)
	}

	if user != dummyUser {
		t.Errorf("User %+v dismatch with found user: %+v", dummyUser, user)
	}
}

func BenchmarkFindUserByID(b *testing.B) {

	userListLen := len(testUserList)
	for i := 0; i < b.N; i++ {
		userID := testUserList[i%userListLen].ID
		FindUserByID(testUserList, userID)
	}

}

func TestFindUserByUsername(t *testing.T) {

	dummyUser := User{
		ID:        5,
		Username:  "lquirke4",
		FirstName: "Lorant",
		LastName:  "Quirke",
		Email:     "lquirke4@bluehost.com",
		Gender:    "M",
		Active:    true}

	user, err := FindUserByUsername(testUserList, dummyUser.Username)

	if err != nil {
		t.Error(err)
	}

	if user != dummyUser {
		t.Errorf("User %+v dismatch with found user: %+v", dummyUser, user)
	}
}

func BenchmarkFindUserByUsername(b *testing.B) {

	userListLen := len(testUserList)
	for i := 0; i < b.N; i++ {
		username := testUserList[i%userListLen].Username
		FindUserByUsername(testUserList, username)
	}

}

func TestFindUserByEmail(t *testing.T) {

	dummyUser := User{
		ID:        4,
		Username:  "mbreit3",
		FirstName: "Marla",
		LastName:  "Breit",
		Email:     "mbreit3@huffingtonpost.com",
		Gender:    "F",
		Active:    true}

	user, err := FindUserByEmail(testUserList, dummyUser.Email)

	if err != nil {
		t.Error(err)
	}

	if user != dummyUser {
		t.Errorf("User %+v dismatch with found user: %+v", dummyUser, user)
	}
}

func BenchmarkFindUserByEmail(b *testing.B) {

	userListLen := len(testUserList)
	for i := 0; i < b.N; i++ {
		email := testUserList[i%userListLen].Email
		FindUserByEmail(testUserList, email)
	}

}

func TestFindUsersByFirstName(t *testing.T) {

	const firstName = "Marla"
	const marlasCnt = 2

	usersResult := FindUsersByFirstName(testUserList, firstName)
	resultAmount := len(usersResult)

	if resultAmount != marlasCnt {
		t.Errorf("Incorrect amount of users with first name: %v.\nThe program found %v instead of %v", firstName, resultAmount, marlasCnt)
	}
}

func BenchmarkFindUsersByFirstName(b *testing.B) {

	userListLen := len(testUserList)
	for i := 0; i < b.N; i++ {
		firstName := testUserList[i%userListLen].FirstName
		FindUsersByFirstName(testUserList, firstName)
	}

}

func TestFindUsersByLastName(t *testing.T) {

	const firstName = "Laverack"
	const laverCnt = 1

	usersResult := FindUsersByLastName(testUserList, firstName)
	resultAmount := len(usersResult)

	if resultAmount != laverCnt {
		t.Errorf("Incorrect amount of users with last name: %v.\nThe program found %v instead of %v", firstName, resultAmount, laverCnt)
	}
}

func BenchmarkFindUsersByLastName(b *testing.B) {

	userListLen := len(testUserList)
	for i := 0; i < b.N; i++ {
		lastName := testUserList[i%userListLen].LastName
		FindUsersByLastName(testUserList, lastName)
	}

}
func TestFindUsersByGender(t *testing.T) {

	const gender = "F"
	const laverCnt = 2

	usersResult := FindUsersByGender(testUserList, gender)
	resultAmount := len(usersResult)

	if resultAmount != laverCnt {
		t.Errorf("Incorrect amount of users with gender: %v.\nThe program found %v instead of %v", gender, resultAmount, laverCnt)
	}
}

func BenchmarkFindUsersByGender(b *testing.B) {

	userListLen := len(testUserList)
	for i := 0; i < b.N; i++ {
		gender := testUserList[i%userListLen].Gender
		FindUsersByGender(testUserList, gender)
	}

}
