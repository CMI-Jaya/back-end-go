package model

import "time"

type Appointment struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	DateOfBooking time.Time `json:"date_of_booking"`
	Time          time.Time `json:"time"`
	Status        string    `json:"status"` // Status akan default ke "pending"
	PDFFile       string    `json:"pdf_file"`
	Img           string    `json:"img"`
}
