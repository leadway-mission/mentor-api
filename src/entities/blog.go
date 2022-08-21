package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title" bson:"title"`
	ShortDescription string `json:"short_description" bson:"short_description"`
	Content string `json:"content" bson:"content"`
}