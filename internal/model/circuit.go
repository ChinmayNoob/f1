package model

import (
	"github.com/google/uuid"
)

// type Podium struct {
// 	FirstPositionDriverName  string `json:"first_position_driver_name"`
// 	SecondPositionDriverName string `json:"second_position_driver_name"`
// 	ThirdPositionDriverName  string `json:"third_position_driver_name"`
// }

type Circuit struct {
	ID       uuid.UUID `json:"id"`
	Ref      string    `json:"ref"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Country  string    `json:"country"`
	Current  bool      `json:"current"`
	URL      string    `json:"url"`
}
