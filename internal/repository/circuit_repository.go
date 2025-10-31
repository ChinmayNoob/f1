package repository

import (
	"context"

	"github.com/ChinmayNoob/f1/internal/model"
	"github.com/ChinmayNoob/f1/internal/utils"
	"github.com/ChinmayNoob/f1/pkg/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CircuitRepository interface {
	CreateCircuit(ctx context.Context, circuit model.Circuit) (model.Circuit, error)
	GetAllCircuits(ctx context.Context, page, limit int) ([]model.Circuit, error)
	GetCircuitByID(ctx context.Context, id uuid.UUID) (model.Circuit, error)
	GetCircuitByRef(ctx context.Context, ref string) (model.Circuit, error)
	GetCircuitByName(ctx context.Context, name string, page, limit int) ([]model.Circuit, error)
	GetCircuitByLocation(ctx context.Context, location string, page, limit int) ([]model.Circuit, error)
	GetCircuitByCountry(ctx context.Context, country string, page, limit int) ([]model.Circuit, error)
	GetCircuitByCurrent(ctx context.Context, current bool) ([]model.Circuit, error)
	GetCircuitByURL(ctx context.Context, url string) (model.Circuit, error)
	UpdateCircuit(ctx context.Context, id uuid.UUID, circuit model.Circuit) (model.Circuit, error)
	DeleteCircuit(ctx context.Context, id uuid.UUID) error
}

type circuitRepository struct{}

func NewCircuitRepository() CircuitRepository {
	return &circuitRepository{}
}

func (r *circuitRepository) CreateCircuit(ctx context.Context, circuit model.Circuit) (model.Circuit, error) {
	query := `
		INSERT INTO circuits (id, ref, name, location, country, "current", url)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING *
	`

	var createdCircuit model.Circuit
	err := db.Conn.QueryRow(ctx, query, circuit.ID, circuit.Ref, circuit.Name, circuit.Location, circuit.Country, circuit.Current, circuit.URL).Scan(
		&createdCircuit.ID,
		&createdCircuit.Ref,
		&createdCircuit.Name,
		&createdCircuit.Location,
		&createdCircuit.Country,
		&createdCircuit.Current,
		&createdCircuit.URL,
	)
	if err != nil {
		return model.Circuit{}, err
	}
	return createdCircuit, nil
}

func (r *circuitRepository) GetAllCircuits(ctx context.Context, page, limit int) ([]model.Circuit, error) {
	query := `SELECT * FROM circuits`

	paginationQuery, err := utils.Paginate(query, page, limit)

	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circuits []model.Circuit
	for rows.Next() {
		var circuit model.Circuit
		err := rows.Scan(
			&circuit.ID,
			&circuit.Ref,
			&circuit.Name,
			&circuit.Location,
			&circuit.Country,
			&circuit.Current,
			&circuit.URL,
		)
		if err != nil {
			return nil, err
		}
		circuits = append(circuits, circuit)
	}
	return circuits, nil
}

func (r *circuitRepository) GetCircuitByID(ctx context.Context, id uuid.UUID) (model.Circuit, error) {
	query := `SELECT * FROM circuits WHERE id = $1`
	var circuit model.Circuit
	err := db.Conn.QueryRow(ctx, query, id).Scan(
		&circuit.ID,
		&circuit.Ref,
		&circuit.Name,
		&circuit.Location,
		&circuit.Country,
		&circuit.Current,
		&circuit.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Circuit{}, nil
		}
		return model.Circuit{}, err
	}
	return circuit, nil
}

func (r *circuitRepository) GetCircuitByRef(ctx context.Context, ref string) (model.Circuit, error) {
	query := `SELECT * FROM circuits WHERE ref = $1`
	var circuit model.Circuit
	err := db.Conn.QueryRow(ctx, query, ref).Scan(
		&circuit.ID,
		&circuit.Ref,
		&circuit.Name,
		&circuit.Location,
		&circuit.Country,
		&circuit.Current,
		&circuit.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Circuit{}, nil
		}
		return model.Circuit{}, err
	}
	return circuit, nil
}

func (r *circuitRepository) GetCircuitByName(ctx context.Context, name string, page, limit int) ([]model.Circuit, error) {
	query := `SELECT * FROM circuits WHERE name = $1`

	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circuits []model.Circuit
	for rows.Next() {
		var circuit model.Circuit
		err := rows.Scan(
			&circuit.ID,
			&circuit.Ref,
			&circuit.Name,
			&circuit.Location,
			&circuit.Country,
			&circuit.Current,
			&circuit.URL,
		)
		if err != nil {
			return nil, err
		}
		circuits = append(circuits, circuit)
	}
	return circuits, nil
}

func (r *circuitRepository) GetCircuitByLocation(ctx context.Context, location string, page, limit int) ([]model.Circuit, error) {
	query := `SELECT * FROM circuits WHERE location = $1`

	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, location)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circuits []model.Circuit
	for rows.Next() {
		var circuit model.Circuit
		err := rows.Scan(
			&circuit.ID,
			&circuit.Ref,
			&circuit.Name,
			&circuit.Location,
			&circuit.Country,
			&circuit.Current,
			&circuit.URL,
		)
		if err != nil {
			return nil, err
		}
		circuits = append(circuits, circuit)
	}
	return circuits, nil
}

func (r *circuitRepository) GetCircuitByCountry(ctx context.Context, country string, page, limit int) ([]model.Circuit, error) {
	query := `SELECT * FROM circuits WHERE country = $1`
	paginationQuery, err := utils.Paginate(query, page, limit)
	if err != nil {
		return nil, err
	}

	rows, err := db.Conn.Query(ctx, paginationQuery, country)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var circuits []model.Circuit
	for rows.Next() {
		var circuit model.Circuit
		err := rows.Scan(
			&circuit.ID,
			&circuit.Ref,
			&circuit.Name,
			&circuit.Location,
			&circuit.Country,
			&circuit.Current,
			&circuit.URL,
		)
		if err != nil {
			return nil, err
		}
		circuits = append(circuits, circuit)
	}
	return circuits, nil
}

func (r *circuitRepository) GetCircuitByCurrent(ctx context.Context, current bool) ([]model.Circuit, error) {
	query := `SELECT * FROM circuits WHERE "current" = $1`

	rows, err := db.Conn.Query(ctx, query, current)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circuits []model.Circuit
	for rows.Next() {
		var circuit model.Circuit
		err := rows.Scan(
			&circuit.ID,
			&circuit.Ref,
			&circuit.Name,
			&circuit.Location,
			&circuit.Country,
			&circuit.Current,
			&circuit.URL,
		)
		if err != nil {
			return nil, err
		}
		circuits = append(circuits, circuit)
	}
	return circuits, nil
}

func (r *circuitRepository) GetCircuitByURL(ctx context.Context, url string) (model.Circuit, error) {
	query := `SELECT * FROM circuits WHERE url = $1`
	var circuit model.Circuit
	err := db.Conn.QueryRow(ctx, query, url).Scan(
		&circuit.ID,
		&circuit.Ref,
		&circuit.Name,
		&circuit.Location,
		&circuit.Country,
		&circuit.Current,
		&circuit.URL,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return model.Circuit{}, nil
		}
		return model.Circuit{}, err
	}
	return circuit, nil
}

func (r *circuitRepository) UpdateCircuit(ctx context.Context, id uuid.UUID, circuit model.Circuit) (model.Circuit, error) {
	query := `
		UPDATE circuits
		SET ref = $1, name = $2, location = $3, country = $4, "current" = $5, url = $6
		WHERE id = $7
		RETURNING *
	`
	var updatedCircuit model.Circuit
	err := db.Conn.QueryRow(ctx, query, circuit.Ref, circuit.Name, circuit.Location, circuit.Country, circuit.Current, circuit.URL, id).Scan(
		&updatedCircuit.ID,
		&updatedCircuit.Ref,
		&updatedCircuit.Name,
		&updatedCircuit.Location,
		&updatedCircuit.Country,
		&updatedCircuit.Current,
		&updatedCircuit.URL,
	)
	if err != nil {
		return model.Circuit{}, err
	}
	return updatedCircuit, nil
}

func (r *circuitRepository) DeleteCircuit(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM circuits WHERE id = $1`
	_, err := db.Conn.Exec(ctx, query, id)
	return err
}
