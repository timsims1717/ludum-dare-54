package constants

type DifficultyType string

const (
	UndefinedDifficultyType = "Undefined"
	Easy                    = "Easy"
	Medium                  = "Medium"
	Hard                    = "Hard"
)

type Difficulty struct {
	Label                    string
	Level                    DifficultyType
	TimeToDepart             int
	TimeToSell               int
	TrunkTargetFill          int
	NumberofMissedDeliveries int
	NumberofAbandonedWares   int
	TargetWares              int
}

var (
	DifficultyLevels = map[DifficultyType]*Difficulty{
		Easy: {
			Label:                    "Easy",
			Level:                    Easy,
			TimeToDepart:             30,
			TimeToSell:               60,
			TrunkTargetFill:          40,
			NumberofMissedDeliveries: 8,
			NumberofAbandonedWares:   21,
		}, Medium: {
			Label:                    "Medium",
			Level:                    Medium,
			TimeToDepart:             25,
			TimeToSell:               18,
			TrunkTargetFill:          50,
			NumberofMissedDeliveries: 5,
			NumberofAbandonedWares:   13,
		}, Hard: {
			Label:                    "Hard",
			Level:                    Hard,
			TimeToDepart:             20,
			TimeToSell:               15,
			TrunkTargetFill:          60,
			NumberofMissedDeliveries: 3,
			NumberofAbandonedWares:   8,
		},
	}
)
