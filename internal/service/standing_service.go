package service

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/google/uuid"
)

type StandingService interface {
	CreateDriverStanding(ctx context.Context, standing *model.DriverStanding) error
	GetDriverStanding(ctx context.Context, id uuid.UUID) (*model.DriverStanding, error)
	GetAllDriverStandings(ctx context.Context) ([]*model.DriverStanding, error)
	UpdateDriverStanding(ctx context.Context, standing *model.DriverStanding) error
	DeleteDriverStanding(ctx context.Context, id uuid.UUID) error

	CreateConstructorStanding(ctx context.Context, standing *model.ConstructorStanding) error
	GetConstructorStanding(ctx context.Context, id uuid.UUID) (*model.ConstructorStanding, error)
	GetAllConstructorStandings(ctx context.Context) ([]*model.ConstructorStanding, error)
	UpdateConstructorStanding(ctx context.Context, standing *model.ConstructorStanding) error
	DeleteConstructorStanding(ctx context.Context, id uuid.UUID) error
}

type standingService struct {
	repo repository.StandingRepository
}

func NewStandingService(repo repository.StandingRepository) StandingService {
	return &standingService{repo: repo}
}

func (s *standingService) CreateDriverStanding(ctx context.Context, standing *model.DriverStanding) error {
	return s.repo.CreateDriverStanding(ctx, standing)
}

func (s *standingService) GetDriverStanding(ctx context.Context, id uuid.UUID) (*model.DriverStanding, error) {
	return s.repo.GetDriverStanding(ctx, id)
}

func (s *standingService) GetAllDriverStandings(ctx context.Context) ([]*model.DriverStanding, error) {
	return s.repo.GetAllDriverStandings(ctx)
}

func (s *standingService) UpdateDriverStanding(ctx context.Context, standing *model.DriverStanding) error {
	return s.repo.UpdateDriverStanding(ctx, standing)
}

func (s *standingService) DeleteDriverStanding(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDriverStanding(ctx, id)
}

func (s *standingService) CreateConstructorStanding(ctx context.Context, standing *model.ConstructorStanding) error {
	return s.repo.CreateConstructorStanding(ctx, standing)
}

func (s *standingService) GetConstructorStanding(ctx context.Context, id uuid.UUID) (*model.ConstructorStanding, error) {
	return s.repo.GetConstructorStanding(ctx, id)
}

func (s *standingService) GetAllConstructorStandings(ctx context.Context) ([]*model.ConstructorStanding, error) {
	return s.repo.GetAllConstructorStandings(ctx)
}

func (s *standingService) UpdateConstructorStanding(ctx context.Context, standing *model.ConstructorStanding) error {
	return s.repo.UpdateConstructorStanding(ctx, standing)
}

func (s *standingService) DeleteConstructorStanding(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteConstructorStanding(ctx, id)
}
