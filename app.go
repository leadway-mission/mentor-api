package main

import (
	_ "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/echo/v4/middleware"
)

func main() {
	r := echo.New()

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://labstack.net"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	  }))

	r.Use(middleware.Logger())
	mentor := InitializeMentorHandler()
	blog := InitializeBlogHandler()
	messageUs := InitializeMessageUsHandler()

	e := r.Group("/api")
	e.POST("/mentors", mentor.Create)
	e.GET("/mentors/:id", mentor.Login)
	e.PUT("/mentors/:id", mentor.UpdateUser)

	e.POST("/blogs", blog.Create)
	e.GET("/blogs", blog.FetchAll)

	e.POST("/message-us", messageUs.Create)

	r.Logger.Fatal(r.Start(":1323"))
}