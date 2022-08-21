package main

import (
	"github.com/google/wire"
	"github.com/leadwaymisson/mentorship/api/src/handlers"
	"github.com/leadwaymisson/mentorship/api/src/repo"
)

func InitializeMentorHandler() *handlers.Mentor {
	wire.Build(handlers.NewMentorHandler, repo.NewMentorRepo)
	return &handlers.Mentor{}
 }

 func InitializeBlogHandler() *handlers.Blog {
	wire.Build(handlers.NewBlogHandler, repo.NewBlogRepo)
	return &handlers.Blog{}
 }

 func InitializeMessageUsHandler() *handlers.MessageUs {
	wire.Build(handlers.NewMessageUsHandler, repo.NewMessageUsRepo)
	return &handlers.MessageUs{}
 }