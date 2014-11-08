package models

type ShipEnum string

const (
	CARRIER    ShipEnum = "Carrier"
	BATTLESHIP ShipEnum = "Battleship"
	DESTROYER  ShipEnum = "Destroyer"
	SUBMARINE  ShipEnum = "Submarine"
	PATROL     ShipEnum = "Patrol Boat"
)

type Ship struct {
	ShipType       *string    `json:"ship_type,omitempty"`
	Size           *int64     `json:"size,omitempty"`
	Hits           *[]bool    `json:"hits,omitempty"`
	Origin         *Point     `json:"origin,omitempty"`
	Orientation    *Direction `json:"orientation,omitempty"`
	Sight          *int64     `json:"-"`
	Range          *int64     `json:"-"`
	MoveCap        *int64     `json:"-"`
	IsUsingSpecial *bool      `json:"is_using_special,omitempty"`
}

type Point struct {
	X *uint `json:"x,omitempty"`
	Y *uint `json:"y,omitempty"`
}

/*
 * Enum for ship direction
 */
type Enum interface {
	name() string
	ordinal() int
	valueOf() *[]string
}

type Direction uint

func (d Direction) name() string {
	return directions[d]
}
func (d Direction) ordinal() int {
	return int(d)
}
func (d Direction) String() string {
	return directions[d]
}
func (d Direction) values() *[]string {
	return &directions
}

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

var directions = []string{
	"North",
	"Eash",
	"South",
	"West",
}
