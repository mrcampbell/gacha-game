package app

type Element string

const (
	FIRE  Element = "fire"
	WATER Element = "water"
	GRASS Element = "grass"
)

type FighterType string

const (
	TANK     FighterType = "tank"
	ATTACKER FighterType = "attacker"
	SUPPORT  FighterType = "support"
)
