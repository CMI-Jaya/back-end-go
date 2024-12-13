package service

import (
	"errors"
	"go-project/internal/admin/model"
	"go-project/internal/admin/repository"
)

type WebinarService interface {
	CreateWebinar(webinar *model.Webinar) error
}

type webinarService struct {
	repo repository.WebinarRepository
}

func NewWebinarService(repo repository.WebinarRepository) WebinarService {
	return &webinarService{repo: repo}
}

func (s *webinarService) CreateWebinar(webinar *model.Webinar) error {
	if webinar.Title == "" || webinar.Description == "" || webinar.HostID == 0 {
		return errors.New("invalid webinar data")
	}
	return s.repo.CreateWebinar(webinar)
}
