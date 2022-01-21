package email

import (
	"errors"
	"log"
	"net/smtp"
	"ofspace-be/features/booking"
	"os"
)

func Smtp(to []string, data interface{}, template string, bookData booking.Core) (bool, error) {
	var email_type string
	if bookData.BookingStatus == "deal" {
		email_type = "Booking"
	} else {
		email_type = "Review"
	}
	EMAIL := os.Getenv("EMAIL")
	PASSWORD := os.Getenv("PASSWORD")
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	Body, err := templateParse(template, data, bookData)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}
	subject := "Subject: " + email_type + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + "\n" + Body)
	status := smtp.SendMail(HOST+":"+PORT, smtp.PlainAuth("", EMAIL, PASSWORD, HOST), EMAIL, to, []byte(message))
	if status != nil {
		log.Printf("Error from SMTP Server: %s", status)
	}
	log.Print("Email Sent Successfully")
	return true, nil
}
