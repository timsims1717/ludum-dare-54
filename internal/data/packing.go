package data

import (
	"github.com/faiface/pixel"
	"ludum-dare-54/pkg/timing"
)

var (
	ItemQueue  [8]*Ware
	HeldItem   *Ware
	BottomDrop pixel.Rect
	LeftDrop   pixel.Rect
)

type DragTimer struct {
	Timer *timing.Timer
	Quick bool
}
