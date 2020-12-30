package mechanics

import (
	"fmt"

	"github.com/mrcampbell/gacha-game/battle-service/internal/app"
	_ "github.com/mrcampbell/gacha-game/battle-service/internal/app"
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
	advantageAttackerDefenderFactorMap[app.ElementFire] = make(map[app.Element]float32)

	advantageAttackerDefenderFactorMap[app.ElementFire][app.ElementFire] = NORMAL
	advantageAttackerDefenderFactorMap[app.ElementFire][app.ElementWater] = INEFFECTIVE
	advantageAttackerDefenderFactorMap[app.ElementFire][app.ElementGrass] = EFFECTIVE
}

func ElementalAdvantageFactor(attacker, defender app.Element) (float32, error) {

	defenderFactorMap, ok := advantageAttackerDefenderFactorMap[attacker]

	if !ok {
		return 0.0, fmt.Errorf("attacker matchup map not found")
	}

	factor, ok := defenderFactorMap[defender]
	if !ok {
		return 1.0, nil // this is the default response.
	}

	return factor, nil

	// switch attacker {
	// case app.ElementFire:
	// 	{
	// 		switch defender {
	// 		case app.ElementFire:
	// 			return NORMAL, nil
	// 		case app.ElementWater:
	// 			return INEFFECTIVE, nil
	// 		case app.ElementGrass:
	// 			return EFFECTIVE, nil
	// 		}
	// 	}

	// case app.ElementWater:
	// 	{
	// 		switch defender {
	// 		case app.ElementFire:
	// 			return NORMAL, nil
	// 		case app.ElementWater:
	// 			return INEFFECTIVE, nil
	// 		case app.ElementGrass:
	// 			return EFFECTIVE, nil
	// 		}
	// 	}
	// case app.ElementGrass:
	// 	{
	// 		switch defender {
	// 		case app.ElementFire:
	// 			return NORMAL, nil
	// 		case app.ElementWater:
	// 			return INEFFECTIVE, nil
	// 		case app.ElementGrass:
	// 			return EFFECTIVE, nil
	// 		}
	// 	}
	// }
	// return 0.0, fmt.Errorf("Invalid element. Attacker: [%s] Defender: [%s]", attacker, defender)
}
