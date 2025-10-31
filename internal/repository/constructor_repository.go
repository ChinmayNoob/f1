package repository

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/utils"
	"github.com/ChinmayNoob/f1/pkg/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ConstructorRepository interface {
	CreateConstructor(ctx context.Context, constructor model.Constructor) (model.Constructor, error)
	GetAllConstructors(ctx context.Context, page, limit int) ([]model.Constructor, error)
	GetConstructorByName(ctx context.Context, name string, page, limit int) ([]model.Constructor, error)
	GetConstructorByID(ctx context.Context, id uuid.UUID) (model.Constructor, error)
	GetConstructorByNationality(ctx context.Context, nationality string, page, limit int) ([]model.Constructor, error)
	GetConstructorByRef(ctx context.Context, ref string) (model.Constructor, error)
	UpdateConstructor(ctx context.Context, id uuid.UUID, constructor model.Constructor) (model.Constructor, error)
	DeleteConstructor(ctx context.Context, id uuid.UUID) error
}

type constructorRepository struct{}

func NewConstructorRepository() ConstructorRepository {
	return &constructorRepository{}
}

func (r *constructorRepository) CreateConstructor(ctx context.Context, constructor model.Constructor) (model.Constructor, error) {
	query := `
		INSERT INTO constructors (id, ref, name, nationality, url)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`
	var createdConstructor model.Constructor
	err := db.Conn.QueryRow(ctx, query, constructor.ID, constructor.Ref, constructor.Name, constructor.Nationality, constructor.URL).Scan(
		&createdConstructor.ID,
		&createdConstructor.Ref,
		&createdConstructor.Name,
		&createdConstructor.Nationality,
		&createdConstructor.URL,
	)
	if err != nil {
		return model.Constructor{}, err
	}
	return createdConstructor, nil
}

func (r *constructorRepository) GetAllConstructors(ctx context.Context, page, limit int) ([]model.Constructor, error) {
	query := `SELECT * FROM constructors`

	paginationQuery, err := utils.Paginate(query, page, limit)

	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var constructors []model.Constructor
	for rows.Next() {
		var constructor model.Constructor
		err := rows.Scan(
			&constructor.ID,
			&constructor.Ref,
			&constructor.Name,
			&constructor.Nationality,
			&constructor.URL,
		)
		if err != nil {
			return nil, err
		}
		constructors = append(constructors, constructor)
	}
	return constructors, nil
}

func (r *constructorRepository) GetConstructorByName(ctx context.Context, name string, page, limit int) ([]model.Constructor, error) {
	query := `SELECT * FROM constructors WHERE name = $1`
	paginationQuery, err := utils.Paginate(query, page, limit)

	if err != nil {
		return nil, err
	}
	rows, err := db.Conn.Query(ctx, paginationQuery, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var constructors []model.Constructor
	for rows.Next() {
		var constructor model.Constructor
		err := rows.Scan(
			&constructor.ID,
			&constructor.Ref,
			&constructor.Name,
			&constructor.Nationality,
			&constructor.URL,
		)
		if err != nil {
			return nil, err
		}
		constructors = append(constructors, constructor)
	}
	return constructors, nil
}

func (r *constructorRepository) GetConstructorByID(ctx context.Context, id uuid.UUID) (model.Constructor, error) {
	query := `SELECT * FROM constructors WHERE id = $1`
	var constructor model.Constructor
	err := db.Conn.QueryRow(ctx, query, id).Scan(
		&constructor.ID,
		&constructor.Ref,
		&constructor.Name,
		&constructor.Nationality,
		&constructor.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Constructor{}, nil
		}
		return model.Constructor{}, err
	}
	return constructor, nil
}

func (r *constructorRepository) GetConstructorByRef(ctx context.Context, ref string) (model.Constructor, error) {
	query := `SELECT * FROM constructors WHERE ref = $1`
	var constructor model.Constructor
	err := db.Conn.QueryRow(ctx, query, ref).Scan(
		&constructor.ID,
		&constructor.Ref,
		&constructor.Name,
		&constructor.Nationality,
		&constructor.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Constructor{}, nil
		}
		return model.Constructor{}, err
	}
	return constructor, nil
}

func (r *constructorRepository) GetConstructorByNationality(ctx context.Context, nationality string, page, limit int) ([]model.Constructor, error) {
	query := `SELECT * FROM constructors WHERE nationality = $1`

	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, nationality)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var constructors []model.Constructor
	for rows.Next() {
		var constructor model.Constructor
		err := rows.Scan(
			&constructor.ID,
			&constructor.Ref,
			&constructor.Name,
			&constructor.Nationality,
			&constructor.URL,
		)
		if err != nil {
			return nil, err
		}
		constructors = append(constructors, constructor)
	}
	return constructors, nil
}

func (r *constructorRepository) UpdateConstructor(ctx context.Context, id uuid.UUID, constructor model.Constructor) (model.Constructor, error) {
	query := `
		UPDATE constructors
		SET ref = $1, name = $2, nationality = $3, url = $4
		WHERE id = $5
		RETURNING *
	`
	var updatedConstructor model.Constructor
	err := db.Conn.QueryRow(ctx, query, constructor.Ref, constructor.Name, constructor.Nationality, constructor.URL, id).Scan(
		&updatedConstructor.ID,
		&updatedConstructor.Ref,
		&updatedConstructor.Name,
		&updatedConstructor.Nationality,
		&updatedConstructor.URL,
	)
	if err != nil {
		return model.Constructor{}, err
	}
	return updatedConstructor, nil
}

func (r *constructorRepository) DeleteConstructor(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM constructors WHERE id = $1`
	_, err := db.Conn.Exec(ctx, query, id)
	return err
}
