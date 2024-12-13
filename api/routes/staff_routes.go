package routes

import (
	"go-project/internal/staff/handler"
	"net/http"

	"github.com/gorilla/mux"
)

// FUNCTION REGISTER STAFF RESTFULLAPI
func RegisterStaffRoutes(router *mux.Router, articleHandler *handler.ArticleHandler, videoHandler *handler.VideoHandler, appointmentHandler *handler.AppointmentHandler, handler *handler.TestimonialHandler, commentHandler *handler.CommentHandler, webinarHandler *handler.WebinarHandler) {
	// ROUTES STAFF ARTICLE || CRUD ||
	router.HandleFunc("/staff/upload/articles", articleHandler.UploadArticle).Methods(http.MethodPost)
	router.HandleFunc("/staff/articles/view", articleHandler.GetArticleByID).Methods(http.MethodGet)
	router.HandleFunc("/staff/articles", articleHandler.GetAllArticles).Methods(http.MethodGet)

	// ROUTES STAFF VIDEO || CRUD ||
	router.HandleFunc("/staff/upload/videos", videoHandler.UploadVideo).Methods(http.MethodPost)
	router.HandleFunc("/staff/videos/view", videoHandler.GetVideoByID).Methods(http.MethodGet)
	router.HandleFunc("/staff/videos", videoHandler.GetAllVideos).Methods(http.MethodGet)

	// ROUTES STAFF TESTIMONIALS || CREATE || GET PENDING || UPDATE || DELETE ||
	router.HandleFunc("/staff/testimonials", handler.CreateTestimonial).Methods("POST")
	router.HandleFunc("/staff/testimonials", handler.GetPendingTestimonials).Methods("GET")
	router.HandleFunc("/staff/testimonials/{id}", handler.UpdateTestimonial).Methods("PUT")
	router.HandleFunc("/staff/testimonials/{id}", handler.DeleteTestimonial).Methods("DELETE")

	// ROUTES STAFF COMMENTS || GET ALL COMMENTS || DELETE STAFF COMMENT || DELETE USER COMMENT || CREATE || REPLY COMMENT ||
	router.HandleFunc("/staff/comments", commentHandler.GetAllComments).Methods("GET")
	router.HandleFunc("/staff/comments/{id}", commentHandler.DeleteOwnComment).Methods("DELETE")
	router.HandleFunc("/staff/comments/user/{id}", commentHandler.DeleteUserComment).Methods("DELETE")
	router.HandleFunc("/staff/comments", commentHandler.CreateComment).Methods("POST")
	router.HandleFunc("/staff/comments/reply/{id}", commentHandler.ReplyComment).Methods("POST")

	router.HandleFunc("/staff/webinars", webinarHandler.GetAllWebinars).Methods("GET")
	router.HandleFunc("/staff/webinar/view", webinarHandler.GetWebinarByID).Methods("GET")

	// ROUTES STAFF APPOINTMENTS || CREATE APPOINTMENTS || LIST APPOINTMENTS
	router.HandleFunc("/staff/appointments", appointmentHandler.CreateAppointment).Methods(http.MethodPost)
	router.HandleFunc("/staff/appointments", appointmentHandler.ListAppointments).Methods(http.MethodGet)
}
