package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"samurai_api/models"
	"samurai_api/service"
)

func SamuraiHandler(samuraiService service.SamuraiService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		switch r.Method {
		case http.MethodPost:
			var s models.Samurai
			if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := samuraiService.CreateSamurai(ctx, s); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)

		case http.MethodGet:
			samurais, err := samuraiService.GetAllSamurais(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(samurais)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
