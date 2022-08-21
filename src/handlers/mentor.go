package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leadwaymisson/mentorship/api/src/entities"
	"github.com/leadwaymisson/mentorship/api/src/repo"
	"github.com/leadwaymisson/mentorship/api/src/utils"
)

type Mentor struct{
	Repo repo.IMentor
}

func(m Mentor) Create(c echo.Context) error {
	mail := utils.Mail{}
	var data entities.Mentor
	if err := c.Bind(&data); err != nil {
		log.Printf("Error while binding mentor data %v", err)
		return echo.ErrBadRequest
	}
	if err := m.Repo.Create(data); err != nil {
		log.Printf("Error while binding mentor data %v", err)
		return echo.ErrInternalServerError
	}
	message := utils.Message{
		From: "mentorship@leadwaymission.ca",
		To: data.Email,
		Subject: "New Signup",
	}

	type Data struct {
		Name string
	}

	mailData := Data{Name: data.Firstname}

	if err := mail.Send(message, mailData, "signup.html"); err != nil {
		log.Printf("Error sending email")
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Mentor created",
	})
}

// func(m Mentor) Login(c echo.Context) error {

// 	var loginData entities.FirebaseToken

// 	if err := c.Bind(&loginData); err != nil {
// 		log.Printf("Error while binding login data %v", err)
// 		return echo.ErrBadRequest
// 	}

// 	log.Printf("idToken is %v Email: %v", loginData.IDToken, loginData.Email)
// 	token, err := m.FirebaseUserRepo.VerifyIdToken(loginData.IDToken)
// 	if err != nil {
// 		return &echo.HTTPError{Code: 401, Message: "Invalid credentials, please try again"}
// 	}

// 	user, _ := m.MenteeRepo.FindByEmail(loginData.Email)

// 	if err != nil {
// 		return echo.ErrInternalServerError
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"uid": token.UID,
// 		"user": user,
// 	})
// }

func(m Mentor) Login(c echo.Context) error {
	authID := c.Param("id")
	mentor, err := m.Repo.FetchUserByAuthID(authID)
	if err != nil {
		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"mentor": mentor,
	})
}

func(m Mentor) UpdateUser(c echo.Context) error {
	var mentor entities.Mentor
	authID := c.Param("id")
	if err := c.Bind(&mentor); err != nil {
		log.Printf("Error while binding mentor's data %v", err)
	}

	if err := m.Repo.UpdateUserByAuthID(authID, mentor); err != nil {
		return &echo.HTTPError{}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Data successfully updated.",
	})
}

func NewMentorHandler(repo *repo.Mentor)*Mentor{
	return &Mentor{
		Repo: repo,
	}
}