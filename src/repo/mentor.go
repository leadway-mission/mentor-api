package repo

import (
	"context"
	"errors"
	"log"

	"github.com/leadwaymisson/mentorship/api/src/drivers"
	"github.com/leadwaymisson/mentorship/api/src/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type IMentor interface{
	Create(m entities.Mentor) error
	FetchUserByAuthID(id string) (*entities.Mentor, error)
	UpdateUserByAuthID(id string, data entities.Mentor) error 
}

type Mentor struct{}

func(m Mentor) Create(data entities.Mentor) error {
	col := drivers.DB.Collection("mentors")
	ctx := context.Background()
	if _, err := col.InsertOne(ctx, data); err != nil {
		log.Printf("Error while creating mentor")
	}
	return nil
}

func(m Mentor) FetchUserByAuthID(id string) (*entities.Mentor, error) {
	var mentor entities.Mentor
	col := drivers.DB.Collection("mentors")
	ctx := context.Background()
	filter := bson.M{"auth_id": id}
	if err := col.FindOne(ctx, filter).Decode(&mentor); err != nil {
		log.Printf("Error while fetching mentor %v", err)
		return nil, err
	}
	return &mentor, nil
}

func(m Mentor) UpdateUserByAuthID(id string, data entities.Mentor) error {
	col := drivers.DB.Collection("mentors")
	ctx := context.Background()
	filter := bson.M{"auth_id": id}
	log.Printf("ID %v", id)
	log.Printf("Profession %v", data.Profession)

	update := 
		bson.M{"$set": 
			bson.M{
				"profession": data.Profession,
				"city": data.City,
				"phone_number": data.Phonenumber,
				"description": data.Description,
			},
		}
		
	if _, err := col.UpdateOne(ctx, filter, update); err != nil {
		log.Printf("Error while updating mentor %v", err)
		return errors.New("Mentor could not be updated")
	}
	return nil
}

func NewMentorRepo()*Mentor {
	return &Mentor{}
}