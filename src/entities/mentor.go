package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mentor struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lasttname string `json:"lastname" bson:"lastname"`
	Email string `json:"email" bson:"email"`
	City string `json:"city" bson:"city"`
	Password string `json:"password" bson:"password"`
	AuthID string `json:"auth_id" bson:"auth_id"`
	Profession string `json:"profession" bson:"profession"`
	Phonenumber string `json:"phone_number" bson:"phone_number"`
	Description string `json:"description" bson:"description"`
}