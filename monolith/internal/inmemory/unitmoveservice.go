package inmemory

import (
	"context"
	"fmt"

	"github.com/mrcampbell/gacha-game/monolith/internal/app"
)

// ensure the interface is implemented correctly
var _ app.UnitMoveService = &UnitMoveService{}

type UnitMoveService struct {
	unitMoves   map[string]app.UnitMove
	moveService app.MoveService
}

func NewUnitMoveService(moveService app.MoveService) UnitMoveService {
	ums := UnitMoveService{moveService: moveService}
	ums.initialize()
	return ums
}

func (ums UnitMoveService) ListUnitMovesForUnit(ctx context.Context, unitID string) ([]app.Move, error) {
	moves := make([]app.Move, 0)
	for _, unitMove := range ums.unitMoves {
		if unitMove.UnitID == unitID {
			move, err := ums.moveService.MoveByID(ctx, unitMove.MoveID)
			if err != nil {
				return nil, fmt.Errorf("failed getting move from move service by id: [%s]", unitMove.MoveID)
			}
			moves = append(moves, move)
		}
	}

	if len(moves) == 0 { // this mean we don't have any moves for the user
		return nil, fmt.Errorf("no unit moves for unit with id: [%s]", unitID)
	}

	return moves, nil
}

func (ums UnitMoveService) ListUnitMovesAtLevelForUnit(ctx context.Context, unitID string, level int32) ([]app.Move, error) {
	moves := make([]app.Move, 0)
	for _, unitMove := range ums.unitMoves {
		if unitMove.UnitID == unitID && unitMove.LevelLearnedAt <= level {
			move, err := ums.moveService.MoveByID(ctx, unitMove.MoveID)
			if err != nil {
				return nil, fmt.Errorf("failed getting move from move service by id: [%s]", unitMove.MoveID)
			}
			moves = append(moves, move)
		}
	}

	if len(moves) == 0 { // this mean we don't have any moves for the user
		return nil, fmt.Errorf("no unit moves for unit with id: [%s]", unitID)
	}

	return moves, nil
}

// this is a private function to populate the "in-memory database".
// Not typical, more proof of concept.
func (ums *UnitMoveService) initialize() {
	ums.unitMoves = make(map[string]app.UnitMove)

	// Fire Tank
	ums.unitMoves["1"] = app.UnitMove{MoveID: "1", UnitID: "1", LevelLearnedAt: 0}  // Smack
	ums.unitMoves["2"] = app.UnitMove{MoveID: "11", UnitID: "1", LevelLearnedAt: 0} // Singe
	ums.unitMoves["3"] = app.UnitMove{MoveID: "3", UnitID: "1", LevelLearnedAt: 5}  // Boast

	// Water Support
	ums.unitMoves["4"] = app.UnitMove{MoveID: "1", UnitID: "7", LevelLearnedAt: 0}  // Smack
	ums.unitMoves["5"] = app.UnitMove{MoveID: "4", UnitID: "7", LevelLearnedAt: 0}  // Splash
	ums.unitMoves["6"] = app.UnitMove{MoveID: "10", UnitID: "7", LevelLearnedAt: 0} // Sleet

	// Grass Healer
	ums.unitMoves["7"] = app.UnitMove{MoveID: "1", UnitID: "10", LevelLearnedAt: 0} // Smack
	ums.unitMoves["8"] = app.UnitMove{MoveID: "5", UnitID: "10", LevelLearnedAt: 0} // Root
	ums.unitMoves["9"] = app.UnitMove{MoveID: "6", UnitID: "10", LevelLearnedAt: 0} // Shade

	// Normal Attacker
	ums.unitMoves["10"] = app.UnitMove{MoveID: "1", UnitID: "16", LevelLearnedAt: 0}  // Smack
	ums.unitMoves["11"] = app.UnitMove{MoveID: "8", UnitID: "16", LevelLearnedAt: 0}  // Intimidate
	ums.unitMoves["12"] = app.UnitMove{MoveID: "12", UnitID: "16", LevelLearnedAt: 0} // Slam
}
