package model

import (
	"time"
)

type Appointment struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	PhoneNumber   string    `json:"phone_number"`
	Email         string    `json:"email"`
	DateOfBooking time.Time `json:"date_of_booking"`
	Time          time.Time `json:"time"`
	LinkMeet      string    `json:"link_meet"`
	HostID        int       `json:"host_id"`
	Status        string    `json:"status"`
	PDFFile       string    `json:"pdf_file"` // New field for PDF file
	Img           string    `json:"img"`      // New field for image
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
type Staff struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
