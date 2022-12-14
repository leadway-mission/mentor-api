// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/leadwaymisson/mentorship/api/src/handlers"
	"github.com/leadwaymisson/mentorship/api/src/repo"
)

import (
	_ "github.com/labstack/echo/middleware"
	_ "github.com/labstack/echo/v4/middleware"
)

// Injectors from wire.go:

func InitializeMentorHandler() *handlers.Mentor {
	mentor := repo.NewMentorRepo()
	handlersMentor := handlers.NewMentorHandler(mentor)
	return handlersMentor
}

func InitializeBlogHandler() *handlers.Blog {
	blog := repo.NewBlogRepo()
	handlersBlog := handlers.NewBlogHandler(blog)
	return handlersBlog
}

func InitializeMessageUsHandler() *handlers.MessageUs {
	messageUs := repo.NewMessageUsRepo()
	handlersMessageUs := handlers.NewMessageUsHandler(messageUs)
	return handlersMessageUs
}
