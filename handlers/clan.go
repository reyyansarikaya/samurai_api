package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"samurai_api/models"
	"samurai_api/service"
)

func ClanHandler(svc service.ClanService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			clans, err := svc.GetAllClans(r.Context())
			if err != nil {
				slog.Error("Failed to get clans", "error", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			slog.Info("Retrieved clans", "count", len(clans))
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(clans)

		case http.MethodPost:
			var clan models.Clan
			if err := json.NewDecoder(r.Body).Decode(&clan); err != nil {
				slog.Warn("Invalid clan request body", "error", err)
				http.Error(w, "invalid request body", http.StatusBadRequest)
				return
			}

			err := svc.CreateClan(r.Context(), clan)
			if err != nil {
				slog.Error("Failed to create clan", "name", clan.Name, "error", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			slog.Info("Clan created", "name", clan.Name)
			w.WriteHeader(http.StatusCreated)

		default:
			slog.Warn("Method not allowed", "method", r.Method)
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
