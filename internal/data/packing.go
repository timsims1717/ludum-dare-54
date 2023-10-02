package data

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/world"
)

var (
	WareQueue      [8]*Ware
	AbandonedWares []*Ware
	HeldWare       *Ware
	SellWares      []*Ware
	BuyWares       int
	BottomDrop     pixel.Rect
	LeftDrop       pixel.Rect
	DepartureTimer *timing.Timer
	IsTimer        bool
	SignObj        *object.Object
	Sign           *ecs.Entity

	LeavePacking bool
	ScoreTween   *gween.Tween
	SignTween    *gween.Tween
	LeaveTimer   *timing.Timer
	TruckTween   *gween.Tween
	LeaveStep    int
	FadeTween    *gween.Tween
	FirstLoad    = true
	Abandon      bool
	Starting     bool
	Paused       bool

	BackgroundItems []*ecs.Entity
	TrunkEntities   []*ecs.Entity
	TrunkObjects    []*object.Object
)

type DragTimer struct {
	Timer *timing.Timer
	Quick bool
}

func AddToAbandoned(ware *Ware) {
	for _, w := range AbandonedWares {
		if w.SpriteKey == ware.SpriteKey {
			return
		}
	}
	AbandonedWares = append(AbandonedWares, ware)
}

func GetFromAbandoned() *Ware {
	if len(AbandonedWares) == 0 {
		return nil
	} else {
		index := constants.GlobalSeededRandom.Intn(len(AbandonedWares))
		rWare := AbandonedWares[index]
		if len(AbandonedWares) > 1 {
			AbandonedWares = append(AbandonedWares[:index], AbandonedWares[index+1:]...)
		} else {
			AbandonedWares = []*Ware{}
		}
		return rWare.CopyWare()
	}
}

func NotInQueue(ware *Ware) bool {
	for _, w := range WareQueue {
		if w != nil && w.SpriteKey == ware.SpriteKey {
			return false
		}
	}
	return true
}

func NotTooDeep(ware *Ware) bool {
	return world.Height(constants.TrunkominoTypes[ware.ShapeKey]) <= CurrentTruck.Depth
}
