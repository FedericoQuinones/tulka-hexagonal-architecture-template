package services

import (
	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/domain/models"
	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/infra/repositories"
)

type UserService struct {
	Repo repositories.Repository
}

func NewUserService(repo repositories.Repository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	return s.Repo.GetUserByID(id)
}
