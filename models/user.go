package models

// User is the struct that contain the information of a user
type User struct {
	ID        int    `json:"id" bson:"_id"`
	Username  string `json:"username" bson:"username" gorm:"unique_index:idx_name_code"`
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Active    bool   `json:"active" bson:"active"`
}
