package data

import (
	"github.com/faiface/pixel"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/timing"
)

var (
	WareQueue      [8]*Ware
	HeldWare       *Ware
	SellWares      []*Ware
	BuyWares       int
	BottomDrop     pixel.Rect
	LeftDrop       pixel.Rect
	DepartureTimer *timing.Timer
	IsTimer        bool

	LeavePacking bool
	ScoreTween   *gween.Tween
	LeaveTimer   *timing.Timer
	LeaveStep    int
	FadeTween    *gween.Tween
	FirstLoad    = true
	Starting     bool
)

type DragTimer struct {
	Timer *timing.Timer
	Quick bool
}
