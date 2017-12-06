package db

type JobModel struct {
	ID     string `json:"id" bson:"_id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
