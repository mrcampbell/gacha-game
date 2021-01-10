package inmemory

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrcampbell/gacha-game/monolith/internal/util"

	"github.com/mrcampbell/gacha-game/monolith/internal/app"
)

// ensure the interface is implemented correctly
var _ app.FighterService = &FighterService{}

type FighterService struct {
	fighters        map[string]app.Fighter
	unitService     app.UnitService
	unitMoveService app.UnitMoveService
}

func NewFighterService(unitService app.UnitService, unitMoveService app.UnitMoveService) FighterService {
	fs := FighterService{unitService: unitService, unitMoveService: unitMoveService}
	fs.initialize()
	return fs
}

func (fs FighterService) CreateFighter(ctx context.Context, userID, unitID string, level int32) (app.Fighter, error) {
	unit, err := fs.unitService.UnitByID(ctx, unitID)
	if err != nil {
		return app.Fighter{}, fmt.Errorf("failed to get unit by id | %s", err.Error())
	}

	moves, err := fs.unitMoveService.ListUnitMovesAtLevelForUnit(ctx, unitID, level)
	if err != nil {
		return app.Fighter{}, fmt.Errorf("failed to get moves for unit by id | %s", err.Error())
	}

	fighter := app.Fighter{Level: level, UserID: userID, BaseUnit: unit, UnitID: unitID}

	fighter.CurrentStats = fighter.CalculateCurrentStats()

	fighter.ID = util.GenerateID()
	fs.fighters[fighter.ID] = fighter

	// we assign moves afterwards because they are their own unique, independent object.  We're only concerned with storing
	// details specific to the fighter, and since our fighter gains the same exact moves as they level up, it will always
	// compute to the same every time.
	// Now, if the fighter could only learn a certain number of moves, and the moves varied between fighters of the same unit
	// this wouldn't work and we would need to create a new service or store it here to ensure that unique mapping persisted.
	fighter.Moves = moves

	return fighter, nil
}

func (fs FighterService) FighterByID(ctx context.Context, userID, fighterID string) (app.Fighter, error) {
	return app.Fighter{}, errors.New("unimplemented")
}

func (fs FighterService) FightersByUserID(ctx context.Context, userID string) ([]app.Fighter, error) {
	return nil, errors.New("unimplemented")
}

func (fs *FighterService) initialize() {
	fs.fighters = make(map[string]app.Fighter)
}
