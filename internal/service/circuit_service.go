package service

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/google/uuid"
)

type CircuitService interface {
	CreateCircuit(ctx context.Context, circuit model.Circuit) (model.Circuit, error)
	GetAllCircuits(ctx context.Context, page, limit int) ([]model.Circuit, error)
	GetCircuitByID(ctx context.Context, id uuid.UUID) (model.Circuit, error)
	GetCircuitByRef(ctx context.Context, ref string) (model.Circuit, error)
	GetCircuitByName(ctx context.Context, name string, page, limit int) ([]model.Circuit, error)
	GetCircuitByLocation(ctx context.Context, location string, page, limit int) ([]model.Circuit, error)
	GetCircuitByCountry(ctx context.Context, country string, page, limit int) ([]model.Circuit, error)
	GetCircuitByCurrent(ctx context.Context, current bool) ([]model.Circuit, error)
	GetCircuitByURL(ctx context.Context, url string) (model.Circuit, error)
	UpdateCircuit(ctx context.Context, id uuid.UUID, circuit model.Circuit) (model.Circuit, error)
	DeleteCircuit(ctx context.Context, id uuid.UUID) error
}

type circuitService struct {
	repo repository.CircuitRepository
}

func NewCircuitService(r repository.CircuitRepository) CircuitService {
	return &circuitService{repo: r}
}

func (s *circuitService) CreateCircuit(ctx context.Context, circuit model.Circuit) (model.Circuit, error) {
	circuit.ID = uuid.New()
	return s.repo.CreateCircuit(ctx, circuit)
}

func (s *circuitService) GetAllCircuits(ctx context.Context, page, limit int) ([]model.Circuit, error) {
	return s.repo.GetAllCircuits(ctx, page, limit)
}

func (s *circuitService) GetCircuitByID(ctx context.Context, id uuid.UUID) (model.Circuit, error) {
	return s.repo.GetCircuitByID(ctx, id)
}

func (s *circuitService) GetCircuitByRef(ctx context.Context, ref string) (model.Circuit, error) {
	return s.repo.GetCircuitByRef(ctx, ref)
}

func (s *circuitService) GetCircuitByName(ctx context.Context, name string, page, limit int) ([]model.Circuit, error) {
	return s.repo.GetCircuitByName(ctx, name, page, limit)
}

func (s *circuitService) GetCircuitByLocation(ctx context.Context, location string, page, limit int) ([]model.Circuit, error) {
	return s.repo.GetCircuitByLocation(ctx, location, page, limit)
}

func (s *circuitService) GetCircuitByCountry(ctx context.Context, country string, page, limit int) ([]model.Circuit, error) {
	return s.repo.GetCircuitByCountry(ctx, country, page, limit)
}

func (s *circuitService) GetCircuitByCurrent(ctx context.Context, current bool) ([]model.Circuit, error) {
	return s.repo.GetCircuitByCurrent(ctx, current)
}

func (s *circuitService) GetCircuitByURL(ctx context.Context, url string) (model.Circuit, error) {
	return s.repo.GetCircuitByURL(ctx, url)
}

func (s *circuitService) UpdateCircuit(ctx context.Context, id uuid.UUID, circuit model.Circuit) (model.Circuit, error) {
	return s.repo.UpdateCircuit(ctx, id, circuit)
}

func (s *circuitService) DeleteCircuit(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCircuit(ctx, id)
}
