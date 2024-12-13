package repository

import (
	"database/sql"
	"errors"
	"go-project/internal/staff/model"
)

type WebinarRepository interface {
	GetAllWebinars() ([]model.Webinar, error)
	GetWebinarByID(id int) (*model.Webinar, error)
}

type webinarRepository struct {
	db *sql.DB
}

func NewWebinarRepository(db *sql.DB) WebinarRepository {
	return &webinarRepository{db: db}
}

func (r *webinarRepository) GetAllWebinars() ([]model.Webinar, error) {
	query := `SELECT id, title, description, link_meet, host_id, created_at, updated_at FROM webinars`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, errors.New("failed to fetch webinars: " + err.Error())
	}
	defer rows.Close()

	var webinars []model.Webinar
	for rows.Next() {
		var webinar model.Webinar
		if err := rows.Scan(&webinar.ID, &webinar.Title, &webinar.Description, &webinar.LinkMeet, &webinar.HostID, &webinar.CreatedAt, &webinar.UpdatedAt); err != nil {
			return nil, errors.New("failed to scan webinar: " + err.Error())
		}
		webinars = append(webinars, webinar)
	}

	return webinars, nil
}

func (r *webinarRepository) GetWebinarByID(id int) (*model.Webinar, error) {
	query := `SELECT id, title, description, link_meet, host_id, created_at, updated_at FROM webinars WHERE id = $1`
	var webinar model.Webinar
	err := r.db.QueryRow(query, id).Scan(&webinar.ID, &webinar.Title, &webinar.Description, &webinar.LinkMeet, &webinar.HostID, &webinar.CreatedAt, &webinar.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("webinar not found")
	} else if err != nil {
		return nil, errors.New("failed to fetch webinar: " + err.Error())
	}
	return &webinar, nil
}
