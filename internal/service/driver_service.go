package service

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/repository"
	"github.com/google/uuid"
)

type DriverService interface {
	CreateDriver(ctx context.Context, driver model.Driver) (model.Driver, error)
	GetAllDrivers(ctx context.Context, page, limit int) ([]model.Driver, error)
	GetDriverByFirstName(ctx context.Context, firstName string, page, limit int) ([]model.Driver, error)
	GetDriverByLastName(ctx context.Context, lastName string, page, limit int) ([]model.Driver, error)
	GetDriverByTeam(ctx context.Context, constructorName string, page, limit int) ([]model.Driver, error)
	GetDriverByID(ctx context.Context, id uuid.UUID) (model.Driver, error)
	GetDriverByRef(ctx context.Context, ref string) (model.Driver, error)
	GetDriverByCode(ctx context.Context, code string) (model.Driver, error)
	GetDriverByNumber(ctx context.Context, number int) (model.Driver, error)
	GetDriverByNationality(ctx context.Context, nationality string, page, limit int) ([]model.Driver, error)
	GetDriverByStatus(ctx context.Context, status string, page, limit int) ([]model.Driver, error)
	GetDriverByURL(ctx context.Context, url string) (model.Driver, error)
	UpdateDriver(ctx context.Context, id uuid.UUID, driver model.Driver) (model.Driver, error)
	DeleteDriver(ctx context.Context, id uuid.UUID) error
}

type driverService struct {
	repo repository.DriverRepository
}

func NewDriverService(r repository.DriverRepository) DriverService {
	return &driverService{repo: r}
}

func (s *driverService) CreateDriver(ctx context.Context, driver model.Driver) (model.Driver, error) {
	driver.ID = uuid.New()
	return s.repo.CreateDriver(ctx, driver)
}

func (s *driverService) GetAllDrivers(ctx context.Context, page, limit int) ([]model.Driver, error) {
	return s.repo.GetAllDrivers(ctx, page, limit)
}

func (s *driverService) GetDriverByFirstName(ctx context.Context, firstName string, page, limit int) ([]model.Driver, error) {
	return s.repo.GetDriverByFirstName(ctx, firstName, page, limit)
}

func (s *driverService) GetDriverByLastName(ctx context.Context, lastName string, page, limit int) ([]model.Driver, error) {
	return s.repo.GetDriverByLastName(ctx, lastName, page, limit)
}

func (s *driverService) GetDriverByTeam(ctx context.Context, constructorName string, page, limit int) ([]model.Driver, error) {
	return s.repo.GetDriverByTeam(ctx, constructorName, page, limit)
}

func (s *driverService) GetDriverByID(ctx context.Context, id uuid.UUID) (model.Driver, error) {
	return s.repo.GetDriverByID(ctx, id)
}

func (s *driverService) GetDriverByRef(ctx context.Context, ref string) (model.Driver, error) {
	return s.repo.GetDriverByRef(ctx, ref)
}

func (s *driverService) GetDriverByCode(ctx context.Context, code string) (model.Driver, error) {
	return s.repo.GetDriverByCode(ctx, code)
}

func (s *driverService) GetDriverByNumber(ctx context.Context, number int) (model.Driver, error) {
	return s.repo.GetDriverByNumber(ctx, number)
}

func (s *driverService) GetDriverByNationality(ctx context.Context, nationality string, page, limit int) ([]model.Driver, error) {
	return s.repo.GetDriverByNationality(ctx, nationality, page, limit)
}

func (s *driverService) GetDriverByStatus(ctx context.Context, status string, page, limit int) ([]model.Driver, error) {
	return s.repo.GetDriverByStatus(ctx, status, page, limit)
}

func (s *driverService) GetDriverByURL(ctx context.Context, url string) (model.Driver, error) {
	return s.repo.GetDriverByURL(ctx, url)
}

func (s *driverService) UpdateDriver(ctx context.Context, id uuid.UUID, driver model.Driver) (model.Driver, error) {
	return s.repo.UpdateDriver(ctx, id, driver)
}

func (s *driverService) DeleteDriver(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDriver(ctx, id)
}
