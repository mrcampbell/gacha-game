package inmemory

import (
	"context"
	"errors"

	"github.com/mrcampbell/gacha-game/monolith/internal/app"
)

// ensure the interface is implemented correctly
var _ app.MoveService = &MoveService{}

type MoveService struct {
	moves map[string]app.Move
}

func (ms MoveService) MoveByID(ctx context.Context, id string) (app.Move, error) {
	return app.Move{}, errors.New("unimplemented")
}

// this is a private function to populate the "in-memory database".
// Not typical, more proof of concept.
func (ms MoveService) initialize() {
	ms.moves = make(map[string]app.Move)

	smack := app.Move{
		Name:        "Smack",
		Description: "Slap an enemy to inflict damage",
		Element:     app.ElementNormal,
		Metas: []app.MoveMeta{
			app.NewMoveMetaDamage(app.MoveTargetTypeSingleEnemy, 3),
		},
	}

	meditate := app.Move{
		Name:        "Meditate",
		Description: "Center yourself and heal 25% of Max Health",
		Element:     app.ElementNormal,
		Metas: []app.MoveMeta{
			app.NewMoveMetaHeal(app.MoveTargetTypeSelf, 10),
		},
	}

	boast := app.Move{
		Name:        "Boast",
		Description: "Draw attention to yourself for 3 turns, boost your defense for 2",
		Element:     app.ElementNormal,
		Metas: []app.MoveMeta{
			app.NewMoveMetaLockOnMe(3),
			app.NewMoveMetaStatChange(app.MoveTargetTypeSelf, app.StatDefense, 2, 2), // leave them vulnerable for the last turn
		},
	}

	splash := app.Move{
		Name:        "Splash",
		Description: "Actually deal some water damage",
		Element:     app.ElementWater,
		Metas: []app.MoveMeta{
			app.NewMoveMetaDamage(app.MoveTargetTypeSingleEnemy, 2),
		},
	}

	ms.moves["1"] = smack
	ms.moves["2"] = meditate
	ms.moves["3"] = boast
	ms.moves["4"] = splash

}
