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

func TestValidateFriendShip(t *testing.T) {

	dummyUserA := User{
		ID:        1,
		Username:  "xsami",
		FirstName: "first_name",
		LastName:  "last_name",
		Gender:    "M",
		Active:    true}

	dummyUserB := User{
		ID:        2,
		Username:  "jhony",
		FirstName: "first_name",
		LastName:  "last_name",
		Gender:    "M",
		Active:    true}

	dummyUserC := User{
		ID:        3,
		Username:  "patrick",
		FirstName: "first_name",
		LastName:  "last_name",
		Gender:    "M",
		Active:    true}

	if ValidateFriendShip(friendList, dummyUserA.ID, dummyUserB.ID) {
		t.Errorf("Users %+v and %+v should not be friends", dummyUserA, dummyUserB)
	}

	if !ValidateFriendShip(friendList, dummyUserA.ID, dummyUserC.ID) {
		t.Errorf("Users %+v and %+v are friends but the response was false", dummyUserA, dummyUserC)
	}
}

func BenchmarkValidateFriendShip(b *testing.B) {

}

func TestFindTwoUserRelationShip(t *testing.T) {

	dummyUserA := User{
		ID:        1,
		Username:  "xsami",
		FirstName: "first_name",
		LastName:  "last_name",
		Gender:    "M",
		Active:    true}

	dummyUserB := User{
		ID:        2,
		Username:  "jhony",
		FirstName: "first_name",
		LastName:  "last_name",
		Gender:    "M",
		Active:    true}

	relationShip, counter := FindTwoUserRelationShip(friendList, dummyUserA, dummyUserB, 0)

	t.Logf("Users: %+v and %+v have the following relationship:\n%+v\n\nWith the following amount of iterations: %v", dummyUserA, dummyUserB, relationShip, counter)
}

func BenchmarkFindTwoUserRelationShip(b *testing.B) {

}
