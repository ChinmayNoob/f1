package model

import (
	"time"

	"github.com/google/uuid"
)

type Season struct {
	ID   uuid.UUID `json:"id"`
	Year int       `json:"year"`
	URL  string    `json:"url"`
}

type Race struct {
	ID        uuid.UUID `json:"id"`
	SeasonID  uuid.UUID `json:"season_id"`
	CircuitID uuid.UUID `json:"circuit_id"`
	Round     int       `json:"round"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	URL       string    `json:"url"`
}

type Result struct {
	ID            uuid.UUID `json:"id"`
	RaceID        uuid.UUID `json:"race_id"`
	DriverID      uuid.UUID `json:"driver_id"`
	ConstructorID uuid.UUID `json:"constructor_id"`
	Number        int       `json:"number"`
	Grid          int       `json:"grid"`
	Position      *int      `json:"position"`
	PositionText  string    `json:"position_text"`
	Points        float64   `json:"points"`
	Laps          int       `json:"laps"`
	Time          string    `json:"time"`
	Status        string    `json:"status"`
}

type DriverStanding struct {
	ID       uuid.UUID `json:"id"`
	SeasonID uuid.UUID `json:"season_id"`
	DriverID uuid.UUID `json:"driver_id"`
	Position int       `json:"position"`
	Points   float64   `json:"points"`
	Wins     int       `json:"wins"`
}

type ConstructorStanding struct {
	ID            uuid.UUID `json:"id"`
	SeasonID      uuid.UUID `json:"season_id"`
	ConstructorID uuid.UUID `json:"constructor_id"`
	Position      int       `json:"position"`
	Points        float64   `json:"points"`
	Wins          int       `json:"wins"`
}
