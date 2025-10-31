package handler

import (
	"context"
	"net/http"

	"github.com/ChinmayNoob/f1/internal/service"
)

type ResultHandler struct {
	ctx     context.Context
	service service.ResultService
}

func NewResultHandler(ctx context.Context, service service.ResultService) *ResultHandler {
	return &ResultHandler{
		ctx:     ctx,
		service: service,
	}
}

func (h *ResultHandler) CreateResult(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusCreated)
}

func (h *ResultHandler) GetResult(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *ResultHandler) GetResultByID(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *ResultHandler) UpdateResult(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}

func (h *ResultHandler) DeleteResult(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement handler logic
	w.WriteHeader(http.StatusOK)
}
