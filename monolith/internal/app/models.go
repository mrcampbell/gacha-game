package app

type UnitType int32

const (
	UnitTypeUnknown  UnitType = 0
	UnitTypeTank     UnitType = 1
	UnitTypeHealer   UnitType = 2
	UnitTypeSupport  UnitType = 3
	UnitTypeAttacker UnitType = 4
)

type ElementType int32

const (
	ElementUnknown ElementType = 0
	ElementFire    ElementType = 1
	ElementWater   ElementType = 2
	ElementGrass   ElementType = 3
	ElementNormal  ElementType = 4
)

type Unit struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Element  ElementType `json:"element"`
	UnitType UnitType    `json:"unit_type"` // because type is a reserved word
}
