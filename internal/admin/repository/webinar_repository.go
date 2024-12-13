package repository

import (
	"database/sql"
	"errors"
	"go-project/internal/admin/model"
)

type WebinarRepository interface {
	CreateWebinar(webinar *model.Webinar) error
}

type webinarRepository struct {
	db *sql.DB
}

func NewWebinarRepository(db *sql.DB) WebinarRepository {
	return &webinarRepository{db: db}
}

func (r *webinarRepository) CreateWebinar(webinar *model.Webinar) error {
	query := `INSERT INTO webinars (title, description, link_meet, host_id, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`
	err := r.db.QueryRow(query, webinar.Title, webinar.Description, webinar.LinkMeet, webinar.HostID).Scan(&webinar.ID)
	if err != nil {
		return errors.New("failed to create webinar: " + err.Error())
	}
	return nil
}
