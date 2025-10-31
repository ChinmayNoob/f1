package service

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/google/uuid"
)

type ResultService interface {
	CreateResult(ctx context.Context, result *model.Result) error
	GetResult(ctx context.Context, id uuid.UUID) (*model.Result, error)
	GetAllResults(ctx context.Context) ([]*model.Result, error)
	UpdateResult(ctx context.Context, result *model.Result) error
	DeleteResult(ctx context.Context, id uuid.UUID) error
}

type resultService struct {
	repo repository.ResultRepository
}

func NewResultService(repo repository.ResultRepository) ResultService {
	return &resultService{repo: repo}
}

func (s *resultService) CreateResult(ctx context.Context, result *model.Result) error {
	return s.repo.CreateResult(ctx, result)
}

func (s *resultService) GetResult(ctx context.Context, id uuid.UUID) (*model.Result, error) {
	return s.repo.GetResult(ctx, id)
}

func (s *resultService) GetAllResults(ctx context.Context) ([]*model.Result, error) {
	return s.repo.GetAllResults(ctx)
}

func (s *resultService) UpdateResult(ctx context.Context, result *model.Result) error {
	return s.repo.UpdateResult(ctx, result)
}

func (s *resultService) DeleteResult(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteResult(ctx, id)
}
