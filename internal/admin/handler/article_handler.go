package handler

import (
	"encoding/json"
	"fmt"
	"go-project/internal/admin/model"
	"go-project/internal/admin/service"
	"log"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	Service service.ArticleService
}

// NewArticleHandler
// ------------------
// Fungsi ini digunakan untuk menginisialisasi handler Article
// dengan menghubungkan ke layer service.
//
// Parameter:
// - service: Instance dari ArticleService yang menyediakan logika bisnis.
//
// Return:
// - Pointer ke ArticleHandler yang telah diinisialisasi.
func NewArticleHandler(service service.ArticleService) *ArticleHandler {
	return &ArticleHandler{Service: service}
}

// CreateArticle
// --------------
// Fungsi ini digunakan untuk membuat artikel baru.
//
// Parameter:
// - JSON body: Mengandung informasi artikel, termasuk title dan content.

func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
		return
	}

	if article.Title == "" || article.Content == "" {
		http.Error(w, "Title and Content are required", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateArticle(&article); err != nil {
		log.Printf("Error creating article: %v", err)
		http.Error(w, "Failed to create article: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Article created successfully"})
}

// GetArticleByID
// ---------------
// Fungsi ini digunakan untuk mengambil artikel berdasarkan ID.
//
// Parameter:
// - id (query parameter): ID artikel yang akan diambil.

func (h *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Printf("Invalid ID: %v", err)
		return
	}

	article, err := h.Service.GetArticleByID(id)
	if err != nil {
		log.Printf("Error retrieving article with id %d: %v", id, err)
		if err.Error() == fmt.Sprintf("article with id %d not found", id) {
			http.Error(w, "Article not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	json.NewEncoder(w).Encode(article)
}

// UpdateArticle
// --------------
// Fungsi ini digunakan untuk memperbarui artikel yang sudah ada.
//
// Parameter:
// - JSON body: Mengandung ID artikel serta field yang akan diperbarui (title/content).

func (h *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var article model.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if article.ID == 0 {
		http.Error(w, "Article ID is required", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateArticle(&article); err != nil {
		log.Printf("Error updating article: %v", err)
		http.Error(w, "Error updating article: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Article updated successfully"})
}

// DeleteArticle
// --------------
// Fungsi ini digunakan untuk menghapus artikel berdasarkan ID.
//
// Parameter:
// - id (query parameter): ID artikel yang akan dihapus.

func (h *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.DeleteArticle(id); err != nil {
		log.Printf("Error deleting article: %v", err)
		http.Error(w, "Error deleting article: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Article deleted successfully"})
}

// GetAllArticles
// ---------------
// Fungsi ini digunakan untuk mengambil semua artikel yang tersedia.

func (h *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := h.Service.GetAllArticles()
	if err != nil {
		log.Printf("Error retrieving articles: %v", err)
		http.Error(w, "Error retrieving articles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articles)
}
