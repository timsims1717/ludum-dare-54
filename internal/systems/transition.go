package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"math"
	"strings"
)

func TransitionSystem() {
	if data.LeaveTransition {
		if data.FadeTween == nil {
			data.FadeTween = gween.New(255., 0, 0.4, ease.Linear)
		}
		if data.TransitionTimer == nil {
			data.TransitionTimer = timing.New(1)
		}
		if data.TransitionTimer.UpdateDone() {
			state.SwitchState(constants.PackingStateKey)
		}
	} else {
		switch data.TransitionStep {
		case 0:
			if data.TransitionTimer == nil {
				data.TransitionTimer = timing.New(2)
			}
			if data.TransitionTimer.UpdateDone() {
				lp := data.CartPositions[len(data.CartPositions)-1]
				ips := []*object.Interpolation{
					object.NewInterpolation(object.InterpolateX).
						AddGween(data.MiniTruckObj.Pos.X, lp.X, 3, ease.Linear),
					object.NewInterpolation(object.InterpolateY).
						AddGween(data.MiniTruckObj.Pos.Y, lp.Y, 3, ease.Linear),
				}
				data.MiniTruckEntity.AddComponent(myecs.Interpolation, ips)
				data.TransitionStep++
				data.TransitionTimer = timing.New(3)
			}
		case 1:
			if data.TransitionTimer.UpdateDone() {
				data.TransitionStep = 0
				data.TransitionTimer = nil
				data.LeaveTransition = true
			}
		}
	}
}

func DrawPaths() {
	data.IMDraw.Clear()
	for i, stall := range data.CartPositions {
		if i+1 < len(data.CartPositions) {
			data.IMDraw.Color = constants.PathColor
			data.IMDraw.Push(stall)
			data.IMDraw.Push(data.CartPositions[i+1])
			data.IMDraw.Line(3)
		}
	}
	data.IMDraw.Draw(data.GameView.Canvas)
}

func AddNewStall() {
	if len(data.CartPositions) == 0 {
		data.CartPositions = append(data.CartPositions, pixel.ZV)
	}
	lp := data.CartPositions[len(data.CartPositions)-1]
	p := pixel.V(1, 0)
	p = p.Rotated((constants.GlobalSeededRandom.Float64() - 0.5) * math.Pi)
	data.CartPositions = append(data.CartPositions, p.Scaled(108).Add(lp))
}

func CreateMiniTruck() {
	if data.MiniTruckObj == nil {
		data.MiniTruckObj = object.New()
		data.MiniTruckObj.Pos = pixel.ZV
		data.MiniTruckObj.Layer = 1
	}
	if data.MiniTruckSpr == nil || strings.Contains(data.MiniTruckSpr.Key, data.CurrentTruck.SpriteKey) {
		str := fmt.Sprintf("%s_tiny", data.CurrentTruck.SpriteKey)
		spr := img.Batchers[constants.TestBatch].GetSprite(str)
		offY := spr.Frame().H() * 0.5
		data.MiniTruckSpr = img.NewOffsetSprite(str, constants.TestBatch, pixel.V(0, offY))
	}
	if data.MiniTruckEntity == nil {
		data.MiniTruckEntity = myecs.Manager.NewEntity()
		data.MiniTruckEntity.AddComponent(myecs.Object, data.MiniTruckObj).
			AddComponent(myecs.Drawable, data.MiniTruckSpr).
			AddComponent(myecs.Update, data.NewTimerFunc(func() bool {
				if data.MiniTruckObj.Offset.Y == 0 {
					data.MiniTruckObj.Offset.Y = 1
				} else {
					data.MiniTruckObj.Offset.Y = 0
				}
				return false
			}, 0.1))
	}
}
