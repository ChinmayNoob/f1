package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/service"
	"github.com/ChinmayNoob/f1/internal/utils"
	"github.com/google/uuid"
)

type CircuitHandler struct {
	service service.CircuitService
	ctx     context.Context
}

func NewCircuitHandler(ctx context.Context, s service.CircuitService) *CircuitHandler {
	return &CircuitHandler{
		service: s,
		ctx:     ctx,
	}
}

func (h *CircuitHandler) GetCircuit(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, limit := utils.ParsePagination(query.Get("page"), query.Get("limit"))

	switch {
	case query.Has("name"):
		h.getByName(w, query.Get("name"), page, limit)
	case query.Has("location"):
		h.getByLocation(w, query.Get("location"), page, limit)
	case query.Has("country"):
		h.getByCountry(w, query.Get("country"), page, limit)
	case query.Has("current"):
		h.getByCurrent(w, query.Get("current") == "true")
	case query.Has("ref"):
		h.getByRef(w, query.Get("ref"))
	case query.Has("url"):
		h.getByURL(w, query.Get("url"))
	default:
		h.getAll(w, page, limit)
	}
}

func (h *CircuitHandler) GetCircuitByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		log.Printf("Error: Invalid ID format: %v", err)
		return
	}

	h.getByID(w, id)
}

func (h *CircuitHandler) CreateCircuit(w http.ResponseWriter, r *http.Request) {
	var circuit model.Circuit
	if err := json.NewDecoder(r.Body).Decode(&circuit); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("CreateCircuit error: Invalid request body: %v", err)
		return
	}

	createdCircuit, err := h.service.CreateCircuit(h.ctx, circuit)
	if err != nil {
		http.Error(w, "Failed to create circuit", http.StatusInternalServerError)
		log.Printf("CreateCircuit error: Failed to create circuit: %v", err)
		return
	}

	h.respond(w, createdCircuit)
}

func (h *CircuitHandler) UpdateCircuit(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		log.Printf("Error: Invalid ID format: %v", err)
		return
	}

	var circuit model.Circuit
	if err := json.NewDecoder(r.Body).Decode(&circuit); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("UpdateCircuit error: Invalid request body: %v", err)
		return
	}

	updatedCircuit, err := h.service.UpdateCircuit(h.ctx, id, circuit)
	if err != nil {
		http.Error(w, "Failed to update circuit", http.StatusInternalServerError)
		log.Printf("UpdateCircuit error: Failed to update circuit: %v", err)
		return
	}

	h.respond(w, updatedCircuit)
}

func (h *CircuitHandler) DeleteCircuit(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		return
	}

	idStr := parts[2]
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		log.Printf("Error: Invalid ID format: %v", err)
		return
	}

	if err := h.service.DeleteCircuit(h.ctx, id); err != nil {
		http.Error(w, "Failed to delete circuit", http.StatusInternalServerError)
		log.Printf("DeleteCircuit error: Failed to delete circuit: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ------------------------
// Private methods
// ------------------------

func (h *CircuitHandler) getAll(w http.ResponseWriter, page, limit int) {
	circuits, err := h.service.GetAllCircuits(h.ctx, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch circuits", http.StatusInternalServerError)
		log.Printf("GetAll error: %v", err)
		return
	}
	h.respond(w, circuits)
}

func (h *CircuitHandler) getByName(w http.ResponseWriter, name string, page, limit int) {
	circuits, err := h.service.GetCircuitByName(h.ctx, name, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch circuits by name", http.StatusInternalServerError)
		log.Printf("GetByName error: %v", err)
		return
	}
	h.respond(w, circuits)
}

func (h *CircuitHandler) getByLocation(w http.ResponseWriter, location string, page, limit int) {
	circuits, err := h.service.GetCircuitByLocation(h.ctx, location, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch circuits by location", http.StatusInternalServerError)
		log.Printf("GetByLocation error: %v", err)
		return
	}
	h.respond(w, circuits)
}

func (h *CircuitHandler) getByCountry(w http.ResponseWriter, country string, page, limit int) {
	circuits, err := h.service.GetCircuitByCountry(h.ctx, country, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch circuits by country", http.StatusInternalServerError)
		log.Printf("GetByCountry error: %v", err)
		return
	}
	h.respond(w, circuits)
}

func (h *CircuitHandler) getByCurrent(w http.ResponseWriter, current bool) {
	circuits, err := h.service.GetCircuitByCurrent(h.ctx, current)
	if err != nil {
		http.Error(w, "Failed to fetch circuits by current status", http.StatusInternalServerError)
		log.Printf("GetByCurrent error: %v", err)
		return
	}
	h.respond(w, circuits)
}

func (h *CircuitHandler) getByRef(w http.ResponseWriter, ref string) {
	circuit, err := h.service.GetCircuitByRef(h.ctx, ref)
	if err != nil {
		http.Error(w, "Failed to fetch circuit by ref", http.StatusInternalServerError)
		log.Printf("GetByRef error: %v", err)
		return
	}
	if circuit.ID == (uuid.UUID{}) {
		http.Error(w, "Circuit not found", http.StatusNotFound)
		return
	}
	h.respond(w, circuit)
}

func (h *CircuitHandler) getByURL(w http.ResponseWriter, url string) {
	circuit, err := h.service.GetCircuitByURL(h.ctx, url)
	if err != nil {
		http.Error(w, "Failed to fetch circuit by URL", http.StatusInternalServerError)
		log.Printf("GetByURL error: %v", err)
		return
	}
	if circuit.ID == (uuid.UUID{}) {
		http.Error(w, "Circuit not found", http.StatusNotFound)
		return
	}
	h.respond(w, circuit)
}

func (h *CircuitHandler) getByID(w http.ResponseWriter, id uuid.UUID) {
	circuit, err := h.service.GetCircuitByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Circuit not found", http.StatusNotFound)
		log.Printf("GetByID error: %v", err)
		return
	}
	h.respond(w, circuit)
}

func (h *CircuitHandler) respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
