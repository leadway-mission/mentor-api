package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"log"
	"os"
	"gopkg.in/gomail.v2"
)

type Message struct{
	From string
	To string
	Subject string
	Body string
}

type Mail struct {}

func(ma Mail) Send(message Message, data interface{}, templateName string) error {

	password := os.Getenv("EMAIL_PASSWORD")
	host := os.Getenv("EMAIL_HOST")
	from := os.Getenv("EMAIL_USERNAME")

	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
    msg += fmt.Sprintf("From: %s\r\n", message.From)
    msg += fmt.Sprintf("To: %s\r\n;", message.To)
    msg += fmt.Sprintf("Subject: %s\r\n", message.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", message.Body)

	t, err := template.ParseFiles(fmt.Sprintf("src/templates/%v", templateName))
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer

	err = t.Execute(&tpl, data)

	if err != nil {
		log.Println(err)
	}

	result := tpl.String()

	m := gomail.NewMessage()
	
	m.SetAddressHeader("From", from , "Mentorship")
	m.SetHeader("To", message.To)
	
	m.SetHeader("Subject", message.Subject)
	m.SetBody("text/html", result)

	d := gomail.NewDialer(host, 587, from, password,)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error while sending email %v", err)
		return err
	}
	return nil
}