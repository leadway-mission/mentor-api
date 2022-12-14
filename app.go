package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/labstack/echo/v4/middleware"
)

func main() {
	var port string
	if err := godotenv.Load(".env"); err != nil{
		port = "8181"
	} else {
		port = os.Getenv("PORT")
	}

	r := echo.New()

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://labstack.net", "http://localhost:8080", "http://localhost:8081"},
	  }))

	r.Use(middleware.Logger())
	mentor := InitializeMentorHandler()
	blog := InitializeBlogHandler()
	messageUs := InitializeMessageUsHandler()

	e := r.Group("/api")
	e.POST("/mentors", mentor.Create)
	e.GET("/mentors/:id", mentor.Login)
	e.PUT("/mentors/:id", mentor.UpdateUser)
	e.GET("/mentors/:id/fetch", mentor.FetchByID)

	e.POST("/blogs", blog.Create)
	e.GET("/blogs", blog.FetchAll)

	e.POST("/message-us", messageUs.Create)

	r.Logger.Fatal(r.Start(":" + port))
}
