package mechanics

import (
	"fmt"

	"github.com/mrcampbell/gacha-game/battle-service/internal/app"
)

type AdvantageFactor float32

const (
	EFFECTIVE   = 2.0
	NORMAL      = 1.0
	INEFFECTIVE = 0.5
)

var advantageAttackerDefenderFactorMap map[app.Element]map[app.Element]float32

func init() {
	advantageAttackerDefenderFactorMap = make(map[app.Element]map[app.Element]float32)
	advantageAttackerDefenderFactorMap[app.FIRE] = make(map[app.Element]float32)

	advantageAttackerDefenderFactorMap[app.FIRE][app.FIRE] = NORMAL
	advantageAttackerDefenderFactorMap[app.FIRE][app.WATER] = INEFFECTIVE
	advantageAttackerDefenderFactorMap[app.FIRE][app.GRASS] = EFFECTIVE
}

func ElementalAdvantageFactor(attacker, defender app.Element) (float32, error) {

	defenderFactorMap, ok := advantageAttackerDefenderFactorMap[attacker]

	if ok == false {
		return 0.0, fmt.Errorf("attacker matchup map not found")
	}

	factor, ok := defenderFactorMap[defender]
	if ok == false {
		return 0.0, fmt.Errorf("defender matchup map not found")
	}

	return factor, nil

	// switch attacker {
	// case app.FIRE:
	// 	{
	// 		switch defender {
	// 		case app.FIRE:
	// 			return NORMAL, nil
	// 		case app.WATER:
	// 			return INEFFECTIVE, nil
	// 		case app.GRASS:
	// 			return EFFECTIVE, nil
	// 		}
	// 	}

	// case app.WATER:
	// 	{
	// 		switch defender {
	// 		case app.FIRE:
	// 			return NORMAL, nil
	// 		case app.WATER:
	// 			return INEFFECTIVE, nil
	// 		case app.GRASS:
	// 			return EFFECTIVE, nil
	// 		}
	// 	}
	// case app.GRASS:
	// 	{
	// 		switch defender {
	// 		case app.FIRE:
	// 			return NORMAL, nil
	// 		case app.WATER:
	// 			return INEFFECTIVE, nil
	// 		case app.GRASS:
	// 			return EFFECTIVE, nil
	// 		}
	// 	}
	// }
	// return 0.0, fmt.Errorf("Invalid element. Attacker: [%s] Defender: [%s]", attacker, defender)
}
