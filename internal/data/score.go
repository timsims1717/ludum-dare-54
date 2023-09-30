package data

import (
	"ludum-dare-54/internal/constants"
)

type score struct {
	Cash                 int
	SuccessfulDeliveries int
	MissedDeliveries     int
	FailCondition        int
}

func NewScore() {
	CurrentScore = &score{}
}

var CurrentScore *score

func CheckForFailure() {
	if CurrentScore.MissedDeliveries >= 3 {
		CurrentScore.FailCondition = constants.TooManyMisses
	}
}
