package data

import (
	"github.com/faiface/pixel"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/timing"
)

var (
	ItemQueue  [8]*Ware
	HeldItem   *Ware
	BottomDrop pixel.Rect
	LeftDrop   pixel.Rect

	LeavePacking bool
	ScoreTween   *gween.Tween
)

type DragTimer struct {
	Timer *timing.Timer
	Quick bool
}
