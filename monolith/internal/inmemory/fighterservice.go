package inmemory

import (
	"context"
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

	fighter := app.Fighter{Level: level, UserID: userID, BaseUnit: unit, UnitID: unitID}

	fighter.CurrentStats = fighter.CalculateCurrentStats()

	fighter.ID = util.GenerateID()
	fs.fighters[fighter.ID] = fighter

	return fighter, nil
}

func (fs FighterService) FighterByID(ctx context.Context, userID, fighterID string) (app.Fighter, error) {
	fighter, ok := fs.fighters[fighterID]
	if !ok || fighter.UserID != userID { // userID is a simple level of security
		return app.Fighter{}, fmt.Errorf("fighter with id [%s] not found", fighterID)
	}

	moves, err := fs.unitMoveService.ListUnitMovesAtLevelForUnit(ctx, fighter.UnitID, fighter.Level)
	if err != nil {
		return app.Fighter{}, fmt.Errorf("failed to get moves for unit by id | %s", err.Error())
	}

	// we assign moves afterwards because they are their own unique, independent object.  We're only concerned with storing
	// details specific to the fighter, and since our fighter gains the same exact moves as they level up, it will always
	// compute to the same every time.
	// Now, if the fighter could only learn a certain number of moves, and the moves varied between fighters of the same unit
	// this wouldn't work and we would need to create a new service or store it here to ensure that unique mapping persisted.
	fighter.Moves = moves

	return fighter, nil
}

func (fs FighterService) FightersByUserID(ctx context.Context, userID string) ([]app.Fighter, error) {
	fighters := make([]app.Fighter, 0)
	for _, fighter := range fs.fighters { // not efficient, but we're still small, so it's okay
		if fighter.UserID == userID {
			fighters = append(fighters, fighter)
		}
	}

	return fighters, nil
}

func (fs *FighterService) initialize() {
	fs.fighters = make(map[string]app.Fighter)
}
