package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/util"
	"math"
)

var (
	bottomSlot = -40.
	slotSize   = 64.
	slotX      = 240.
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
		data.FirstLoad = false
		if data.ScoreTween == nil {
			data.ScoreTween = gween.New(data.ScoreView.PortPos.Y, 1000, 1, ease.InBack)
		}
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
	}
}

func StartTimer() {
	data.DepartureTimer = timing.New(float64(data.CurrentDifficulty.TimeToDepart))
	data.IsTimer = true
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
