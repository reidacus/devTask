package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

type Note struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Note string             `json:"note" bson:"note"`
}
