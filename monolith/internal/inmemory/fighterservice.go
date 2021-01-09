package inmemory

import (
	"context"
	"fmt"

	"github.com/mrcampbell/gacha-game/monolith/internal/app"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ensure the interface is implemented correctly
var _ app.FighterService = &FighterService{}

type FighterService struct {
	units map[string]app.Unit
}

// Stats By Element

//				Hp	Atk	Def	Speed	Average
// Fire		1		3		1		3			2
// Water	2		2		2		2			2
// Grass	3		1		3		1			2
// Norm		x		1		1		x			1

// Stats By Class

//					Hp	Atk	Def	Speed	Average
// Tank			5		1		5		1			3
// Healer		2		2		2		4			2
// Support	3		1		2		2			2
// Attacker	2		5		2		4			1

var FireStatMod = app.StatGroup{
	Health:  1,
	Attack:  3,
	Defense: 1,
	Speed:   3,
}

var WaterStatMod = app.StatGroup{
	Health:  2,
	Attack:  2,
	Defense: 2,
	Speed:   2,
}

var GrassStatMod = app.StatGroup{
	Health:  3,
	Attack:  1,
	Defense: 3,
	Speed:   1,
}

func NewFighterService() FighterService {
	units := make(map[string]app.Unit)

	units["1"] = app.Unit{Element: app.ElementFire, Name: "Fire Tank", UnitType: app.UnitTypeTank}
	units["2"] = app.Unit{Element: app.ElementFire, Name: "Fire Healer", UnitType: app.UnitTypeHealer}
	units["3"] = app.Unit{Element: app.ElementFire, Name: "Fire Support", UnitType: app.UnitTypeSupport}
	units["4"] = app.Unit{Element: app.ElementFire, Name: "Fire Attacker", UnitType: app.UnitTypeAttacker}
	units["5"] = app.Unit{Element: app.ElementWater, Name: "Water Tank", UnitType: app.UnitTypeTank}
	units["6"] = app.Unit{Element: app.ElementWater, Name: "Water Healer", UnitType: app.UnitTypeHealer}
	units["7"] = app.Unit{Element: app.ElementWater, Name: "Water Support", UnitType: app.UnitTypeSupport}
	units["8"] = app.Unit{Element: app.ElementWater, Name: "Water Attacker", UnitType: app.UnitTypeAttacker}
	units["9"] = app.Unit{Element: app.ElementGrass, Name: "Water Tank", UnitType: app.UnitTypeTank}
	units["10"] = app.Unit{Element: app.ElementGrass, Name: "Grass Healer", UnitType: app.UnitTypeHealer}
	units["11"] = app.Unit{Element: app.ElementGrass, Name: "Grass Support", UnitType: app.UnitTypeSupport}
	units["12"] = app.Unit{Element: app.ElementGrass, Name: "Grass Attacker", UnitType: app.UnitTypeAttacker}
	units["13"] = app.Unit{Element: app.ElementNormal, Name: "Normal Tank", UnitType: app.UnitTypeTank}
	units["14"] = app.Unit{Element: app.ElementNormal, Name: "Normal Healer", UnitType: app.UnitTypeHealer}
	units["15"] = app.Unit{Element: app.ElementNormal, Name: "Normal Support", UnitType: app.UnitTypeSupport}
	units["16"] = app.Unit{Element: app.ElementNormal, Name: "Normal Attacker", UnitType: app.UnitTypeAttacker}

	return FighterService{units: units}
}

func (f FighterService) UnitByID(ctx context.Context, id string) (app.Unit, error) {
	unit, ok := f.units[id]
	if !ok {
		return app.Unit{}, status.Error(codes.NotFound, fmt.Sprintf("no fighter with id [%s]", id))
	}

	// we purposefully didn't set this in the map above so we could shift around the order without having to worry
	// about inconsistencies between the unit's hardcoded ID and the ID in reference to where it is in the map
	unit.ID = id

	return unit, nil
}

func (f FighterService) ListAllUnits(ctx context.Context) ([]app.Unit, error) {
	units := []app.Unit{}
	for _, unit := range f.units {
		units = append(units, unit)
	}
	return units, nil
}
