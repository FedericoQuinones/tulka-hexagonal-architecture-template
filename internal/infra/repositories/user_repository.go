package repositories

import "github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/domain/models"

type Repository interface {
	GetUserByID(id int) (*models.User, error)
}

type UserRepository struct {
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	return &models.User{ID: id, Name: "Federico"}, nil
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}
