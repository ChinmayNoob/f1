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

type ConstructorHandler struct {
	service service.ConstructorService
	ctx     context.Context
}

func NewConstructorHandler(ctx context.Context, s service.ConstructorService) *ConstructorHandler {
	return &ConstructorHandler{
		service: s,
		ctx:     ctx,
	}
}

func (h *ConstructorHandler) GetConstructor(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, limit := utils.ParsePagination(query.Get("page"), query.Get("limit"))

	switch {
	case query.Has("name"):
		h.getByName(w, query.Get("name"), page, limit)
	case query.Has("nationality"):
		h.getByNationality(w, query.Get("nationality"), page, limit)
	case query.Has("ref"):
		h.getByRef(w, query.Get("ref"))
	default:
		h.getAll(w, page, limit)
	}
}

func (h *ConstructorHandler) GetConstructorByID(w http.ResponseWriter, r *http.Request) {
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

func (h *ConstructorHandler) CreateConstructor(w http.ResponseWriter, r *http.Request) {
	var constructor model.Constructor
	if err := json.NewDecoder(r.Body).Decode(&constructor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("CreateConstructor error: Invalid request body: %v", err)
		return
	}

	createdConstructor, err := h.service.CreateConstructor(h.ctx, constructor)
	if err != nil {
		http.Error(w, "Failed to create constructor", http.StatusInternalServerError)
		log.Printf("CreateConstructor error: Failed to create constructor: %v", err)
		return
	}

	h.respond(w, createdConstructor)
}

func (h *ConstructorHandler) UpdateConstructor(w http.ResponseWriter, r *http.Request) {
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

	var constructor model.Constructor
	if err := json.NewDecoder(r.Body).Decode(&constructor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("UpdateConstructor error: Invalid request body: %v", err)
		return
	}

	updatedConstructor, err := h.service.UpdateConstructor(h.ctx, id, constructor)
	if err != nil {
		http.Error(w, "Failed to update constructor", http.StatusInternalServerError)
		log.Printf("UpdateConstructor error: Failed to update constructor: %v", err)
		return
	}

	h.respond(w, updatedConstructor)
}

func (h *ConstructorHandler) DeleteConstructor(w http.ResponseWriter, r *http.Request) {
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

	if err := h.service.DeleteConstructor(h.ctx, id); err != nil {
		http.Error(w, "Failed to delete constructor", http.StatusInternalServerError)
		log.Printf("DeleteConstructor error: Failed to delete constructor: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ------------------------
// Private methods
// ------------------------

func (h *ConstructorHandler) getAll(w http.ResponseWriter, page, limit int) {
	constructors, err := h.service.GetAllConstructors(h.ctx, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch constructors", http.StatusInternalServerError)
		log.Printf("GetAll error: %v", err)
		return
	}
	h.respond(w, constructors)
}

func (h *ConstructorHandler) getByName(w http.ResponseWriter, name string, page, limit int) {
	constructor, err := h.service.GetConstructorByName(h.ctx, name, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch constructor by name", http.StatusInternalServerError)
		log.Printf("GetByName error: %v", err)
		return
	}
	h.respond(w, constructor)
}

func (h *ConstructorHandler) getByNationality(w http.ResponseWriter, nationality string, page, limit int) {
	constructors, err := h.service.GetConstructorByNationality(h.ctx, nationality, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch constructors by nationality", http.StatusInternalServerError)
		log.Printf("GetByNationality error: %v", err)
		return
	}
	h.respond(w, constructors)
}

func (h *ConstructorHandler) getByRef(w http.ResponseWriter, ref string) {
	constructor, err := h.service.GetConstructorByRef(h.ctx, ref)
	if err != nil {
		http.Error(w, "Failed to fetch constructor by ref", http.StatusInternalServerError)
		log.Printf("GetByRef error: %v", err)
		return
	}

	if constructor.ID == (uuid.UUID{}) {
		http.Error(w, "Constructor not found", http.StatusNotFound)
		return
	}

	h.respond(w, constructor)
}

func (h *ConstructorHandler) getByID(w http.ResponseWriter, id uuid.UUID) {
	constructor, err := h.service.GetConstructorByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Constructor not found", http.StatusNotFound)
		log.Printf("GetByID error: %v", err)
		return
	}
	h.respond(w, constructor)
}

func (h *ConstructorHandler) respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
