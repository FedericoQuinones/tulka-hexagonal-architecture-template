package interfaces

import "github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/domain/models"

type Repository interface {
	GetUserByID(id int) (*models.User, error)
}
