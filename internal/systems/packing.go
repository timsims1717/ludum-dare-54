package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/util"
	"math"
)

var (
	bottomSlot = -40.
	slotSize   = 64.
	slotX      = 230.
)

func rightQueueY(i int) float64 {
	return -bottomSlot + (float64(i) * slotSize)
}

func GetNearestPos(pos pixel.Vec, r pixel.Rect) pixel.Vec {
	nPos := pos
	if !data.BottomDrop.Contains(pos) && !data.LeftDrop.Contains(pos) {
		bdx := 0.
		if pos.X > data.BottomDrop.Max.X {
			bdx = pos.X - data.BottomDrop.Max.X
		} else if pos.X < data.BottomDrop.Min.X {
			bdx = pos.X - data.BottomDrop.Min.X
		}
		bdy := 0.
		if pos.Y > data.BottomDrop.Max.Y {
			bdy = pos.Y - data.BottomDrop.Max.Y
		} else if pos.Y < data.BottomDrop.Min.Y {
			bdy = pos.Y - data.BottomDrop.Min.Y
		}
		ldx := 0.
		if pos.X > data.LeftDrop.Max.X {
			ldx = pos.X - data.LeftDrop.Max.X
		} else if pos.X < data.LeftDrop.Min.X {
			ldx = pos.X - data.LeftDrop.Min.X
		}
		ldy := 0.
		if pos.Y > data.LeftDrop.Max.Y {
			ldy = pos.Y - data.LeftDrop.Max.Y
		} else if pos.Y < data.LeftDrop.Min.Y {
			ldy = pos.Y - data.LeftDrop.Min.Y
		}
		bd := math.Sqrt(bdx*bdx + bdy*bdy)
		ld := math.Sqrt(ldx*ldx + ldy*ldy)
		if bd > ld {
			if ldx > 0 {
				nPos.X = data.LeftDrop.Max.X
			} else if ldx < 0 {
				nPos.X = data.LeftDrop.Min.X
			}
			if ldy > 0 {
				nPos.Y = data.LeftDrop.Max.Y
			} else if ldy < 0 {
				nPos.Y = data.LeftDrop.Min.Y
			}
		} else {
			if bdx > 0 {
				nPos.X = data.BottomDrop.Max.X
			} else if bdx < 0 {
				nPos.X = data.BottomDrop.Min.X
			}
			if bdy > 0 {
				nPos.Y = data.BottomDrop.Max.Y
			} else if bdy < 0 {
				nPos.Y = data.BottomDrop.Min.Y
			}
		}
	}
	if data.BottomDrop.Contains(nPos) {
		return util.ConstrainR(nPos, data.BottomDrop.Center(), r, data.BottomDrop)
	}
	if data.LeftDrop.Contains(nPos) {
		return util.ConstrainR(nPos, data.LeftDrop.Center(), r, data.LeftDrop)
	}
	return pixel.V(-100, -100)
}

func LeavePackingSystem() {
	if data.LeavePacking {
		switch data.LeaveStep {
		case 0:
			if !data.FirstLoad {
				for _, ware := range data.SellWares {
					if !ware.Sold {
						data.CurrentScore.MissedDeliveries++
					}
				}
				for _, result := range myecs.Manager.Query(myecs.IsWare) {
					_, okO := result.Components[myecs.Object].(*object.Object)
					ware, okW := result.Components[myecs.Ware].(*data.Ware)
					if okO && okW {
						if ware.TIndex < 0 && ware.QueueIndex < 0 {
							data.CurrentScore.AbandonedWares++
							data.AddToAbandoned(ware)
						}
					}
				}
			}
			data.FirstLoad = false
			data.CheckForFailure()
			if data.CurrentScore.FailCondition == constants.NotFailing {
				data.ScoreTween = gween.New(data.ScoreView.PortPos.Y, 1000, 1, ease.InBack)
				data.LeaveStep = 1
			} else {
				failMsg := "Game Over"
				switch data.CurrentScore.FailCondition {
				case constants.TooManyMisses:
					failMsg = "You missed too many deliveries."
				case constants.AbandonToManyItems:
					failMsg = "You left behind too many wares."
				case constants.TooFewItems:
					failMsg = "You don't have enough wares in your truck."
				case constants.Abandoned:
					failMsg = "You left your truck."
				}
				SetBigMessage(failMsg, constants.BadUIText, 8)
				data.LeaveTimer = timing.New(5)
				data.LeaveStep = 2
			}
		case 1:
			y, done := data.ScoreTween.Update(timing.DT)
			data.ScoreView.PortPos.Y = y
			if done {
				data.ScoreTween = nil
			}
			if data.FadeTween == nil {
				data.FadeTween = gween.New(255., 0, 1, ease.Linear)
			}
			if done {
				state.SwitchState(constants.TransitionStateKey)
			}
		case 2:
			if data.LeaveTimer.UpdateDone() {
				data.LeaveStep = 3
				data.FadeTween = gween.New(255., 0, 1, ease.Linear)
				data.LeaveTimer = timing.New(1)
			}
		case 3:
			if data.LeaveTimer.UpdateDone() {
				state.SwitchState(constants.MainMenuStateKey)
			}
		}
	}
}

func UpdateTimer() {
	data.DepartureTimer.Update()
	if data.DepartureTimer.Sec()-data.DepartureTimer.Elapsed() <= 0 {
		data.TimerCount.SetText("00.0")
		data.TimerCount.SetColor(pixel.ToRGBA(colornames.Red))
	} else {
		data.TimerCount.SetText(fmt.Sprintf("%.1f", data.DepartureTimer.Sec()-data.DepartureTimer.Elapsed()))
	}
}
