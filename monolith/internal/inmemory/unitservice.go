package inmemory

import (
	"context"
	"fmt"
	"strconv"

	"github.com/mrcampbell/gacha-game/monolith/internal/app"
)

// ensure the interface is implemented correctly
var _ app.UnitService = &UnitService{}

type UnitService struct {
	units map[string]app.Unit
}

func NewUnitService() (UnitService, error) {
	us := UnitService{}
	err := us.initialize()
	return us, err
}

// Stats By Element

//				Hp	Atk	Def	Speed	Average
// Fire		1		3		1		3			2
// Water	2		2		2		2			2
// Grass	3		1		3		1			2
// Norm		3		1		1		3			2

// Stats By Class

//					Hp	Atk	Def	Speed	Average
// Tank			5		1		5		1			3
// Healer		3		2		3		4			3
// Support	2		1		4		5			3
// Attacker	3		5		2		2			3

var FireElementStatMod = app.StatGroup{
	Health:  1,
	Attack:  3,
	Defense: 1,
	Speed:   3,
}

var WaterElementStatMod = app.StatGroup{
	Health:  2,
	Attack:  2,
	Defense: 2,
	Speed:   2,
}

var GrassElementStatMod = app.StatGroup{
	Health:  3,
	Attack:  1,
	Defense: 3,
	Speed:   1,
}

var NormalElementStatMod = app.StatGroup{
	Health:  3,
	Attack:  1,
	Defense: 1,
	Speed:   3,
}

var TankClassStatMod = app.StatGroup{
	Health:  5,
	Attack:  1,
	Defense: 5,
	Speed:   1,
}

var HealerClassStatMod = app.StatGroup{
	Health:  3,
	Attack:  2,
	Defense: 3,
	Speed:   4,
}

var SupportClassStatMod = app.StatGroup{
	Health:  2,
	Attack:  1,
	Defense: 4,
	Speed:   5,
}

var AttackerClassStatMod = app.StatGroup{
	Health:  3,
	Attack:  5,
	Defense: 2,
	Speed:   2,
}

var elementStatMap = map[app.Element]app.StatGroup{
	app.ElementFire:   FireElementStatMod,
	app.ElementGrass:  GrassElementStatMod,
	app.ElementWater:  WaterElementStatMod,
	app.ElementNormal: NormalElementStatMod,
}

var classStatMap = map[app.UnitType]app.StatGroup{
	app.UnitTypeHealer:   HealerClassStatMod,
	app.UnitTypeAttacker: AttackerClassStatMod,
	app.UnitTypeSupport:  SupportClassStatMod,
	app.UnitTypeTank:     TankClassStatMod,
}

func (us UnitService) UnitByID(ctx context.Context, id string) (app.Unit, error) {
	unit, ok := us.units[id]
	if !ok {
		return app.Unit{}, fmt.Errorf("no unit with id [%s]", id)
	}

	// we purposefully didn't set this in the map above so we could shift around the order without having to worry
	// about inconsistencies between the unit's hardcoded ID and the ID in reference to where it is in the map
	unit.ID = id

	return unit, nil
}

func (us UnitService) ListAllUnits(ctx context.Context) ([]app.Unit, error) {
	units := []app.Unit{}
	for key, unit := range us.units {
		unit.ID = key
		units = append(units, unit)
	}
	return units, nil
}

func (us UnitService) CreateUnit(ctx context.Context, unit app.Unit) (app.Unit, error) {
	// https://codereview.stackexchange.com/a/178497
	var maxID string
	for maxID = range us.units { // not really safe, but again - we're small, so it shouldn't cause problems
		break
	}

	maxIDValue, err := strconv.Atoi(maxID)
	if err != nil {
		return app.Unit{}, fmt.Errorf("the max id in the unit's map is not an integer: [%s]", maxID)
	}

	newID := fmt.Sprintf("%d", maxIDValue+1)
	us.units[newID] = unit

	unit.ID = newID

	return unit, nil
}

// this is a private function to populate the "in-memory database".
// Not typical, more proof of concept.
func (us *UnitService) initialize() error {
	us.units = make(map[string]app.Unit)

	us.units["1"] = app.Unit{Element: app.ElementFire, Name: "Fire Tank", UnitType: app.UnitTypeTank}
	us.units["2"] = app.Unit{Element: app.ElementFire, Name: "Fire Healer", UnitType: app.UnitTypeHealer}
	us.units["3"] = app.Unit{Element: app.ElementFire, Name: "Fire Support", UnitType: app.UnitTypeSupport}
	us.units["4"] = app.Unit{Element: app.ElementFire, Name: "Fire Attacker", UnitType: app.UnitTypeAttacker}
	us.units["5"] = app.Unit{Element: app.ElementWater, Name: "Water Tank", UnitType: app.UnitTypeTank}
	us.units["6"] = app.Unit{Element: app.ElementWater, Name: "Water Healer", UnitType: app.UnitTypeHealer}
	us.units["7"] = app.Unit{Element: app.ElementWater, Name: "Water Support", UnitType: app.UnitTypeSupport}
	us.units["8"] = app.Unit{Element: app.ElementWater, Name: "Water Attacker", UnitType: app.UnitTypeAttacker}
	us.units["9"] = app.Unit{Element: app.ElementGrass, Name: "Water Tank", UnitType: app.UnitTypeTank}
	us.units["10"] = app.Unit{Element: app.ElementGrass, Name: "Grass Healer", UnitType: app.UnitTypeHealer}
	us.units["11"] = app.Unit{Element: app.ElementGrass, Name: "Grass Support", UnitType: app.UnitTypeSupport}
	us.units["12"] = app.Unit{Element: app.ElementGrass, Name: "Grass Attacker", UnitType: app.UnitTypeAttacker}
	us.units["13"] = app.Unit{Element: app.ElementNormal, Name: "Normal Tank", UnitType: app.UnitTypeTank}
	us.units["14"] = app.Unit{Element: app.ElementNormal, Name: "Normal Healer", UnitType: app.UnitTypeHealer}
	us.units["15"] = app.Unit{Element: app.ElementNormal, Name: "Normal Support", UnitType: app.UnitTypeSupport}
	us.units["16"] = app.Unit{Element: app.ElementNormal, Name: "Normal Attacker", UnitType: app.UnitTypeAttacker}

	for id, unit := range us.units {
		sg, err := calculateStatsForUnit(unit.Element, unit.UnitType)
		if err != nil {
			return fmt.Errorf("failed to generate stats for unit with id [%s] | %s", id, err)
		}

		unit.Stats = sg
		us.units[id] = unit
	}

	return nil
}

func calculateStatsForUnit(element app.Element, class app.UnitType) (app.StatGroup, error) {
	eStat, ok := elementStatMap[element]
	if !ok {
		return app.StatGroup{}, fmt.Errorf("unknown element: [%v]", element)
	}

	cStat, ok := classStatMap[class]
	if !ok {
		return app.StatGroup{}, fmt.Errorf("unknown class: [%v]", class)
	}

	return app.StatGroup{
		Health:  (eStat.Health + 10) * (cStat.Health + 2),
		Attack:  (eStat.Attack + 3) * (cStat.Attack + 2),
		Defense: (eStat.Defense + 3) * (cStat.Defense + 2),
		Speed:   (eStat.Speed + 3) * (cStat.Speed + 2),
	}, nil
}
