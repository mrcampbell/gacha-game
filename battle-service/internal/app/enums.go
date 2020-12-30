package app

type Element string

const (
	ElementFire  Element = "fire"
	ElementWater Element = "water"
	ElementGrass Element = "grass"
)

type FighterType string

const (
	FighterTypeTank     FighterType = "tank"
	FighterTypeAttacker FighterType = "attacker"
	FighterTypeSupport  FighterType = "support"
)

type MoveTargetType string

const (
	MoveTargetTypeAllEnemies     MoveTargetType = "all_enemies"
	MoveTargetTypeAllAllies      MoveTargetType = "all_allies"
	MoveTargetTypeAllSingleEnemy MoveTargetType = "single_enemy"
	MoveTargetTypeAllSingleAlly  MoveTargetType = "single_ally"
	MoveTargetTypeAllSelf        MoveTargetType = "self"
)

type MoveType string

const (
	MoveTypeHeal            MoveType = "heal"
	MoveTypeDamage          MoveType = "damage"
	MoveTypeBuff            MoveType = "buff"
	MoveTypeDebuff          MoveType = "debuff"
	MoveTypeDamageAndDebuff MoveType = "damage+debuff"
	MoveTypeHealAndBuff     MoveType = "heal+buff"
)
