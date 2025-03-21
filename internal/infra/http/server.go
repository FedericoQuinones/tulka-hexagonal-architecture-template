package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/domain/services"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	Addr        string
	Router      chi.Router
	UserService *services.UserService
}

func (s *Server) Start() error {
	fmt.Println("🚀 Server running at", s.Addr)

	s.Router.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		// Convertir el id de string a int
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		user, err := s.UserService.GetUser(id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "User: %+v", user)
	})

	return http.ListenAndServe(s.Addr, s.Router)
}

func NewServer(addr string, userService *services.UserService) *Server {
	r := chi.NewRouter()

	return &Server{
		Addr:        addr,
		Router:      r,
		UserService: userService,
	}
}
