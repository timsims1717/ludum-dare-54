package data

import (
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/typeface"
)

type score struct {
	Cash                 int
	SuccessfulDeliveries int
	MissedDeliveries     int
	FailCondition        int
	AbandonedWares       int
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
	LeftTitle        *typeface.Text
	LeftCompletes    *typeface.Text
	LeftMisseds      *typeface.Text
	LeftAbandoned    *typeface.Text
	LeftCash         *typeface.Text
	RightTitle       *typeface.Text
	RightLoadedWares *typeface.Text
	RightLoadHeight  *typeface.Text
	RightPercentFull *typeface.Text
	PercCount        *typeface.Text
	TimerCount       *typeface.Text
	ButtonText       *typeface.Text
	ButtonSpr        *img.Sprite
	ButtonObj        *object.Object
	ButtonLock       bool
)
