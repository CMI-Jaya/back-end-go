package service

import (
	"go-project/internal/staff/model"
	"go-project/internal/staff/repository"
)

type WebinarService interface {
	GetAllWebinars() ([]model.Webinar, error)
	GetWebinarByID(id int) (*model.Webinar, error)
}

type webinarService struct {
	repo repository.WebinarRepository
}

func NewWebinarService(repo repository.WebinarRepository) WebinarService {
	return &webinarService{repo: repo}
}

func (s *webinarService) GetAllWebinars() ([]model.Webinar, error) {
	return s.repo.GetAllWebinars()
}

func (s *webinarService) GetWebinarByID(id int) (*model.Webinar, error) {
	return s.repo.GetWebinarByID(id)
}
