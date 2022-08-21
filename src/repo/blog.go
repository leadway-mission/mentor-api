package repo

import (
	"context"
	"log"

	"github.com/leadwaymisson/mentorship/api/src/drivers"
	"github.com/leadwaymisson/mentorship/api/src/entities"
	"go.mongodb.org/mongo-driver/bson"
)

type IBlog interface {
	Create(data entities.Blog) error
	FetchAll()([]entities.Blog, error)
	FetchByID(id string)(*entities.Blog, error)
}

type Blog struct {}

func(b Blog) Create(data entities.Blog) error{
	col := drivers.DB.Collection("blogs")
	ctx := context.Background()
	if _, err := col.InsertOne(ctx, data); err != nil {
		log.Printf("Error while creating blog")
	}
	return nil
}

func(b Blog) FetchAll()([]entities.Blog, error){
	var blogs []entities.Blog
	col := drivers.DB.Collection("blogs")
	ctx := context.Background()

	cur, err := col.Find(ctx, bson.M{})

	if err != nil {
		log.Printf("Error while fetching blogs")
		return nil, err
	}
	
	if err := cur.All(ctx, &blogs); err != nil {
		log.Printf("Error while fetching blogs")
	}
	return blogs, nil
}

func(b Blog) FetchByID(id string)(*entities.Blog, error){
	var blog entities.Blog
	col := drivers.DB.Collection("blogs")
	ctx := context.Background()

	if err := col.FindOne(ctx, bson.M{}).Decode(&blog); err != nil {
		log.Printf("Error while fetching blog %v", err)
		return nil, err
	}
	return &blog, nil
}

func NewBlogRepo()*Blog {
	return &Blog{}
}