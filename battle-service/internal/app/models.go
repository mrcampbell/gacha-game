package app

type FighterUnit struct {
	Name    string      `json:"name"`
	Element Element     `json:"element"`
	Type    FighterType `json:"type"`
}

type Fighter struct {
	Unit  FighterUnit `json:"unit"`
	Level int32       `json:"level"`
}
