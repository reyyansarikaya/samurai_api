package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"samurai_api/models"
	"samurai_api/service"
)

func SamuraiHandler(svc service.SamuraiService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			samurais, err := svc.GetAllSamurais(r.Context())
			if err != nil {
				slog.Error("Failed to get samurais", "error", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			slog.Info("Retrieved samurais", "count", len(samurais))
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(samurais)

		case http.MethodPost:
			var samurai models.Samurai
			if err := json.NewDecoder(r.Body).Decode(&samurai); err != nil {
				slog.Warn("Invalid request body", "error", err)
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			err := svc.CreateSamurai(r.Context(), &samurai)
			if err != nil {
				slog.Error("Failed to create samurai", "name", samurai.Name, "error", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			slog.Info("Samurai created", "name", samurai.Name, "clan", samurai.ClanID)
			w.WriteHeader(http.StatusCreated)

		default:
			slog.Warn("Method not allowed", "method", r.Method)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
