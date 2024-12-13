package routes

import (
	"go-project/internal/admin/handler"

	"github.com/gorilla/mux"
)

// FUNCTION REGISTER ADMIN RESTFULLAPI
func RegisterAdminRoutes(router *mux.Router,
	articleHandler *handler.ArticleHandler,
	videoHandler *handler.VideoHandler,
	appointmentHandler *handler.AppointmentHandler,
	testimonialHandler *handler.TestimonialHandler,
	commentHandler *handler.CommentHandler,
	webinarHandler *handler.WebinarHandler) {

	// ROUTES ARTICLE ADMIN || CRUD ||
	router.HandleFunc("/admin/articles", articleHandler.GetAllArticles).Methods("GET")
	router.HandleFunc("/admin/article", articleHandler.CreateArticle).Methods("POST")
	router.HandleFunc("/admin/article/update", articleHandler.UpdateArticle).Methods("PUT")
	router.HandleFunc("/admin/article/delete", articleHandler.DeleteArticle).Methods("DELETE")
	router.HandleFunc("/admin/article/view", articleHandler.GetArticleByID).Methods("GET")

	// ROUTES VIDEO ADMIN || CRUD ||
	router.HandleFunc("/admin/videos", videoHandler.GetAllVideos).Methods("GET")
	router.HandleFunc("/admin/video", videoHandler.CreateVideo).Methods("POST")
	router.HandleFunc("/admin/video/update", videoHandler.UpdateVideo).Methods("PUT")
	router.HandleFunc("/admin/video/delete", videoHandler.DeleteVideo).Methods("DELETE")
	router.HandleFunc("/admin/video/view", videoHandler.GetVideoByID).Methods("GET")

	// ROUTES APPOINTMENT ADMIN || ASSGIN HOST || CREATE || UPDATE ||
	router.HandleFunc("/admin/staff", appointmentHandler.GetStaffList).Methods("GET")
	router.HandleFunc("/admin/appointments", appointmentHandler.CreateAppointment).Methods("POST")
	router.HandleFunc("/admin/appointments/{id}/assign-host", appointmentHandler.AssignHost).Methods("POST")
	router.HandleFunc("/admin/appointments/{id}/update-status", appointmentHandler.UpdateStatus).Methods("PUT")

	// ROUTES TESTIMONIALS ADMIN || CRUD ||
	router.HandleFunc("/admin/testimonials", testimonialHandler.GetAllTestimonials).Methods("GET")
	router.HandleFunc("/admin/testimonial", testimonialHandler.CreateTestimonial).Methods("POST")
	router.HandleFunc("/admin/testimonial/{id:[0-9]+}", testimonialHandler.GetTestimonialByID).Methods("GET")
	router.HandleFunc("/admin/testimonial/{id:[0-9]+}", testimonialHandler.UpdateTestimonial).Methods("PUT")
	router.HandleFunc("/admin/testimonial/{id:[0-9]+}", testimonialHandler.DeleteTestimonial).Methods("DELETE")
	router.HandleFunc("/admin/testimonial/{id:[0-9]+}/approve", testimonialHandler.ApproveTestimonial).Methods("PUT")
	router.HandleFunc("/admin/testimonial/{id:[0-9]+}/reject", testimonialHandler.RejectTestimonial).Methods("PUT")

	// ROUTES COMMENT ADMIN || APPROVE || REJECT || REPLY||
	router.HandleFunc("/admin/comments", commentHandler.GetAllComments).Methods("GET")
	router.HandleFunc("/admin/comment", commentHandler.CreateComment).Methods("POST")
	router.HandleFunc("/admin/comment/{id:[0-9]+}/approve", commentHandler.ApproveComment).Methods("PUT")
	router.HandleFunc("/admin/comment/{id:[0-9]+}/reject", commentHandler.RejectComment).Methods("PUT")
	router.HandleFunc("/admin/comment/{id:[0-9]+}/delete", commentHandler.DeleteComment).Methods("DELETE")
	router.HandleFunc("/admin/comment/{id:[0-9]+}/reply", commentHandler.ReplyComment).Methods("POST")

	router.HandleFunc("/admin/webinar", webinarHandler.CreateWebinar).Methods("POST")

	// Auth Routes
	router.HandleFunc("/admin/register", handler.RegisterAdmin).Methods("POST")

	router.HandleFunc("/admin/login", handler.LoginAdmin).Methods("POST")
	router.HandleFunc("/admin/logout", handler.LogoutAdmin).Methods("POST")
}
