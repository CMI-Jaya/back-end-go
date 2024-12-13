package handler

import (
	"encoding/json"
	"go-project/internal/staff/model"
	"go-project/internal/staff/service"
	"net/http"
	"strconv"
)

// ArticleHandler handles HTTP requests for articles
type ArticleHandler struct {
	Service *service.ArticleService
}

// UploadArticle handles the article upload request
func (h *ArticleHandler) UploadArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Buat artikel melalui service
	if err := h.Service.CreateArticle(article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Article uploaded successfully and notification sent!"))
}

// GetArticleByID retrieves an article by ID
func (h *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	// Extract the article ID from the URL query
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Call the service to get the article by ID
	article, err := h.Service.GetArticleByID(id)
	if err != nil {
		http.Error(w, "Article not found", http.StatusNotFound)
		return
	}

	// Respond with the article data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// GetAllArticles retrieves all articles
func (h *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	// Call the service to get all articles
	articles, err := h.Service.GetAllArticles()
	if err != nil {
		http.Error(w, "Error fetching articles", http.StatusInternalServerError)
		return
	}

	// Respond with the list of articles
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
