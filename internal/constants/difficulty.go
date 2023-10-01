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
	InitialTrunkTargetFill   int
	NumberofMissedDeliveries int
	NumberofAbandonedWares   int
	TargetWares              int
}

var (
	DifficultyLevels = map[DifficultyType]*Difficulty{
		Easy: {
			Label:                    "Easy",
			Level:                    Easy,
			TimeToDepart:             60,
			TimeToSell:               40,
			InitialTrunkTargetFill:   50,
			NumberofMissedDeliveries: 8,
			NumberofAbandonedWares:   21,
		}, Medium: {
			Label:                    "Medium",
			Level:                    Medium,
			TimeToDepart:             50,
			TimeToSell:               30,
			InitialTrunkTargetFill:   40,
			NumberofMissedDeliveries: 5,
			NumberofAbandonedWares:   13,
		}, Hard: {
			Label:                    "Hard",
			Level:                    Hard,
			TimeToDepart:             40,
			TimeToSell:               25,
			InitialTrunkTargetFill:   30,
			NumberofMissedDeliveries: 3,
			NumberofAbandonedWares:   8,
		},
	}
)
