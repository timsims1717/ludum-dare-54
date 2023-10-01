package data

import (
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/pkg/typeface"
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
	if CurrentScore.MissedDeliveries >= CurrentDifficulty.NumberofMissedDeliveries {
		CurrentScore.FailCondition = constants.TooManyMisses
	}
}

var (
	LeftCount  *typeface.Text
	RightCount *typeface.Text
	PercCount  *typeface.Text
	TimerCount *typeface.Text
)
