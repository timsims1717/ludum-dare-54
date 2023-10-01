package data

import (
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/typeface"
)

type score struct {
	Cash                 int
	SuccessfulDeliveries int
	MissedDeliveries     int
	FailCondition        int
	AbandonedWares       int
	DeliveryCount        int
}

func NewScore() {
	CurrentScore = &score{}
}

var CurrentScore *score

func CheckForFailure() {
	if CurrentScore.MissedDeliveries >= CurrentDifficulty.NumberofMissedDeliveries {
		CurrentScore.FailCondition = constants.TooManyMisses
	} else if CurrentScore.AbandonedWares >= CurrentDifficulty.NumberofAbandonedWares {
		CurrentScore.FailCondition = constants.AbandonToManyItems
	} else if len(CurrentTruck.Wares) < CurrentDifficulty.TargetWares {
		CurrentScore.FailCondition = constants.TooFewItems
	} else {
		CurrentScore.FailCondition = constants.NotFailing
	}
}

var (
	LeftTitle        *typeface.Text
	LeftComplete     *typeface.Text
	LeftMissed       *typeface.Text
	LeftAbandoned    *typeface.Text
	LeftCash         *typeface.Text
	RightTitle       *typeface.Text
	RightLoadedWares *typeface.Text
	RightLoadHeight  *typeface.Text
	RightPercentFull *typeface.Text
	RightWaresCount  *typeface.Text
	MinWaresCount    *typeface.Text
	TimerCount       *typeface.Text
	ButtonText       *typeface.Text
	ButtonSpr        *img.Sprite
	ButtonObj        *object.Object
	ButtonLock       bool

	BigMessage      *typeface.Text
	BigMessageTimer *timing.Timer
)
