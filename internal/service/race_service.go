package service

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/google/uuid"
)

type RaceService interface {
	CreateRace(ctx context.Context, race *model.Race) error
	GetRace(ctx context.Context, id uuid.UUID) (*model.Race, error)
	GetAllRaces(ctx context.Context) ([]*model.Race, error)
	UpdateRace(ctx context.Context, race *model.Race) error
	DeleteRace(ctx context.Context, id uuid.UUID) error
}

type raceService struct {
	repo repository.RaceRepository
}

func NewRaceService(repo repository.RaceRepository) RaceService {
	return &raceService{repo: repo}
}

func (s *raceService) CreateRace(ctx context.Context, race *model.Race) error {
	return s.repo.CreateRace(ctx, race)
}

func (s *raceService) GetRace(ctx context.Context, id uuid.UUID) (*model.Race, error) {
	return s.repo.GetRace(ctx, id)
}

func (s *raceService) GetAllRaces(ctx context.Context) ([]*model.Race, error) {
	return s.repo.GetAllRaces(ctx)
}

func (s *raceService) UpdateRace(ctx context.Context, race *model.Race) error {
	return s.repo.UpdateRace(ctx, race)
}

func (s *raceService) DeleteRace(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteRace(ctx, id)
}
