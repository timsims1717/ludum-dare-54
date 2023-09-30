package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/world"
)

type Item struct {
	Key    string
	Name   string
	Index  int
	Shape  []world.Coords
	Object *object.Object
	Entity *ecs.Entity
	Sprite *img.Sprite
}

var (
	ItemQueue  [8]*Item
	HeldItem   *Item
	BottomDrop pixel.Rect
	LeftDrop   pixel.Rect
)

type DragTimer struct {
	Timer *timing.Timer
	Quick bool
}
