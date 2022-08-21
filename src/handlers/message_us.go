package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leadwaymisson/mentorship/api/src/entities"
	"github.com/leadwaymisson/mentorship/api/src/repo"
)

type MessageUs struct{
	Repo repo.IMessageUs
}

func(m MessageUs) Create(c echo.Context) error {

	var data entities.MessageUs
	if err := c.Bind(&data); err != nil {
		log.Printf("Error while binding message us data %v", err)
		return echo.ErrBadRequest
	}

	if err := m.Repo.Create(data); err != nil {
		log.Printf("Error while binding message us data %v", err)
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "created",
	})
}

func NewMessageUsHandler(repo *repo.MessageUs)*MessageUs{
	return &MessageUs{
		Repo: repo,
	}
}