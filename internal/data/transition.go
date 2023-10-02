package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/timing"
)

var (
	LeaveTransition bool
	TransitionTimer *timing.Timer
	TransitionStep  int

	CartPositions   []pixel.Vec
	MiniTruckSpr    *img.Sprite
	MiniTruckObj    *object.Object
	MiniTruckEntity *ecs.Entity
)
