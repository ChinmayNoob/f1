package repository

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/google/uuid"
)

type RaceRepository interface {
	CreateRace(ctx context.Context, race *model.Race) error
	GetRace(ctx context.Context, id uuid.UUID) (*model.Race, error)
	GetAllRaces(ctx context.Context) ([]*model.Race, error)
	UpdateRace(ctx context.Context, race *model.Race) error
	DeleteRace(ctx context.Context, id uuid.UUID) error
}

type raceRepository struct {
}

func NewRaceRepository() RaceRepository {
	return &raceRepository{}
}

func (r *raceRepository) CreateRace(ctx context.Context, race *model.Race) error {
	// TODO: Implement database logic
	return nil
}

func (r *raceRepository) GetRace(ctx context.Context, id uuid.UUID) (*model.Race, error) {
	// TODO: Implement database logic
	return nil, nil
}

func (r *raceRepository) GetAllRaces(ctx context.Context) ([]*model.Race, error) {
	// TODO: Implement database logic
	return nil, nil
}

func (r *raceRepository) UpdateRace(ctx context.Context, race *model.Race) error {
	// TODO: Implement database logic
	return nil
}

func (r *raceRepository) DeleteRace(ctx context.Context, id uuid.UUID) error {
	// TODO: Implement database logic
	return nil
}
