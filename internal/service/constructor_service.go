package service

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/google/uuid"
)

type ConstructorService interface {
	CreateConstructor(ctx context.Context, constructor model.Constructor) (model.Constructor, error)
	GetAllConstructors(ctx context.Context, page, limit int) ([]model.Constructor, error)
	GetConstructorByName(ctx context.Context, name string, page, limit int) ([]model.Constructor, error)
	GetConstructorByID(ctx context.Context, id uuid.UUID) (model.Constructor, error)
	GetConstructorByNationality(ctx context.Context, nationality string, page, limit int) ([]model.Constructor, error)
	GetConstructorByRef(ctx context.Context, ref string) (model.Constructor, error)
	UpdateConstructor(ctx context.Context, id uuid.UUID, constructor model.Constructor) (model.Constructor, error)
	DeleteConstructor(ctx context.Context, id uuid.UUID) error
}

type constructorService struct {
	repo repository.ConstructorRepository
}

func NewConstructorService(r repository.ConstructorRepository) ConstructorService {
	return &constructorService{repo: r}
}

func (s *constructorService) CreateConstructor(ctx context.Context, constructor model.Constructor) (model.Constructor, error) {
	constructor.ID = uuid.New()
	return s.repo.CreateConstructor(ctx, constructor)
}

func (s *constructorService) GetAllConstructors(ctx context.Context, page, limit int) ([]model.Constructor, error) {
	return s.repo.GetAllConstructors(ctx, page, limit)
}

func (s *constructorService) GetConstructorByName(ctx context.Context, name string, page, limit int) ([]model.Constructor, error) {
	return s.repo.GetConstructorByName(ctx, name, page, limit)
}

func (s *constructorService) GetConstructorByNationality(ctx context.Context, nationality string, page, limit int) ([]model.Constructor, error) {
	return s.repo.GetConstructorByNationality(ctx, nationality, page, limit)
}

func (s *constructorService) GetConstructorByRef(ctx context.Context, ref string) (model.Constructor, error) {
	return s.repo.GetConstructorByRef(ctx, ref)
}

func (s *constructorService) GetConstructorByID(ctx context.Context, id uuid.UUID) (model.Constructor, error) {
	return s.repo.GetConstructorByID(ctx, id)
}

func (s *constructorService) UpdateConstructor(ctx context.Context, id uuid.UUID, constructor model.Constructor) (model.Constructor, error) {
	return s.repo.UpdateConstructor(ctx, id, constructor)
}

func (s *constructorService) DeleteConstructor(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteConstructor(ctx, id)
}
