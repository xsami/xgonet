package models

import "fmt"

// User is the struct that contain the information of a user
type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Active    bool   `json:"active"`
}

// FindUserByID return the user with the given ID
func FindUserByID(users []User, id int) (User, error) {

	resultUsers := FilterUsers(users, func(u User) bool {
		return u.ID == id
	})

	if len(resultUsers) > 0 {
		return resultUsers[0], nil
	}

	return User{}, fmt.Errorf("User with ID=%v was not found", id)
}

// FindUserByUsername return the user with the given username
func FindUserByUsername(users []User, username string) (User, error) {

	resultUsers := FilterUsers(users, func(u User) bool {
		return u.Username == username
	})

	if len(resultUsers) > 0 {
		return resultUsers[0], nil
	}

	return User{}, fmt.Errorf("User with username=%v was not found", username)
}

// FindUserByEmail return the user with the given email
func FindUserByEmail(users []User, email string) (User, error) {

	resultUsers := FilterUsers(users, func(u User) bool {
		return u.Email == email
	})

	if len(resultUsers) > 0 {
		return resultUsers[0], nil
	}

	return User{}, fmt.Errorf("User with email=%v was not found", email)
}

// FindUsersByFirstName return a slice of users with the given firstName
func FindUsersByFirstName(users []User, firstName string) []User {

	return FilterUsers(users, func(u User) bool {
		return u.FirstName == firstName
	})

}

// FindUsersByLastName return a slice of users with the given lastName
func FindUsersByLastName(users []User, lastName string) []User {

	return FilterUsers(users, func(u User) bool {
		return u.LastName == lastName
	})

}

// FindUsersByGender return a slice of users with the given gender
func FindUsersByGender(users []User, gender string) []User {
	return FilterUsers(users, func(u User) bool {
		return u.Gender == gender
	})
}

// FilterUsers return a slice of type User.
// Given a slice of type User, this function evaluate the condition
// that is passed as parameter and return a new slice of slice
// where the condition meet the evaluator and if the user is active
func FilterUsers(users []User, evaluator func(u User) bool) []User {

	resultArray := []User{}
	for _, user := range users {
		if evaluator(user) && user.Active {
			resultArray = append(resultArray, user)
		}
	}

	return resultArray
}
