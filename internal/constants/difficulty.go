package constants

type DifficultyType string

const (
	UndefinedDifficultyType = "Undefined"
	Easy                    = "Easy"
	Medium                  = "Medium"
	Hard                    = "Hard"
)

type Difficulty struct {
	Level                    DifficultyType
	TimeToDepart             int
	TimeToSell               int
	InitialTrunkTargetFill   int
	NumberofMissedDeliveries int
	NumberofAbandonedWares   int
}

var (
	DifficultyLevels = map[DifficultyType]*Difficulty{
		Easy: {
			Level:                    Easy,
			TimeToDepart:             100,
			TimeToSell:               100,
			InitialTrunkTargetFill:   50,
			NumberofMissedDeliveries: 8,
			NumberofAbandonedWares:   21,
		}, Medium: {
			Level:                    Medium,
			TimeToDepart:             80,
			TimeToSell:               80,
			InitialTrunkTargetFill:   40,
			NumberofMissedDeliveries: 5,
			NumberofAbandonedWares:   13,
		}, Hard: {
			Level:                    Hard,
			TimeToDepart:             60,
			TimeToSell:               60,
			InitialTrunkTargetFill:   30,
			NumberofMissedDeliveries: 3,
			NumberofAbandonedWares:   8,
		},
	}
)
