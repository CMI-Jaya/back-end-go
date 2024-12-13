package model

import "time"

type Appointment struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	DateOfBooking time.Time `json:"date_of_booking"`
	LinkMeet      string    `json:"link_meet"`
	PDFFile       string    `json:"pdf_file"`
	Img           string    `json:"img"`
	Time          time.Time `json:"time"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
