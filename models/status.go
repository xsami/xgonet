package models

// Status struct represent all the status
// These status could be: user blocked, waiting acceptance, declined
type Status struct {
	ID          int    `json:"id" bson:"_id"`
	Value       string `json:"value" bson:"value"`
	Description string `json:"descriptio n" bson:"description"`
}
