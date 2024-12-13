package routes

import (
	"go-project/internal/user/handler"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(
	router *mux.Router,
	appointmentHandler *handler.AppointmentHandler,
) {
	router.HandleFunc("/user/appointments", appointmentHandler.CreateAppointment).Methods("POST")
	// router.HandleFunc("/user/appointments", appointmentHandler.Get).Methods("POST")

}
