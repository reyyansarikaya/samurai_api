package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"samurai_api/models"
	"samurai_api/service"
)

func ClanHandler(clanService service.ClanService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		switch r.Method {
		case http.MethodPost:
			var clan models.Clan
			if err := json.NewDecoder(r.Body).Decode(&clan); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := clanService.CreateClan(ctx, clan); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)

		case http.MethodGet:
			clans, err := clanService.GetAllClans(ctx)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(clans)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
