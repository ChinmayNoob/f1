package handler

import (
	"context"
	"net/http"

	"github.com/ChinmayNoob/f1/internal/service"
)

type StandingHandler struct {
	ctx     context.Context
	service service.StandingService
}

func NewStandingHandler(ctx context.Context, service service.StandingService) *StandingHandler {
	return &StandingHandler{
		ctx:     ctx,
		service: service,
	}
}

func (h *StandingHandler) CreateDriverStanding(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusCreated)
}

func (h *StandingHandler) GetDriverStanding(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *StandingHandler) GetDriverStandingByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *StandingHandler) UpdateDriverStanding(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *StandingHandler) DeleteDriverStanding(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *StandingHandler) CreateConstructorStanding(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusCreated)
}

func (h *StandingHandler) GetConstructorStanding(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *StandingHandler) GetConstructorStandingByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *StandingHandler) UpdateConstructorStanding(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *StandingHandler) DeleteConstructorStanding(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}
