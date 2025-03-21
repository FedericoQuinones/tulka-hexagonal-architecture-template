package main

import (
	"log"

	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/domain/services"
	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/infra/http"
	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/infra/repositories"
)

func main() {
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)

	server := http.NewServer(":8080", userService)

	err := server.Start()
	if err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
