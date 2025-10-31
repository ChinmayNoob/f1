package service

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/google/uuid"
)

type SeasonService interface {
	CreateSeason(ctx context.Context, season *model.Season) error
	GetSeason(ctx context.Context, id uuid.UUID) (*model.Season, error)
	GetAllSeasons(ctx context.Context) ([]*model.Season, error)
	UpdateSeason(ctx context.Context, season *model.Season) error
	DeleteSeason(ctx context.Context, id uuid.UUID) error
}

type seasonService struct {
	repo repository.SeasonRepository
}

func NewSeasonService(repo repository.SeasonRepository) SeasonService {
	return &seasonService{repo: repo}
}

func (s *seasonService) CreateSeason(ctx context.Context, season *model.Season) error {
	return s.repo.CreateSeason(ctx, season)
}

func (s *seasonService) GetSeason(ctx context.Context, id uuid.UUID) (*model.Season, error) {
	return s.repo.GetSeason(ctx, id)
}

func (s *seasonService) GetAllSeasons(ctx context.Context) ([]*model.Season, error) {
	return s.repo.GetAllSeasons(ctx)
}

func (s *seasonService) UpdateSeason(ctx context.Context, season *model.Season) error {
	return s.repo.UpdateSeason(ctx, season)
}

func (s *seasonService) DeleteSeason(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteSeason(ctx, id)
}
