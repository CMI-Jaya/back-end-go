package repository

import (
	"database/sql"
	"go-project/internal/staff/model"
)

type VideoRepository struct {
	DB *sql.DB
}

func (r *VideoRepository) SaveVideo(video model.Video) error {
	query := `
    INSERT INTO videos (
        title, description, link_video, category_id, status, author_id, meta_title, meta_description, created_at, updated_at
    ) VALUES (
        $1, $2, $3, $4, 'pending approval', $5, $6, $7, NOW(), NOW()
    ) RETURNING id`
	err := r.DB.QueryRow(query, video.Title, video.Description, video.LinkVideo, video.CategoryID, video.AuthorID, video.MetaTitle, video.MetaDescription).Scan(&video.ID)
	return err
}

func (r *VideoRepository) GetVideoByID(id int) (*model.Video, error) {
	query := `SELECT id, title, description, link_video, category_id, status, author_id, meta_title, meta_description, created_at, updated_at FROM videos WHERE id = $1`
	row := r.DB.QueryRow(query, id)

	var video model.Video
	err := row.Scan(
		&video.ID,
		&video.Title,
		&video.Description,
		&video.LinkVideo,
		&video.CategoryID,
		&video.Status,
		&video.AuthorID,
		&video.MetaTitle,
		&video.MetaDescription,
		&video.CreatedAt,
		&video.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &video, nil
}

func (r *VideoRepository) GetAllVideos() ([]model.Video, error) {
	query := `SELECT id, title, description, link_video, category_id, status, author_id, meta_title, meta_description, created_at, updated_at FROM videos`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var videos []model.Video
	for rows.Next() {
		var video model.Video
		if err := rows.Scan(
			&video.ID,
			&video.Title,
			&video.Description,
			&video.LinkVideo,
			&video.CategoryID,
			&video.Status,
			&video.AuthorID,
			&video.MetaTitle,
			&video.MetaDescription,
			&video.CreatedAt,
			&video.UpdatedAt,
		); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	return videos, nil
}
