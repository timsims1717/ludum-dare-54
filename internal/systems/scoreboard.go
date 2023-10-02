package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	if data.Sign == nil {
		data.SignObj = object.New()
		data.SignObj.Layer = 2
		data.SignObj.Pos = pixel.V(-170, 135)
		spr := img.NewSprite("sign", constants.TestBatch)
		data.Sign = myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.SignObj).
			AddComponent(myecs.Drawable, spr)
	} else {
		data.SignObj.Pos = pixel.V(-170, 135)
	}
	data.SignTween = nil

	if data.LeftTitle == nil {
		data.LeftTitle = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 0, 0.)
		data.LeftTitle.Obj.Layer = 30
		data.LeftTitle.SetColor(pixel.ToRGBA(colornames.Black))
		data.LeftTitle.SetPos(pixel.V(leftScoreboard, topScoreboard))
		data.LeftTitle.SetText("DELIVERIES")
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.LeftTitle.Obj).
			AddComponent(myecs.Drawable, data.LeftTitle).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.LeftComplete == nil {
		data.LeftComplete = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 0, 0.)
		data.LeftComplete.Obj.Layer = 30
		data.LeftComplete.SetColor(pixel.ToRGBA(colornames.Black))
		data.LeftComplete.SetPos(pixel.V(leftScoreboard, topScoreboard-lineHeight))
		data.LeftComplete.SetText(fmt.Sprintf("%d Complete", 42))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.LeftComplete.Obj).
			AddComponent(myecs.Drawable, data.LeftComplete).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.LeftMissed == nil {
		data.LeftMissed = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 0, 0.)
		data.LeftMissed.Obj.Layer = 30
		data.LeftMissed.SetColor(pixel.ToRGBA(colornames.Black))
		data.LeftMissed.SetPos(pixel.V(leftScoreboard, topScoreboard-lineHeight*2))
		data.LeftMissed.SetText(fmt.Sprintf("%d Missed", 42))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.LeftMissed.Obj).
			AddComponent(myecs.Drawable, data.LeftMissed).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.LeftAbandoned == nil {
		data.LeftAbandoned = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 0, 0.)
		data.LeftAbandoned.Obj.Layer = 30
		data.LeftAbandoned.SetColor(pixel.ToRGBA(colornames.Black))
		data.LeftAbandoned.SetPos(pixel.V(leftScoreboard, topScoreboard-lineHeight*3))
		data.LeftAbandoned.SetText(fmt.Sprintf("%d Abandonded", 42))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.LeftAbandoned.Obj).
			AddComponent(myecs.Drawable, data.LeftAbandoned).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.LeftCash == nil {
		data.LeftCash = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 0, 0.)
		data.LeftCash.Obj.Layer = 30
		data.LeftCash.SetColor(pixel.ToRGBA(colornames.Black))
		data.LeftCash.SetPos(pixel.V(leftScoreboard, topScoreboard-lineHeight*4))
		data.LeftCash.SetText(fmt.Sprintf("$%d", 42))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.LeftCash.Obj).
			AddComponent(myecs.Drawable, data.LeftCash).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.RightTitle == nil {
		data.RightTitle = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 0, 0.)
		data.RightTitle.Obj.Layer = 30
		data.RightTitle.SetColor(pixel.ToRGBA(colornames.Black))
		data.RightTitle.SetPos(pixel.V(rightScoreboard, topScoreboard))
		data.RightTitle.SetText(data.CurrentTruck.TruckLabel)
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.RightTitle.Obj).
			AddComponent(myecs.Drawable, data.RightTitle).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.RightLoadedWares == nil {
		data.RightLoadedWares = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 0, 0.)
		data.RightLoadedWares.Obj.Layer = 30
		data.RightLoadedWares.SetColor(pixel.ToRGBA(colornames.Black))
		data.RightLoadedWares.SetPos(pixel.V(rightScoreboard, topScoreboard-lineHeight))
		data.RightLoadedWares.SetText(fmt.Sprintf("Loaded Wares: 42"))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.RightLoadedWares.Obj).
			AddComponent(myecs.Drawable, data.RightLoadedWares).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.RightLoadHeight == nil {
		data.RightLoadHeight = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 0, 0.)
		data.RightLoadHeight.Obj.Layer = 30
		data.RightLoadHeight.SetColor(pixel.ToRGBA(colornames.Black))
		data.RightLoadHeight.SetPos(pixel.V(rightScoreboard, topScoreboard-lineHeight*2))
		data.RightLoadHeight.SetText(fmt.Sprintf("Load Height: %d / %d", 0, data.CurrentTruck.Depth))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.RightLoadHeight.Obj).
			AddComponent(myecs.Drawable, data.RightLoadHeight).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.RightPercentFull == nil {
		data.RightPercentFull = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 0, 0.)
		data.RightPercentFull.Obj.Layer = 30
		data.RightPercentFull.SetColor(pixel.ToRGBA(colornames.Black))
		data.RightPercentFull.SetPos(pixel.V(rightScoreboard, topScoreboard-lineHeight*3))
		data.RightPercentFull.SetText(fmt.Sprintf("%d%% Full", 42))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.RightPercentFull.Obj).
			AddComponent(myecs.Drawable, data.RightPercentFull).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.RightWaresCount == nil {
		data.RightWaresCount = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 0, 0.)
		data.RightWaresCount.Obj.Layer = 30
		data.RightWaresCount.SetColor(pixel.ToRGBA(colornames.Black))
		data.RightWaresCount.SetPos(pixel.V(rightScoreboard, topScoreboard-lineHeight*4))
		data.RightWaresCount.SetText(fmt.Sprintf("%d/%d Min Wares", 42, 42))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.RightWaresCount.Obj).
			AddComponent(myecs.Drawable, data.RightWaresCount).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.MinWaresCount == nil {
		data.MinWaresCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.2, 0, 0.)
		data.MinWaresCount.Obj.Layer = 30
		data.MinWaresCount.SetColor(pixel.ToRGBA(colornames.Black))
		data.MinWaresCount.SetPos(pixel.V(centerScoreboard, middleScoreboard))
		data.MinWaresCount.SetText("Min Wares xx / yy")
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.MinWaresCount.Obj).
			AddComponent(myecs.Drawable, data.MinWaresCount).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.TimerCount == nil {
		data.TimerCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		data.TimerCount.Obj.Layer = 30
		data.TimerCount.SetColor(pixel.ToRGBA(colornames.Black))
		data.TimerCount.SetPos(pixel.V(centerScoreboard, middleScoreboard))
		data.TimerCount.SetText("16.012")
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.TimerCount.Obj).
			AddComponent(myecs.Drawable, data.TimerCount).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	if data.ButtonText == nil {
		data.ButtonText = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.24, 0, 0.)
		data.ButtonText.Obj.Layer = 30
		data.ButtonText.SetColor(pixel.ToRGBA(colornames.Black))
		data.ButtonText.SetPos(pixel.V(centerScoreboard, buttonYHeight))
		data.ButtonText.SetText("Set Out")
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.ButtonText.Obj).
			AddComponent(myecs.Drawable, data.ButtonText).
			AddComponent(myecs.DrawTarget, data.ScoreView)
	}

	data.ButtonLock = true
	clickedOn := false
	if data.ButtonObj == nil {
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
}

func ScoreboardReset() {
	data.TimerCount.SetColor(constants.HoverUIText)
	data.LeftTitle.SetColor(constants.HoverUIText)
	data.LeftComplete.SetColor(constants.HoverUIText)
	data.LeftMissed.SetColor(constants.HoverUIText)
	data.LeftAbandoned.SetColor(constants.HoverUIText)
	data.LeftCash.SetColor(constants.HoverUIText)
	data.RightTitle.SetColor(constants.HoverUIText)
	data.RightLoadedWares.SetColor(constants.HoverUIText)
	data.RightLoadHeight.SetColor(constants.HoverUIText)
	data.RightPercentFull.SetColor(constants.HoverUIText)
	data.MinWaresCount.SetColor(constants.HoverUIText)
	data.TimerCount.SetColor(constants.HoverUIText)
	data.ButtonText.SetColor(constants.HoverUIText)
	data.SignObj.Pos = pixel.V(-170, 135)
	data.SignTween = nil
}

func ScoreSystem() {
	data.RightLoadedWares.SetText(fmt.Sprintf("Loaded Wares: %d", len(data.CurrentTruck.Wares)))
	data.RightLoadHeight.SetText(fmt.Sprintf("Load Height: %d / %d", data.CurrentTruck.CurrHeight, data.CurrentTruck.Height))
	//data.LeftTitle.SetText(fmt.Sprintf("DELIVERIES\n%d Complete\n%d Missed\n$%d.00", data.CurrentScore.SuccessfulDeliveries,
	//data.CurrentScore.MissedDeliveries, data.CurrentScore.Cash))
	data.LeftComplete.SetText(fmt.Sprintf("%d Complete", data.CurrentScore.SuccessfulDeliveries))
	data.LeftMissed.SetText(fmt.Sprintf("%d Missed", data.CurrentScore.MissedDeliveries))
	data.LeftAbandoned.SetText(fmt.Sprintf("%d Abandoned", data.CurrentScore.AbandonedWares))
	data.LeftCash.SetText(fmt.Sprintf("$%d", data.CurrentScore.Cash))
	data.MinWaresCount.SetText(fmt.Sprintf("Min Wares: %d/%d", len(data.CurrentTruck.Wares), data.CurrentDifficulty.TargetWares))
	data.RightPercentFull.SetText(fmt.Sprintf("%d%% Full", data.CurrentTruck.PercentFilled))
	data.RightWaresCount.SetText(fmt.Sprintf("Min Wares: %d/%d", len(data.CurrentTruck.Wares), data.CurrentDifficulty.TargetWares))

	if data.FirstLoad && !data.IsTimer && len(data.CurrentTruck.Wares) >= data.CurrentDifficulty.TargetWares {
		data.DepartureTimer = timing.New(float64(data.CurrentDifficulty.TimeToDepart))
		data.IsTimer = true
	}
	data.CheckForFailure()
	data.ButtonLock = data.CurrentScore.FailCondition != constants.NotFailing
	if data.IsTimer {
		UpdateTimer()
		if len(data.CurrentTruck.Wares) >= data.CurrentDifficulty.TargetWares {
			data.RightLoadedWares.SetColor(constants.HoverUIText)
		} else {
			data.RightLoadedWares.SetColor(constants.BadUIText)
		}
	}
	if data.FirstLoad {
		if len(data.CurrentTruck.Wares) >= data.CurrentDifficulty.TargetWares && data.DepartureTimer.Done() {
			data.LeavePacking = true
		}
	} else if data.DepartureTimer.Done() {
		data.LeavePacking = true
	}
	if data.Abandon {
		data.LeavePacking = true
	}
}

func BigMessageInit() {
	if data.BigMessage == nil {
		data.BigMessage = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.22, 600., 0.)
		data.BigMessage.Obj.Layer = 50
		data.BigMessage.SetColor(pixel.ToRGBA(colornames.Black))
		data.BigMessage.SetText("Test")
		data.BigMessage.Obj.Hidden = true
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.BigMessage.Obj).
			AddComponent(myecs.Drawable, data.BigMessage).
			AddComponent(myecs.DrawTarget, data.GameView)
	}
	data.BigMessage.SetPos(data.GameView.CamPos)
}

func SetBigMessage(raw string, col pixel.RGBA, dur float64) {
	BigMessageInit()
	data.BigMessage.SetColor(col)
	data.BigMessage.SetText(raw)
	data.BigMessage.Obj.SetRect(pixel.R(0, 0, 700, data.BigMessage.Height*1.2))
	data.BigMessage.Obj.Rect = data.BigMessage.Obj.Rect.Moved(pixel.V(0, data.BigMessage.Height*-0.2))
	data.BigMessage.Obj.Hidden = false
	if dur > 0 {
		data.BigMessageTimer = timing.New(dur)
	} else {
		data.BigMessageTimer = nil
	}
}

func HideBigMessage() {
	data.BigMessage.Obj.Hidden = true
}

func BigMessageSystem() {
	if data.BigMessageTimer != nil {
		if data.BigMessageTimer.UpdateDone() {
			HideBigMessage()
		}
	}
}

func DrawBigMessage(win *pixelgl.Window) {
	if !data.BigMessage.Obj.Hidden {
		data.IMDraw.Clear()
		data.IMDraw.Color = constants.UIBGColor
		data.IMDraw.Push(data.BigMessage.Obj.Pos.Add(data.BigMessage.Obj.Rect.Min))
		data.IMDraw.Push(data.BigMessage.Obj.Pos.Add(pixel.V(data.BigMessage.Obj.Rect.Min.X, data.BigMessage.Obj.Rect.Max.Y)))
		data.IMDraw.Push(data.BigMessage.Obj.Pos.Add(data.BigMessage.Obj.Rect.Max))
		data.IMDraw.Push(data.BigMessage.Obj.Pos.Add(pixel.V(data.BigMessage.Obj.Rect.Max.X, data.BigMessage.Obj.Rect.Min.Y)))
		data.IMDraw.Polygon(0)
		data.IMDraw.Draw(data.GameView.Canvas)
	}
	DrawSystem(win, 50)
}
