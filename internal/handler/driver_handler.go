package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/service"
	"github.com/ChinmayNoob/f1/internal/utils"
	"github.com/google/uuid"
)

type DriverHandler struct {
	service service.DriverService
	ctx     context.Context
}

func NewDriverHandler(ctx context.Context, s service.DriverService) *DriverHandler {
	return &DriverHandler{
		service: s,
		ctx:     ctx,
	}
}

func (h *DriverHandler) GetDriver(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, limit := utils.ParsePagination(query.Get("page"), query.Get("limit"))

	switch {
	case query.Has("firstName"):
		h.getByFirstName(w, query.Get("firstName"), page, limit)
	case query.Has("lastName"):
		h.getByLastName(w, query.Get("lastName"), page, limit)
	case query.Has("team"):
		h.getByTeam(w, query.Get("team"), page, limit)
	case query.Has("nationality"):
		h.getByNationality(w, query.Get("nationality"), page, limit)
	case query.Has("status"):
		h.getByStatus(w, query.Get("status"), page, limit)
	case query.Has("ref"):
		h.getByRef(w, query.Get("ref"))
	case query.Has("code"):
		h.getByCode(w, query.Get("code"))
	case query.Has("number"):
		number, err := strconv.Atoi(query.Get("number"))
		if err != nil {
			http.Error(w, "Invalid number format", http.StatusBadRequest)
			return
		}
		h.getByNumber(w, number)
	case query.Has("url"):
		h.getByURL(w, query.Get("url"))
	default:
		h.getAll(w, page, limit)
	}
}

func (h *DriverHandler) GetDriverByID(w http.ResponseWriter, r *http.Request) {
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

func (h *DriverHandler) CreateDriver(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("CreateDriver error: Invalid request body: %v", err)
		return
	}

	createdDriver, err := h.service.CreateDriver(h.ctx, driver)
	if err != nil {
		if errors.Is(err, errors.New("constructor not found")) {
			http.Error(w, "Constructor not found", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to create driver", http.StatusInternalServerError)
		}
		log.Printf("CreateDriver error: Failed to create driver: %v", err)
		return
	}

	h.respond(w, createdDriver)
}

func (h *DriverHandler) UpdateDriver(w http.ResponseWriter, r *http.Request) {
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

	var driver model.Driver
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("UpdateDriver error: Invalid request body: %v", err)
		return
	}

	updatedDriver, err := h.service.UpdateDriver(h.ctx, id, driver)
	if err != nil {
		if errors.Is(err, errors.New("constructor not found")) {
			http.Error(w, "Constructor not found", http.StatusBadRequest)
		} else {
			http.Error(w, "Failed to update driver", http.StatusInternalServerError)
		}
		log.Printf("UpdateDriver error: Failed to update driver: %v", err)
		return
	}

	h.respond(w, updatedDriver)
}

func (h *DriverHandler) DeleteDriver(w http.ResponseWriter, r *http.Request) {
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

	if err := h.service.DeleteDriver(h.ctx, id); err != nil {
		http.Error(w, "Failed to delete driver", http.StatusInternalServerError)
		log.Printf("DeleteDriver error: Failed to delete driver: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ------------------------
// Private methods
// ------------------------

func (h *DriverHandler) getAll(w http.ResponseWriter, page, limit int) {
	drivers, err := h.service.GetAllDrivers(h.ctx, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch drivers", http.StatusInternalServerError)
		log.Printf("GetAll error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByFirstName(w http.ResponseWriter, firstName string, page, limit int) {
	drivers, err := h.service.GetDriverByFirstName(h.ctx, firstName, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch drivers by first name", http.StatusInternalServerError)
		log.Printf("GetByFirstName error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByLastName(w http.ResponseWriter, lastName string, page, limit int) {
	drivers, err := h.service.GetDriverByLastName(h.ctx, lastName, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch drivers by last name", http.StatusInternalServerError)
		log.Printf("GetByLastName error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByTeam(w http.ResponseWriter, team string, page, limit int) {
	drivers, err := h.service.GetDriverByTeam(h.ctx, team, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch drivers by team", http.StatusInternalServerError)
		log.Printf("GetByTeam error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByNationality(w http.ResponseWriter, nationality string, page, limit int) {
	drivers, err := h.service.GetDriverByNationality(h.ctx, nationality, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch drivers by nationality", http.StatusInternalServerError)
		log.Printf("GetByNationality error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByStatus(w http.ResponseWriter, status string, page, limit int) {
	drivers, err := h.service.GetDriverByStatus(h.ctx, status, page, limit)
	if err != nil {
		http.Error(w, "Failed to fetch drivers by status", http.StatusInternalServerError)
		log.Printf("GetByStatus error: %v", err)
		return
	}
	h.respond(w, drivers)
}

func (h *DriverHandler) getByRef(w http.ResponseWriter, ref string) {
	driver, err := h.service.GetDriverByRef(h.ctx, ref)
	if err != nil {
		http.Error(w, "Failed to fetch driver by ref", http.StatusInternalServerError)
		log.Printf("GetByRef error: %v", err)
		return
	}
	if driver.ID == (uuid.UUID{}) {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}
	h.respond(w, driver)
}

func (h *DriverHandler) getByCode(w http.ResponseWriter, code string) {
	driver, err := h.service.GetDriverByCode(h.ctx, code)
	if err != nil {
		http.Error(w, "Failed to fetch driver by code", http.StatusInternalServerError)
		log.Printf("GetByCode error: %v", err)
		return
	}
	if driver.ID == (uuid.UUID{}) {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}
	h.respond(w, driver)
}

func (h *DriverHandler) getByNumber(w http.ResponseWriter, number int) {
	driver, err := h.service.GetDriverByNumber(h.ctx, number)
	if err != nil {
		http.Error(w, "Failed to fetch driver by number", http.StatusInternalServerError)
		log.Printf("GetByNumber error: %v", err)
		return
	}
	if driver.ID == (uuid.UUID{}) {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}
	h.respond(w, driver)
}

func (h *DriverHandler) getByURL(w http.ResponseWriter, url string) {
	driver, err := h.service.GetDriverByURL(h.ctx, url)
	if err != nil {
		http.Error(w, "Failed to fetch driver by URL", http.StatusInternalServerError)
		log.Printf("GetByURL error: %v", err)
		return
	}
	if driver.ID == (uuid.UUID{}) {
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}
	h.respond(w, driver)
}

func (h *DriverHandler) getByID(w http.ResponseWriter, id uuid.UUID) {
	driver, err := h.service.GetDriverByID(h.ctx, id)
	if err != nil {
		http.Error(w, "Driver not found", http.StatusNotFound)
		log.Printf("GetByID error: %v", err)
		return
	}
	h.respond(w, driver)
}

func (h *DriverHandler) respond(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}
