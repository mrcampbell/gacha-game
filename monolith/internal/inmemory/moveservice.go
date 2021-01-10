package inmemory

import (
	"context"
	"fmt"

	"github.com/mrcampbell/gacha-game/monolith/internal/app"
)

// ensure the interface is implemented correctly
var _ app.MoveService = &MoveService{}

type MoveService struct {
	moves map[string]app.Move
}

func NewMoveService() MoveService {
	ms := MoveService{}
	ms.initialize()
	return ms
}

func (ms MoveService) MoveByID(ctx context.Context, id string) (app.Move, error) {
	move, ok := ms.moves[id]
	if !ok {
		return app.Move{}, fmt.Errorf("no move with id [%s]", id)
	}

	// we purposefully didn't set this in the map above so we could shift around the order without having to worry
	// about inconsistencies between the move's hardcoded ID and the ID in reference to where it is in the map
	move.ID = id

	return move, nil
	// return app.Move{}, errors.New("unimplemented")
}

// this is a private function to populate the "in-memory database".
// Not typical, more proof of concept.
func (ms *MoveService) initialize() {
	ms.moves = make(map[string]app.Move)

	smack := app.Move{
		Name:        "Smack",
		Description: "Slap an enemy to inflict damage",
		Element:     app.ElementNormal,
		Metas: []app.MoveMeta{
			app.NewMoveMetaDamage(app.MoveTargetTypeSingleEnemy, 3),
		},
		RestDuration: 0,
	}

	meditate := app.Move{
		Name:        "Meditate",
		Description: "Center yourself and heal 25% of Max Health",
		Element:     app.ElementNormal,
		Metas: []app.MoveMeta{
			app.NewMoveMetaHeal(app.MoveTargetTypeSelf, 10),
		},
		RestDuration: 3,
	}

	boast := app.Move{
		Name:        "Boast",
		Description: "Draw attention to yourself for 3 turns, but boost your defense greatly for only 2 turns",
		Element:     app.ElementNormal,
		Metas: []app.MoveMeta{
			app.NewMoveMetaLockOnMe(3),
			app.NewMoveMetaStatChange(app.MoveTargetTypeSelf, app.StatDefense, 3, 2), // leave them vulnerable for the last turn
		},
		RestDuration: 6,
	}

	splash := app.Move{
		Name:        "Splash",
		Description: "Actually deal some water damage",
		Element:     app.ElementWater,
		Metas: []app.MoveMeta{
			app.NewMoveMetaDamage(app.MoveTargetTypeSingleEnemy, 2),
		},
		RestDuration: 1,
	}

	root := app.Move{
		Name:        "Root",
		Description: "Dig into the Earth with roots, and heal 25%",
		Element:     app.ElementGrass,
		Metas: []app.MoveMeta{
			app.NewMoveMetaHeal(app.MoveTargetTypeSelf, 25),
		},
		RestDuration: 4,
	}

	shade := app.Move{
		Name:        "Shade",
		Description: "Extend vines and leaves over allies and grant moderate defense and healing for 2 turns",
		Element:     app.ElementGrass,
		Metas: []app.MoveMeta{
			app.NewMoveMetaStatChange(app.MoveTargetTypeAllAllies, app.StatDefense, 2, 2),
			app.NewMoveMetaHeal(app.MoveTargetTypeAllAllies, 15),
		},
		RestDuration: 6,
	}

	scatterShot := app.Move{
		Name:        "Scatter Shot",
		Description: "Fire seeds in every enemys' direction",
		Element:     app.ElementGrass,
		Metas: []app.MoveMeta{
			app.NewMoveMetaStatChange(app.MoveTargetTypeAllAllies, app.StatDefense, 2, 2),
		},
		RestDuration: 3,
	}

	fireUp := app.Move{
		Name:        "Fire-up",
		Description: "Excite your team and up their attack, defense and speed a bit for 2 turns",
		Element:     app.ElementFire,
		Metas: []app.MoveMeta{
			app.NewMoveMetaStatChange(app.MoveTargetTypeAllAllies, app.StatAttack, 1, 2),
			app.NewMoveMetaStatChange(app.MoveTargetTypeAllAllies, app.StatDefense, 1, 2),
			app.NewMoveMetaStatChange(app.MoveTargetTypeAllAllies, app.StatSpeed, 1, 2),
		},
		RestDuration: 6,
	}

	intimidate := app.Move{
		Name:        "Intimidate",
		Description: "Scowl at your enemies and drop their defense a bit for 3 turns",
		Element:     app.ElementNormal,
		Metas: []app.MoveMeta{
			app.NewMoveMetaStatChange(app.MoveTargetTypeAllEnemies, app.StatDefense, -2, 3),
		},
	}

	sleet := app.Move{
		Name:        "Sleet",
		Description: "Shoot icy rain at an enemy to deal damage and reduce their speed",
		Element:     app.ElementWater, // ice?
		Metas: []app.MoveMeta{
			app.NewMoveMetaStatChange(app.MoveTargetTypeSingleAlly, app.StatSpeed, -2, 2),
			app.NewMoveMetaDamage(app.MoveTargetTypeSingleEnemy, 3),
		},
	}

	singe := app.Move{
		Name:        "Singe",
		Description: "Graze an opponent with a hot hand, causing fire damage",
		Element:     app.ElementFire,
		Metas: []app.MoveMeta{
			app.NewMoveMetaDamage(app.MoveTargetTypeSingleEnemy, 2),
		},
	}

	slam := app.Move{
		Name:        "Slam",
		Description: "Deal massive damage to one enemy",
		Element:     app.ElementNormal,
		Metas: []app.MoveMeta{
			app.NewMoveMetaDamage(app.MoveTargetTypeSingleEnemy, 6),
		},
	}

	ms.moves["1"] = smack
	ms.moves["2"] = meditate
	ms.moves["3"] = boast
	ms.moves["4"] = splash
	ms.moves["5"] = root
	ms.moves["6"] = shade
	ms.moves["7"] = fireUp
	ms.moves["8"] = intimidate
	ms.moves["9"] = scatterShot
	ms.moves["10"] = sleet
	ms.moves["11"] = singe
	ms.moves["12"] = slam
}
