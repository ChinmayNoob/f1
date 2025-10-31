package handler

import (
	"context"
	"net/http"

	"github.com/ChinmayNoob/f1/internal/service"
)

type SeasonHandler struct {
	ctx     context.Context
	service service.SeasonService
}

func NewSeasonHandler(ctx context.Context, service service.SeasonService) *SeasonHandler {
	return &SeasonHandler{
		ctx:     ctx,
		service: service,
	}
}

func (h *SeasonHandler) CreateSeason(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusCreated)
}

func (h *SeasonHandler) GetSeason(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *SeasonHandler) GetSeasonByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *SeasonHandler) UpdateSeason(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *SeasonHandler) DeleteSeason(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}
