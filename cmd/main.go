package main

import (
	"go-project/api/routes"
	"go-project/config"
	"go-project/db"
	adminHandler "go-project/internal/admin/handler"
	adminRepo "go-project/internal/admin/repository"
	adminService "go-project/internal/admin/service"
	staffHandler "go-project/internal/staff/handler"
	staffRepo "go-project/internal/staff/repository"
	staffService "go-project/internal/staff/service"
	userHandler "go-project/internal/user/handler"
	userRepo "go-project/internal/user/repository"
	userService "go-project/internal/user/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Recovery untuk menghindari panic yang menyebabkan aplikasi crash
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db.ConnectDB(cfg)
	defer db.DB.Close() // Ensure the DB connection is closed when main exits.

	// Initialize router
	router := mux.NewRouter()

	// Admin initialization
	adminArticleRepo := adminRepo.NewArticleRepository(db.DB)
	adminArticleService := adminService.NewArticleService(adminArticleRepo)
	adminArticleHandler := adminHandler.NewArticleHandler(adminArticleService)

	adminVideoRepo := adminRepo.NewVideoRepository(db.DB)
	adminVideoService := adminService.NewVideoService(adminVideoRepo)
	adminVideoHandler := adminHandler.NewVideoHandler(adminVideoService)

	// Appointment initialization for Admin
	adminAppointmentRepo := adminRepo.NewAppointmentRepository(db.DB)
	adminAppointmentService := adminService.NewAppointmentService(adminAppointmentRepo)
	adminAppointmentHandler := adminHandler.NewAppointmentHandler(adminAppointmentService)

	// Testimonial initialization for Admin
	adminTestimonialRepo := adminRepo.NewTestimonialRepository(db.DB)
	adminTestimonialService := adminService.NewTestimonialService(adminTestimonialRepo)
	adminTestimonialHandler := adminHandler.NewTestimonialHandler(adminTestimonialService)

	adminCommentRepo := adminRepo.NewCommentRepository(db.DB)
	adminCommentService := adminService.NewCommentService(adminCommentRepo)
	adminCommentHandler := adminHandler.NewCommentHandler(adminCommentService)

	adminWebinarRepo := adminRepo.NewWebinarRepository(db.DB)
	adminWebinarService := adminService.NewWebinarService(adminWebinarRepo)
	adminWebinarHandler := adminHandler.NewWebinarHandler(adminWebinarService)

	// Register admin routes (including CommentHandler)
	routes.RegisterAdminRoutes(router, adminArticleHandler, adminVideoHandler, adminAppointmentHandler, adminTestimonialHandler, adminCommentHandler, adminWebinarHandler)

	// Staff initialization
	staffArticleRepo := staffRepo.ArticleRepository{DB: db.DB}
	staffArticleService := staffService.ArticleService{Repo: &staffArticleRepo}
	staffArticleHandler := staffHandler.ArticleHandler{Service: &staffArticleService}

	staffVideoRepo := staffRepo.VideoRepository{DB: db.DB}
	staffVideoService := staffService.VideoService{Repo: &staffVideoRepo}
	staffVideoHandler := staffHandler.VideoHandler{Service: &staffVideoService}

	// Appointment initialization for Staff
	staffAppointmentRepo := staffRepo.NewAppointmentRepository(db.DB)
	staffAppointmentService := staffService.NewAppointmentService(staffAppointmentRepo)
	staffAppointmentHandler := staffHandler.NewAppointmentHandler(staffAppointmentService)

	staffTestimonialRepo := staffRepo.NewTestimonialRepository(db.DB)
	staffTestimonialService := staffService.NewTestimonialService(staffTestimonialRepo)
	staffTestimonialHandler := staffHandler.NewTestimonialHandler(staffTestimonialService)

	staffCommentRepo := staffRepo.NewCommentRepository(db.DB)
	staffCommentService := staffService.NewCommentService(staffCommentRepo)
	staffCommentHandler := staffHandler.NewCommentHandler(staffCommentService)

	staffWebinarRepo := staffRepo.NewWebinarRepository(db.DB)
	staffWebinarService := staffService.NewWebinarService(staffWebinarRepo)
	staffWebinarHandler := staffHandler.NewWebinarHandler(staffWebinarService)

	// Register staff routes
	routes.RegisterStaffRoutes(router, &staffArticleHandler, &staffVideoHandler, staffAppointmentHandler, staffTestimonialHandler, staffCommentHandler, staffWebinarHandler)

	appointmentRepo := userRepo.NewAppointmentRepository(db.DB)
	appointmentService := userService.NewAppointmentService(appointmentRepo)
	appointmentHandler := userHandler.NewAppointmentHandler(appointmentService)

	// Routing
	routes.RegisterUserRoutes(router, appointmentHandler)

	// Start the server
	log.Println("Starting server on http://localhost:8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
