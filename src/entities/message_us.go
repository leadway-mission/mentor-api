package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type MessageUs struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Subject string `json:"subject" bson:"subject"`
	Message string `json:"message" bson:"message"`
}