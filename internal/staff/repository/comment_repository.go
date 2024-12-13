package repository

import (
	"database/sql"
	"go-project/internal/staff/model"
)

type CommentRepository interface {
	GetAllComments() ([]model.Comment, error)
	GetCommentByID(commentID int) (*model.Comment, error)
	DeleteComment(commentID int) error
	CreateComment(comment *model.Comment) (*model.Comment, error)
}

type commentRepository struct {
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) GetAllComments() ([]model.Comment, error) {
	query := `SELECT id, article_id, username, email, comment, parent_id, status, created_at, updated_at FROM comments`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []model.Comment
	for rows.Next() {
		var comment model.Comment
		var parentID sql.NullInt32
		if err := rows.Scan(&comment.ID, &comment.ArticleID, &comment.Username, &comment.Email,
			&comment.Comment, &parentID, &comment.Status, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		if parentID.Valid {
			id := int(parentID.Int32)
			comment.ParentID = &id
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *commentRepository) GetCommentByID(commentID int) (*model.Comment, error) {
	query := `SELECT id, article_id, username, email, comment, parent_id, status, created_at, updated_at FROM comments WHERE id = $1`
	row := r.db.QueryRow(query, commentID)

	var comment model.Comment
	var parentID sql.NullInt32
	if err := row.Scan(&comment.ID, &comment.ArticleID, &comment.Username, &comment.Email,
		&comment.Comment, &parentID, &comment.Status, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}
	if parentID.Valid {
		id := int(parentID.Int32)
		comment.ParentID = &id
	}
	return &comment, nil
}

func (r *commentRepository) DeleteComment(commentID int) error {
	query := `DELETE FROM comments WHERE id = $1`
	_, err := r.db.Exec(query, commentID)
	return err
}

func (r *commentRepository) CreateComment(comment *model.Comment) (*model.Comment, error) {
	query := `INSERT INTO comments (article_id, username, email, comment, parent_id, status, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.db.QueryRow(query, comment.ArticleID, comment.Username, comment.Email, comment.Comment, comment.ParentID, comment.Status).
		Scan(&comment.ID, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
