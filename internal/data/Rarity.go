package data

import "ludum-dare-54/internal/constants"

type Rarity string

const (
	UndefinedDifficultyType = "Undefined"
	Common                  = "Common"
	Rare                    = "Rare"
)

func PickRarity() Rarity {
	if constants.GlobalSeededRandom.Intn(100) > 90 {
		return Rare
	}
	return Common
}
