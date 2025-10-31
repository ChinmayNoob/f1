package repository

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/google/uuid"
)

type SeasonRepository interface {
	CreateSeason(ctx context.Context, season *model.Season) error
	GetSeason(ctx context.Context, id uuid.UUID) (*model.Season, error)
	GetAllSeasons(ctx context.Context) ([]*model.Season, error)
	UpdateSeason(ctx context.Context, season *model.Season) error
	DeleteSeason(ctx context.Context, id uuid.UUID) error
}

type seasonRepository struct {
}

func NewSeasonRepository() SeasonRepository {
	return &seasonRepository{}
}

func (r *seasonRepository) CreateSeason(ctx context.Context, season *model.Season) error {
	// TODO: Implement database logic
	return nil
}

func (r *seasonRepository) GetSeason(ctx context.Context, id uuid.UUID) (*model.Season, error) {
	// TODO: Implement database logic
	return nil, nil
}

func (r *seasonRepository) GetAllSeasons(ctx context.Context) ([]*model.Season, error) {
	// TODO: Implement database logic
	return nil, nil
}

func (r *seasonRepository) UpdateSeason(ctx context.Context, season *model.Season) error {
	// TODO: Implement database logic
	return nil
}

func (r *seasonRepository) DeleteSeason(ctx context.Context, id uuid.UUID) error {
	// TODO: Implement database logic
	return nil
}
