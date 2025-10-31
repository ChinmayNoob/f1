package repository

import (
	"context"
	"errors"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/utils"
	"github.com/ChinmayNoob/f1/pkg/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type DriverRepository interface {
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

type driverRepository struct{}

func NewDriverRepository() DriverRepository {
	return &driverRepository{}
}

func (r *driverRepository) CreateDriver(ctx context.Context, driver model.Driver) (model.Driver, error) {
	var constructorID uuid.UUID
	err := db.Conn.QueryRow(ctx, "SELECT id FROM constructors WHERE name = $1", driver.Constructor).Scan(&constructorID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Driver{}, errors.New("constructor not found")
		}
		return model.Driver{}, err
	}

	query := `
		INSERT INTO drivers (id, constructor_id, ref, code, number, first_name, last_name, date_of_birth, nationality, status, url)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, ref, code, number, first_name, last_name, date_of_birth, nationality, status, url
	`
	var createdDriver model.Driver
	err = db.Conn.QueryRow(ctx, query, driver.ID, constructorID, driver.Ref, driver.Code, driver.Number, driver.FirstName, driver.LastName, driver.DateOfBirth, driver.Nationality, driver.Status, driver.URL).Scan(
		&createdDriver.ID,
		&createdDriver.Ref,
		&createdDriver.Code,
		&createdDriver.Number,
		&createdDriver.FirstName,
		&createdDriver.LastName,
		&createdDriver.DateOfBirth,
		&createdDriver.Nationality,
		&createdDriver.Status,
		&createdDriver.URL,
	)
	if err != nil {
		return model.Driver{}, err
	}
	createdDriver.Constructor = driver.Constructor
	return createdDriver, nil
}

func (r *driverRepository) GetAllDrivers(ctx context.Context, page, limit int) ([]model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
	`
	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []model.Driver
	for rows.Next() {
		var driver model.Driver
		err := rows.Scan(
			&driver.ID,
			&driver.Constructor,
			&driver.Ref,
			&driver.Code,
			&driver.Number,
			&driver.FirstName,
			&driver.LastName,
			&driver.DateOfBirth,
			&driver.Nationality,
			&driver.Status,
			&driver.URL,
		)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (r *driverRepository) GetDriverByFirstName(ctx context.Context, firstName string, page, limit int) ([]model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.first_name = $1
	`
	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, firstName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []model.Driver
	for rows.Next() {
		var driver model.Driver
		err := rows.Scan(
			&driver.ID,
			&driver.Constructor,
			&driver.Ref,
			&driver.Code,
			&driver.Number,
			&driver.FirstName,
			&driver.LastName,
			&driver.DateOfBirth,
			&driver.Nationality,
			&driver.Status,
			&driver.URL,
		)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (r *driverRepository) GetDriverByLastName(ctx context.Context, lastName string, page, limit int) ([]model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.last_name = $1
	`
	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, lastName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []model.Driver
	for rows.Next() {
		var driver model.Driver
		err := rows.Scan(
			&driver.ID,
			&driver.Constructor,
			&driver.Ref,
			&driver.Code,
			&driver.Number,
			&driver.FirstName,
			&driver.LastName,
			&driver.DateOfBirth,
			&driver.Nationality,
			&driver.Status,
			&driver.URL,
		)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (r *driverRepository) GetDriverByTeam(ctx context.Context, constructorName string, page, limit int) ([]model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE c.name = $1
	`
	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, constructorName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []model.Driver
	for rows.Next() {
		var driver model.Driver
		err := rows.Scan(
			&driver.ID,
			&driver.Constructor,
			&driver.Ref,
			&driver.Code,
			&driver.Number,
			&driver.FirstName,
			&driver.LastName,
			&driver.DateOfBirth,
			&driver.Nationality,
			&driver.Status,
			&driver.URL,
		)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (r *driverRepository) GetDriverByID(ctx context.Context, id uuid.UUID) (model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.id = $1
	`
	var driver model.Driver
	err := db.Conn.QueryRow(ctx, query, id).Scan(
		&driver.ID,
		&driver.Constructor,
		&driver.Ref,
		&driver.Code,
		&driver.Number,
		&driver.FirstName,
		&driver.LastName,
		&driver.DateOfBirth,
		&driver.Nationality,
		&driver.Status,
		&driver.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Driver{}, nil
		}
		return model.Driver{}, err
	}
	return driver, nil
}

func (r *driverRepository) GetDriverByRef(ctx context.Context, ref string) (model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.ref = $1
	`
	var driver model.Driver
	err := db.Conn.QueryRow(ctx, query, ref).Scan(
		&driver.ID,
		&driver.Constructor,
		&driver.Ref,
		&driver.Code,
		&driver.Number,
		&driver.FirstName,
		&driver.LastName,
		&driver.DateOfBirth,
		&driver.Nationality,
		&driver.Status,
		&driver.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Driver{}, nil
		}
		return model.Driver{}, err
	}
	return driver, nil
}

func (r *driverRepository) GetDriverByCode(ctx context.Context, code string) (model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.code = $1
	`
	var driver model.Driver
	err := db.Conn.QueryRow(ctx, query, code).Scan(
		&driver.ID,
		&driver.Constructor,
		&driver.Ref,
		&driver.Code,
		&driver.Number,
		&driver.FirstName,
		&driver.LastName,
		&driver.DateOfBirth,
		&driver.Nationality,
		&driver.Status,
		&driver.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Driver{}, nil
		}
		return model.Driver{}, err
	}
	return driver, nil
}

func (r *driverRepository) GetDriverByNumber(ctx context.Context, number int) (model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.number = $1
	`
	var driver model.Driver
	err := db.Conn.QueryRow(ctx, query, number).Scan(
		&driver.ID,
		&driver.Constructor,
		&driver.Ref,
		&driver.Code,
		&driver.Number,
		&driver.FirstName,
		&driver.LastName,
		&driver.DateOfBirth,
		&driver.Nationality,
		&driver.Status,
		&driver.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Driver{}, nil
		}
		return model.Driver{}, err
	}
	return driver, nil
}

func (r *driverRepository) GetDriverByNationality(ctx context.Context, nationality string, page, limit int) ([]model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.nationality = $1
	`
	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, nationality)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []model.Driver
	for rows.Next() {
		var driver model.Driver
		err := rows.Scan(
			&driver.ID,
			&driver.Constructor,
			&driver.Ref,
			&driver.Code,
			&driver.Number,
			&driver.FirstName,
			&driver.LastName,
			&driver.DateOfBirth,
			&driver.Nationality,
			&driver.Status,
			&driver.URL,
		)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (r *driverRepository) GetDriverByStatus(ctx context.Context, status string, page, limit int) ([]model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.status = $1
	`
	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drivers []model.Driver
	for rows.Next() {
		var driver model.Driver
		err := rows.Scan(
			&driver.ID,
			&driver.Constructor,
			&driver.Ref,
			&driver.Code,
			&driver.Number,
			&driver.FirstName,
			&driver.LastName,
			&driver.DateOfBirth,
			&driver.Nationality,
			&driver.Status,
			&driver.URL,
		)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func (r *driverRepository) GetDriverByURL(ctx context.Context, url string) (model.Driver, error) {
	query := `
		SELECT d.id, c.name as constructor, d.ref, d.code, d.number, d.first_name, d.last_name, d.date_of_birth, d.nationality, d.status, d.url
		FROM drivers d
		INNER JOIN constructors c ON d.constructor_id = c.id
		WHERE d.url = $1
	`
	var driver model.Driver
	err := db.Conn.QueryRow(ctx, query, url).Scan(
		&driver.ID,
		&driver.Constructor,
		&driver.Ref,
		&driver.Code,
		&driver.Number,
		&driver.FirstName,
		&driver.LastName,
		&driver.DateOfBirth,
		&driver.Nationality,
		&driver.Status,
		&driver.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Driver{}, nil
		}
		return model.Driver{}, err
	}
	return driver, nil
}

func (r *driverRepository) UpdateDriver(ctx context.Context, id uuid.UUID, driver model.Driver) (model.Driver, error) {
	var constructorID uuid.UUID
	err := db.Conn.QueryRow(ctx, "SELECT id FROM constructors WHERE name = $1", driver.Constructor).Scan(&constructorID)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Driver{}, errors.New("constructor not found")
		}
		return model.Driver{}, err
	}

	query := `
		UPDATE drivers
		SET constructor_id = $1, ref = $2, code = $3, number = $4, first_name = $5, last_name = $6, date_of_birth = $7, nationality = $8, status = $9, url = $10
		WHERE id = $11
		RETURNING id, ref, code, number, first_name, last_name, date_of_birth, nationality, status, url
	`
	var updatedDriver model.Driver
	err = db.Conn.QueryRow(ctx, query, constructorID, driver.Ref, driver.Code, driver.Number, driver.FirstName, driver.LastName, driver.DateOfBirth, driver.Nationality, driver.Status, driver.URL, id).Scan(
		&updatedDriver.ID,
		&updatedDriver.Ref,
		&updatedDriver.Code,
		&updatedDriver.Number,
		&updatedDriver.FirstName,
		&updatedDriver.LastName,
		&updatedDriver.DateOfBirth,
		&updatedDriver.Nationality,
		&updatedDriver.Status,
		&updatedDriver.URL,
	)
	if err != nil {
		return model.Driver{}, err
	}
	updatedDriver.Constructor = driver.Constructor
	return updatedDriver, nil
}

func (r *driverRepository) DeleteDriver(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM drivers WHERE id = $1`
	_, err := db.Conn.Exec(ctx, query, id)
	return err
}
