package main

import (
	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/domain/services"
	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/infra/http"

	"github.com/google/wire"
)

func CreateServer() (*http.Server, error) {
	wire.Build(
		services.NewUserService,
		http.NewServer,
	)
	return nil, nil
}
