package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ToDoList model
type ToDoList struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string             `json:"task,omitempty"`
	Status bool               `json:"status,omitempty"`
}

type ToDoListResponse struct {
   ID     string `json:"_id,omitempty"`
   response string `json:"response,omitempty"`
}
