package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/typeface"
)

var (
	leftScoreboard   = -200.
	rightScoreboard  = -leftScoreboard
	centerScoreboard = 0.
	topScoreboard    = 170.
	lineHeight       = 30.
	middleScoreboard = -20.
	buttonYHeight    = -100.
)

func ScoreboardInit() {
	data.LeftTitle = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftTitle.Obj.Layer = 30
	data.LeftTitle.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftTitle.SetPos(pixel.V(leftScoreboard, topScoreboard))
	data.LeftTitle.SetText("DELIVERIES")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftTitle.Obj).
		AddComponent(myecs.Drawable, data.LeftTitle).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.LeftComplete = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftComplete.Obj.Layer = 30
	data.LeftComplete.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftComplete.SetPos(pixel.V(leftScoreboard, topScoreboard-lineHeight))
	data.LeftComplete.SetText(fmt.Sprintf("%d Complete", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftComplete.Obj).
		AddComponent(myecs.Drawable, data.LeftComplete).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.LeftMissed = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftMissed.Obj.Layer = 30
	data.LeftMissed.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftMissed.SetPos(pixel.V(leftScoreboard, topScoreboard-lineHeight*2))
	data.LeftMissed.SetText(fmt.Sprintf("%d Missed", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftMissed.Obj).
		AddComponent(myecs.Drawable, data.LeftMissed).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.LeftAbandoned = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftAbandoned.Obj.Layer = 30
	data.LeftAbandoned.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftAbandoned.SetPos(pixel.V(leftScoreboard, topScoreboard-lineHeight*3))
	data.LeftAbandoned.SetText(fmt.Sprintf("%d Abandonded", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftAbandoned.Obj).
		AddComponent(myecs.Drawable, data.LeftAbandoned).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.LeftCash = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftCash.Obj.Layer = 30
	data.LeftCash.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftCash.SetPos(pixel.V(leftScoreboard, topScoreboard-lineHeight*4))
	data.LeftCash.SetText(fmt.Sprintf("$%d", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftCash.Obj).
		AddComponent(myecs.Drawable, data.LeftCash).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightTitle = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightTitle.Obj.Layer = 30
	data.RightTitle.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightTitle.SetPos(pixel.V(rightScoreboard, topScoreboard))
	data.RightTitle.SetText(data.CurrentTruck.TruckLabel)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightTitle.Obj).
		AddComponent(myecs.Drawable, data.RightTitle).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightLoadedWares = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightLoadedWares.Obj.Layer = 30
	data.RightLoadedWares.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightLoadedWares.SetPos(pixel.V(rightScoreboard, topScoreboard-lineHeight))
	data.RightLoadedWares.SetText(fmt.Sprintf("Loaded Wares: 42"))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightLoadedWares.Obj).
		AddComponent(myecs.Drawable, data.RightLoadedWares).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightLoadHeight = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightLoadHeight.Obj.Layer = 30
	data.RightLoadHeight.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightLoadHeight.SetPos(pixel.V(rightScoreboard, topScoreboard-lineHeight*2))
	data.RightLoadHeight.SetText(fmt.Sprintf("Load Height: %d / %d", 0, data.CurrentTruck.Depth))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightLoadHeight.Obj).
		AddComponent(myecs.Drawable, data.RightLoadHeight).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightPercentFull = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightPercentFull.Obj.Layer = 30
	data.RightPercentFull.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightPercentFull.SetPos(pixel.V(rightScoreboard, topScoreboard-lineHeight*3))
	data.RightPercentFull.SetText(fmt.Sprintf("%d%% Full", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightPercentFull.Obj).
		AddComponent(myecs.Drawable, data.RightPercentFull).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.PercCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.2, 300., 0.)
	data.PercCount.Obj.Layer = 30
	data.PercCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.PercCount.SetPos(pixel.V(centerScoreboard, middleScoreboard))
	data.PercCount.SetText("Min Wares xx / yy")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.PercCount.Obj).
		AddComponent(myecs.Drawable, data.PercCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.TimerCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
	data.TimerCount.Obj.Layer = 30
	data.TimerCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.TimerCount.SetPos(pixel.V(centerScoreboard, middleScoreboard))
	data.TimerCount.SetText("16.012")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.TimerCount.Obj).
		AddComponent(myecs.Drawable, data.TimerCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.ButtonText = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.24, 300., 0.)
	data.ButtonText.Obj.Layer = 30
	data.ButtonText.SetColor(pixel.ToRGBA(colornames.Black))
	data.ButtonText.SetPos(pixel.V(centerScoreboard, buttonYHeight))
	data.ButtonText.SetText("Set Out")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.ButtonText.Obj).
		AddComponent(myecs.Drawable, data.ButtonText).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.ButtonLock = true
	clickedOn := false
	data.ButtonSpr = img.NewSprite("button", constants.TestBatch)
	spr := img.Batchers[constants.TestBatch].GetSprite("button")
	data.ButtonObj = object.New()
	data.ButtonObj.SetRect(spr.Frame())
	data.ButtonObj.Layer = 29
	data.ButtonObj.Pos = pixel.V(centerScoreboard, buttonYHeight-15.)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.ButtonObj).
		AddComponent(myecs.Drawable, data.ButtonSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.ScoreView, func(hvc *data.HoverClick) {
			if !data.ButtonLock {
				click := hvc.Input.Get("click")
				data.ButtonSpr.Color = pixel.RGB(1, 1, 1)
				data.ButtonText.SetColor(pixel.RGB(0, 0, 0))
				if hvc.Hover {
					if click.JustPressed() {
						clickedOn = true
					}
					if click.Pressed() && clickedOn {
						data.ButtonSpr.Color = pixel.RGB(0.9, 0.9, 0.9)
					} else if click.JustReleased() && clickedOn {
						// leave the packing state
						data.LeavePacking = true
					}
				} else if clickedOn && !click.Pressed() {
					clickedOn = false
				}
			} else {
				data.ButtonText.SetColor(pixel.RGB(0.7, 0.7, 0.7))
			}
		}))
}

func ScoreboardReset() {
	data.TimerCount.SetColor(constants.BaseUIText)
	data.LeftTitle.SetColor(constants.BaseUIText)
	data.LeftComplete.SetColor(constants.BaseUIText)
	data.LeftMissed.SetColor(constants.BaseUIText)
	data.LeftAbandoned.SetColor(constants.BaseUIText)
	data.LeftCash.SetColor(constants.BaseUIText)
	data.RightTitle.SetColor(constants.BaseUIText)
	data.RightLoadedWares.SetColor(constants.BaseUIText)
	data.RightLoadHeight.SetColor(constants.BaseUIText)
	data.RightPercentFull.SetColor(constants.BaseUIText)
	data.PercCount.SetColor(constants.BaseUIText)
	data.TimerCount.SetColor(constants.BaseUIText)
	data.ButtonText.SetColor(constants.BaseUIText)
}

func ScoreSystem() {
	data.RightLoadedWares.SetText(fmt.Sprintf("Loaded Wares: %d", len(data.CurrentTruck.Wares)))
	data.RightLoadHeight.SetText(fmt.Sprintf("Load Height: %d / %d", data.CurrentTruck.CurrHeight, data.CurrentTruck.Height))
	//data.LeftTitle.SetText(fmt.Sprintf("DELIVERIES\n%d Complete\n%d Missed\n$%d.00", data.CurrentScore.SuccessfulDeliveries,
	//data.CurrentScore.MissedDeliveries, data.CurrentScore.Cash))
	data.LeftComplete.SetText(fmt.Sprintf("%d Complete", data.CurrentScore.SuccessfulDeliveries))
	data.LeftMissed.SetText(fmt.Sprintf("%d Missed", data.CurrentScore.MissedDeliveries))
	data.LeftAbandoned.SetText(fmt.Sprintf("%d Abandonded", data.CurrentScore.AbandonedWares))
	data.LeftCash.SetText(fmt.Sprintf("$%d", data.CurrentScore.Cash))
	data.PercCount.SetText(fmt.Sprintf("Min Wares: %d / %d", len(data.CurrentTruck.Wares), data.CurrentDifficulty.TargetWares))
	data.RightPercentFull.SetText(fmt.Sprintf("%d%% Full", data.CurrentTruck.PercentFilled))

	data.CheckForFailure()
	if data.FirstLoad && !data.IsTimer && len(data.CurrentTruck.Wares) >= data.CurrentDifficulty.TargetWares {
		data.DepartureTimer = timing.New(float64(data.CurrentDifficulty.TimeToDepart))
		data.IsTimer = true
	}
	if data.IsTimer {
		UpdateTimer()
		if len(data.CurrentTruck.Wares) >= data.CurrentDifficulty.TargetWares {
			data.RightLoadedWares.SetColor(pixel.ToRGBA(colornames.Black))
			if data.FirstLoad {
				data.ButtonLock = false
			}
		} else {
			data.RightLoadedWares.SetColor(pixel.ToRGBA(colornames.Red))
			if data.FirstLoad {
				data.ButtonLock = true
			}
		}
	}
	if len(data.CurrentTruck.Wares) >= data.CurrentDifficulty.TargetWares && data.DepartureTimer.Done() {
		data.LeavePacking = true
	}
}
