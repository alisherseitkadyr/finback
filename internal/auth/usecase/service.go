package usecase

import (
	"errors"
	"finback/internal/auth/model"
	"finback/internal/auth/repository"
	"strings"
)

type RegisterRequest struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	ExperienceLevel string   `json:"experience_level"`
	Goals           []string `json:"goals"`
	PreferredTopics []string `json:"preferred_topics"`
}

type LoginRequest struct {
	Email string `json:"email"`
}

type Service struct {
	repo *repository.UserRepository
}

func NewService(repo *repository.UserRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(req RegisterRequest) (*model.User, error) {
	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.Email) == "" {
		return nil, errors.New("name and email are required")
	}
	level := strings.ToLower(strings.TrimSpace(req.ExperienceLevel))
	if level == "" {
		level = "beginner"
	}
	return s.repo.Register(req.Name, req.Email, level, req.Goals, req.PreferredTopics)
}

func (s *Service) Login(req LoginRequest) (*model.User, error) {
	return s.repo.GetByEmail(req.Email)
}

func (s *Service) Me(userID string) (*model.User, error) {
	return s.repo.GetByID(userID)
}
