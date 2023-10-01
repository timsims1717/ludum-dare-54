package data

import (
	"ludum-dare-54/internal/constants"
)

var (
	CurrentDifficulty *constants.Difficulty
	TotalWareSize     int
	PickedDiffKey     constants.DifficultyType
)

func SetTotalWareSize() {
	TotalWareSize = 0
	for _, ware := range Wares {
		TotalWareSize += len(constants.TrunkominoTypes[ware.ShapeKey])
	}
}

func SetDifficulty(d constants.DifficultyType) {
	CurrentDifficulty = &constants.Difficulty{
		Level:                    constants.DifficultyLevels[d].Level,
		TimeToDepart:             constants.DifficultyLevels[d].TimeToDepart,
		TimeToSell:               constants.DifficultyLevels[d].TimeToSell,
		TrunkTargetFill:          constants.DifficultyLevels[d].TrunkTargetFill,
		NumberofMissedDeliveries: constants.DifficultyLevels[d].NumberofMissedDeliveries,
		NumberofAbandonedWares:   constants.DifficultyLevels[d].NumberofAbandonedWares,
	}
	CurrentDifficulty.TargetWares = TargetWares(CurrentTruck, CurrentDifficulty)
}

func TargetWares(truck *Truck, d *constants.Difficulty) int {
	return int(float64((truck.Height*truck.Depth*truck.Width)/(TotalWareSize/len(Wares))) * float64(d.TrunkTargetFill) / 100)
}
