package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leadwaymisson/mentorship/api/src/entities"
	"github.com/leadwaymisson/mentorship/api/src/repo"
)

type Blog struct{
	Repo repo.IBlog
}

func(m Blog) Create(c echo.Context) error {

	var data entities.Blog
	if err := c.Bind(&data); err != nil {
		log.Printf("Error while binding blog data %v", err)
		return echo.ErrBadRequest
	}

	if err := m.Repo.Create(data); err != nil {
		log.Printf("Error while binding blog data %v", err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "created",
	})
}

func(m Blog) FetchAll(c echo.Context) error {

	blogs, err := m.Repo.FetchAll()
	if err != nil {
		log.Printf("Error while binding blog data %v", err)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": blogs,
	})
}

func NewBlogHandler(repo *repo.Blog)*Blog{
	return &Blog{
		Repo: repo,
	}
}