package app

type FighterUnit struct {
	ID      string      `json:"id"`
	Name    string      `json:"name"`
	Element Element     `json:"element"`
	Type    FighterType `json:"type"`
}

type Fighter struct {
	ID    string      `json:"id"`
	Unit  FighterUnit `json:"unit"`
	Level int32       `json:"level"`
	Moves []Move      `json:"moves"`
}

type Move struct {
	ID         string         `json:"id"`
	Name       string         `json:"name"`
	TargetType MoveTargetType `json:"target"`
	Type       MoveType       `json:"type"`
	Power      int            `json:"power"` // percent based?  -10% enemy health, etc
}

type FighterUnitMove struct {
	LevelLearnAt int
	FighterID    string
	MoveID       string
}
