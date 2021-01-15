package app

type UnitType int32

const (
	UnitTypeUnknown  UnitType = 0
	UnitTypeTank     UnitType = 1
	UnitTypeHealer   UnitType = 2
	UnitTypeSupport  UnitType = 3
	UnitTypeAttacker UnitType = 4
)

type Element int32

const (
	ElementUnknown Element = 0
	ElementFire    Element = 1
	ElementWater   Element = 2
	ElementGrass   Element = 3
	ElementNormal  Element = 4
)

type StatType string

const (
	StatHealth  StatType = "HEALTH"
	StatAttack  StatType = "ATTACK"
	StatDefense StatType = "DEFENSE"
	StatSpeed   StatType = "SPEED"
)

type MoveTargetType string

const (
	MoveTargetTypeSelf        MoveTargetType = "SELF"
	MoveTargetTypeAllAllies   MoveTargetType = "ALL_ALLIES"
	MoveTargetTypeAllEnemies  MoveTargetType = "ALL_ENEMIES"
	MoveTargetTypeSingleAlly  MoveTargetType = "SINGLE_ALLY"
	MoveTargetTypeSingleEnemy MoveTargetType = "SINGLE_ENEMY"
)

type MoveMetaEffectType string

const (
	MoveMetaEffectTypeHeal       MoveMetaEffectType = "HEAL"
	MoveMetaEffectTypeDamage     MoveMetaEffectType = "DAMAGE"
	MoveMetaEffectTypeStatChange MoveMetaEffectType = "STAT_CHANGE"
	MoveMetaEffectTypeLockOnMe   MoveMetaEffectType = "LOCK_ON_ME"
)

type MoveMeta interface {
	Type() MoveMetaEffectType
}

type Move struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Element      Element    `json:"element"`
	Metas        []MoveMeta `json:"metas"` // this is tricky, more a demo than anything
	RestDuration int32      `json:"rest_duration"`
}

type MoveMetaHeal struct {
	moveMetaEffectType MoveMetaEffectType
	Target             MoveTargetType `json:"target"`
	PercentOfMaxHealth int32          `json:"percent_of_max_health"`
}

func (mhh MoveMetaHeal) Type() MoveMetaEffectType {
	return mhh.moveMetaEffectType
}
func NewMoveMetaHeal(target MoveTargetType, percentOfMaxHealth int32) MoveMetaHeal {
	return MoveMetaHeal{
		moveMetaEffectType: MoveMetaEffectTypeHeal,
		Target:             target,
		PercentOfMaxHealth: percentOfMaxHealth,
	}
}

type MoveMetaDamage struct {
	moveMetaEffectType MoveMetaEffectType
	Target             MoveTargetType `json:"target"`
	Power              int32          `json:"power"`
}

func (m MoveMetaDamage) Type() MoveMetaEffectType {
	return m.moveMetaEffectType
}
func NewMoveMetaDamage(target MoveTargetType, power int32) MoveMetaDamage {
	return MoveMetaDamage{
		moveMetaEffectType: MoveMetaEffectTypeDamage,
		Target:             target,
		Power:              power,
	}
}

type MoveMetaStatChange struct {
	moveMetaEffectType MoveMetaEffectType
	Target             MoveTargetType `json:"target"`
	StatAffected       StatType       `json:"stat_affected"`
	Magnitude          int32          `json:"magnitude"` // positive number for buff, negative number for debuff
	Duration           int32          `json:"duration"`
}

func (m MoveMetaStatChange) Type() MoveMetaEffectType {
	return m.moveMetaEffectType
}
func NewMoveMetaStatChange(target MoveTargetType, statAffected StatType, magnitude int32, duration int32) MoveMetaStatChange {
	return MoveMetaStatChange{
		moveMetaEffectType: MoveMetaEffectTypeStatChange,
		Target:             target,
		Magnitude:          magnitude,
		StatAffected:       statAffected,
		Duration:           duration,
	}
}

type MoveMetaLockOnMe struct {
	moveMetaEffectType MoveMetaEffectType
	Duration           int32 `json:"duration"`
}

func (m MoveMetaLockOnMe) Type() MoveMetaEffectType {
	return m.moveMetaEffectType
}
func NewMoveMetaLockOnMe(duration int32) MoveMetaLockOnMe {
	return MoveMetaLockOnMe{
		moveMetaEffectType: MoveMetaEffectTypeLockOnMe,
		Duration:           duration,
	}
}

type UnitMove struct {
	UnitID         string `json:"unit_id"`
	MoveID         string `json:"move_id"`
	LevelLearnedAt int32  `json:"level_learned_at"`
}

type StatGroup struct {
	Health  int32 `json:"health"`
	Attack  int32 `json:"attack"`
	Defense int32 `json:"defense"`
	Speed   int32 `json:"speed"`
}

type Unit struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Element  Element   `json:"element"`
	UnitType UnitType  `json:"unit_type"` // because `type` is a reserved word
	Stats    StatGroup `json:"stats"`
}

type Fighter struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UnitID       string    `json:"unit_id"` // for hydration
	BaseUnit     Unit      `json:"base_unit"`
	Level        int32     `json:"string"`
	CurrentStats StatGroup `json:"stats"` // calculated at level up
	Moves        []Move    `json:"moves"`
}

func (f Fighter) CalculateCurrentStats() StatGroup {
	calcStat := func(level, baseStat int32) int32 {
		return level * baseStat / 4 // this is arbitrary.  We'll work this out later
	}

	sg := StatGroup{
		Health:  calcStat(f.Level, f.BaseUnit.Stats.Health) + 20,
		Attack:  calcStat(f.Level, f.BaseUnit.Stats.Attack),
		Defense: calcStat(f.Level, f.BaseUnit.Stats.Defense),
		Speed:   calcStat(f.Level, f.BaseUnit.Stats.Speed),
	}

	return sg
}
