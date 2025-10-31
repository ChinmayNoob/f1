package handler

import (
	"context"
	"net/http"

	"github.com/ChinmayNoob/f1/internal/service"
)

type RaceHandler struct {
	ctx     context.Context
	service service.RaceService
}

func NewRaceHandler(ctx context.Context, service service.RaceService) *RaceHandler {
	return &RaceHandler{
		ctx:     ctx,
		service: service,
	}
}

func (h *RaceHandler) CreateRace(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusCreated)
}

func (h *RaceHandler) GetRace(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *RaceHandler) GetRaceByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *RaceHandler) UpdateRace(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *RaceHandler) DeleteRace(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}
