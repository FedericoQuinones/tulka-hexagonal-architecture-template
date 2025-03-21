package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/FedericoQuinones/tulka-hexagotal-architecture-template/internal/domain/models"
)

func ErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		if r.Context().Err() != nil {
			log.Println("Error:", r.Context().Err())

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)

			errorResponse := models.ErrorResponse{Message: "Something went wrong."}
			json.NewEncoder(w).Encode(errorResponse)
		}
	})
}
