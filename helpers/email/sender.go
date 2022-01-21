package email

import (
	"log"
	"ofspace-be/features/booking"
	"time"
)

type Message struct {
	ReceiverName  string
	SenderName    string
	Total         float32
	Type          string
	Date          time.Time
	BookingStatus string
	Email         string
}

func SendEmail(bookData booking.Core) error {
	data := Message{
		ReceiverName:  bookData.ConfirmedName,
		SenderName:    "Ofspace",
		Total:         bookData.Price,
		Type:          bookData.Unit.UnitType,
		Date:          bookData.DealDate,
		BookingStatus: bookData.BookingStatus,
		Email:         bookData.User.Email,
	}
	to := []string{bookData.User.Email}
	var message string
	if bookData.BookingStatus == "deal" {
		message = "booking.html"
	} else {
		message = "review.html"
	}
	status, err := Smtp(to, data, message, bookData)
	if err != nil {
		log.Println(err)
		return err
	}
	if status {
		log.Println("Email Succes!!")
	}
	return nil
}
