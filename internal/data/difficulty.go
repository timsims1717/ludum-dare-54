package data

import "ludum-dare-54/internal/constants"

var (
	CurrentDifficulty *constants.Difficulty
)

func SetDifficulty(d constants.DifficultyType) {
	CurrentDifficulty = &constants.Difficulty{
		Level:                    constants.DifficultyLevels[d].Level,
		TimeToDepart:             constants.DifficultyLevels[d].TimeToDepart,
		TimeToSell:               constants.DifficultyLevels[d].TimeToSell,
		InitialTrunkTargetFill:   constants.DifficultyLevels[d].InitialTrunkTargetFill,
		NumberofMissedDeliveries: constants.DifficultyLevels[d].NumberofMissedDeliveries,
		NumberofAbandonedWares:   constants.DifficultyLevels[d].NumberofAbandonedWares,
	}
}
