package data

import (
	"ludum-dare-54/internal/constants"
)

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
	TotalWareSize := 0
	for _, ware := range Wares {
		TotalWareSize += len(constants.TrunkominoTypes[ware.ShapeKey])
	}
	CurrentDifficulty.TargetWares = int(float64((CurrentTruck.Height*CurrentTruck.Depth*CurrentTruck.Width)/(TotalWareSize/len(Wares))) * float64(CurrentDifficulty.InitialTrunkTargetFill) / 100)
}
