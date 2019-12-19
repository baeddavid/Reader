package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"task, omitempty"`
	Author string `json:"task,omitempty"`
	Status string `json:"task,omitempty"`
}