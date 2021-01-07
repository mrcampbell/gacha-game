package inmemory

import (
	"context"
	"fmt"

	"github.com/mrcampbell/gacha-game/monolith/internal/app"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

type FighterService struct {
	units map[string]app.Unit
}

func NewFighterService() FighterService {
	units := make(map[string]app.Unit)

	units["1"] = app.Unit{Element: app.ElementFire, ID: "1", Name: "Fire Tank", UnitType: app.UnitTypeTank}
	units["2"] = app.Unit{Element: app.ElementFire, ID: "2", Name: "Fire Healer", UnitType: app.UnitTypeHealer}

	return FighterService{units: units}
}

func (f FighterService) UnitByID(ctx context.Context, id string) (app.Unit, error) {
	unit, ok := f.units[id]
	if !ok {
		return app.Unit{}, status.Error(codes.NotFound, fmt.Sprintf("no fighter with id [%s]", id))
	}
	return unit, nil
}

func (f FighterService) ListAllUnits(ctx context.Context) ([]app.Unit, error) {
	units := []app.Unit{}
	for _, unit := range f.units {
		units = append(units, unit)
	}
	return units, nil
}
