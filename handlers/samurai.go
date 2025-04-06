package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"samurai_api/models"
	"samurai_api/service"

	"github.com/go-chi/chi"
)

type SamuraiHandler struct {
	svc service.SamuraiService
}

func NewSamuraiHandler(svc service.SamuraiService) *SamuraiHandler {
	return &SamuraiHandler{svc: svc}
}

func (h *SamuraiHandler) GetAllSamurais(w http.ResponseWriter, r *http.Request) {
	samurais, err := h.svc.GetAllSamurais(r.Context())
	if err != nil {
		slog.Error("Failed to get samurais", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("Retrieved samurais", "count", len(samurais))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(samurais)
}

func (h *SamuraiHandler) CreateSamurai(w http.ResponseWriter, r *http.Request) {
	var samurai models.Samurai
	if err := json.NewDecoder(r.Body).Decode(&samurai); err != nil {
		slog.Warn("Invalid request body", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := h.svc.CreateSamurai(r.Context(), &samurai)
	if err != nil {
		slog.Error("Failed to create samurai", "name", samurai.Name, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	slog.Info("Samurai created", "name", samurai.Name, "clan", samurai.ClanID)
	w.WriteHeader(http.StatusCreated)
}

func (h *SamuraiHandler) Attack(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req models.AttackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Warn("Invalid attack request body", "error", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := h.svc.Attack(id, req)
	if err != nil {
		slog.Error("Attack failed", "samurai", id, "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.svc.PublishAttackEvent(result)

	slog.Info("Attack processed", "samurai", id, "result", result.Result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
