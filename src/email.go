package main

import (
	"os"
	// "github.com/joho/godotenv"
	"time"
	"net/smtp"
	"strings"
	"log"
)

// Sends email using SMTP
func Emailer(message string) error {
	// uncomment below to load envs vars from .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	host := strings.TrimSpace(os.Getenv("HOST"))
	port := strings.TrimSpace(os.Getenv("PORT"))
	from := strings.TrimSpace(os.Getenv("FROM"))
	to := strings.TrimSpace(os.Getenv("TO"))
	pwd := strings.TrimSpace(os.Getenv("PASSWD"))
	addr := host + ":" + port
	auth := smtp.PlainAuth("", from, pwd, host)

	log.Println("Sending email from " + from + " to " + to)
	subject := "Daily RSS Feeds for " + time.Now().Format("Jan 02, 2006")
	body := strings.Builder{}
	body.WriteString("From: \"io-golang: Daily RSS Feed\" <" + from + ">\n")
	body.WriteString("To: " + to + "\n")
	body.WriteString("Subject: " + subject + "\n")
	body.WriteString("MIME-version: 1.0;\n")
	body.WriteString("Content-Type: text/html;charset=\"UTF-8\";\n")
	body.WriteString("\n")
	body.WriteString(message)

	err := smtp.SendMail(addr, auth, from, []string{to}, []byte(body.String()))
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}