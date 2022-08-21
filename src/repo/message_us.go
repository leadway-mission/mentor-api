package repo

import (
	"context"
	"log"

	"github.com/leadwaymisson/mentorship/api/src/drivers"
	"github.com/leadwaymisson/mentorship/api/src/entities"
)

type IMessageUs interface {
	Create(data entities.MessageUs) error
}

type MessageUs struct {}

func(m MessageUs) Create(data entities.MessageUs) error{
	col := drivers.DB.Collection("message_us")
	ctx := context.Background()
	if _, err := col.InsertOne(ctx, data); err != nil {
		log.Printf("Error while creating message us")
	}
	return nil
}

func NewMessageUsRepo()*MessageUs {
	return &MessageUs{}
}