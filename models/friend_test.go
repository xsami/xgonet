package models

import "testing"

var testFriendList = []Friend{
	{
		ID:         1,
		UserIDFrom: 1,
		UserIDTo:   2,
		Accepted:   true},
	{
		ID:         2,
		UserIDFrom: 3,
		UserIDTo:   4,
		Accepted:   true},
	{
		ID:         3,
		UserIDFrom: 2,
		UserIDTo:   5,
		Accepted:   true},
	{
		ID:         4,
		UserIDFrom: 5,
		UserIDTo:   4,
		Accepted:   true}}

func injectUser() {
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

	UserList = testUserList
	FriendMap = BuildFriendMap(testFriendList)
}

func TestFindUserFriends(t *testing.T) {

	injectUser()
	dummyUser := User{
		ID:        1,
		Username:  "mhallihan0",
		FirstName: "Marla",
		LastName:  "Hallihan",
		Email:     "mhallihan0@google.co.jp",
		Gender:    "F",
		Active:    true}

	result := FindUserFriends(testFriendList, dummyUser)

	if len(result) != 1 {
		t.Errorf("Failed finding the amount of friend for the user %v. It found: %v", dummyUser, result)
	}
}

func BenchmarkFindUserFriends(b *testing.B) {

	injectUser()
	lenUserList := len(UserList)
	for i := 0; i < b.N; i++ {
		user := UserList[i%lenUserList]
		FindUserFriends(testFriendList, user)
	}

}

func TestValidateFriendShip(t *testing.T) {

	dummyUserA := User{
		ID:        1,
		Username:  "mhallihan0",
		FirstName: "Marla",
		LastName:  "Hallihan",
		Email:     "mhallihan0@google.co.jp",
		Gender:    "F",
		Active:    true}

	dummyUserB := User{
		ID:        2,
		Username:  "mbrereton1",
		FirstName: "Mozes",
		LastName:  "Brereton",
		Email:     "mbrereton1@icq.com",
		Gender:    "M",
		Active:    true}

	if !ValidateFriendShip(testFriendList, dummyUserA.ID, dummyUserB.ID) {
		t.Errorf("Users %+v and %+v are friends. But the response was false", dummyUserA, dummyUserB)
	}

}

func BenchmarkValidateFriendShip(b *testing.B) {

	injectUser()
	lenUserList := len(UserList)
	for i := 0; i < b.N; i++ {
		userA := UserList[i%lenUserList]
		userB := UserList[(i+1)%lenUserList]
		ValidateFriendShip(testFriendList, userA.ID, userB.ID)
	}

}

func TestFindTwoUserRelationShip(t *testing.T) {
	injectUser()
	var tresHold uint

	userA := User{
		ID:        1,
		Username:  "mhallihan0",
		FirstName: "Marla",
		LastName:  "Hallihan",
		Email:     "mhallihan0@google.co.jp",
		Gender:    "F",
		Active:    true}
	userB := User{
		ID:        2,
		Username:  "mbrereton1",
		FirstName: "Mozes",
		LastName:  "Brereton",
		Email:     "mbrereton1@icq.com",
		Gender:    "M",
		Active:    true}

	res, counter := FindTwoUserRelationShip(FriendMap, make(map[RelateFriend]bool, len(FriendMap)), userA, []int{userB.ID}, 0, tresHold)

	if counter != 0 {
		t.Error("Failed on the response for friendship: ", res, " : ", counter)
	}
}

func BenchmarkFindTwoUserRelationShip(b *testing.B) {

	injectUser()
	var tresHold uint = 0
	lenUserList := len(UserList)

	for i := 0; i < b.N; i++ {
		userA := UserList[i%lenUserList]
		userB := UserList[(i+1)%lenUserList]
		FindTwoUserRelationShip(FriendMap, make(map[RelateFriend]bool, len(FriendMap)), userA, []int{userB.ID}, 0, tresHold)
	}
}
