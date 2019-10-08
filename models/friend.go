package models

// Friend struct contains the friend relationship
type Friend struct {
	ID         int `json:"id" bson:"_id"`
	UserIDFrom int `json:"from_id" bson:"from_id"`
	UserIDTo   int `json:"to_id" bson:"to_id"`
	StatusID   int `json:"status_id" bson:"status_id"`
}
