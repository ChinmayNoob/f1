package model

import (
	"github.com/google/uuid"
)

type Podium struct {
	FirstPositionDriverID  uuid.UUID `json:"first_position_driver_id"`
	SecondPositionDriverID uuid.UUID `json:"second_position_driver_id"`
	ThirdPositionDriverID  uuid.UUID `json:"third_position_driver_id"`
}

type Circuit struct {
	ID       uuid.UUID `json:"id"`
	Ref      string    `json:"ref"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Country  string    `json:"country"`
	Current  bool      `json:"current"`
	URL      string    `json:"url"`
	Podiums  Podium    `json:"podiums"`
}
