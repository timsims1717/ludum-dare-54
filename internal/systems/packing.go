package systems

import (
	"fmt"
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/typeface"
	"ludum-dare-54/pkg/util"
	"math"
)

var (
	bottomSlot = 60.
	slotSize   = 120.
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
				data.TruckTween = gween.New(0, 400, 4, ease.InQuad)
				data.LeaveTimer = timing.New(2)
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
			if data.TruckTween != nil {
				y, done := data.TruckTween.Update(timing.DT)
				TrunkOffset(y)
				if done {
					data.TruckTween = nil
				}
			}
			if data.LeaveTimer.UpdateDone() {
				data.ScoreTween = gween.New(data.ScoreView.PortPos.Y, 1000, 1, ease.InBack)
				data.SignTween = gween.New(data.SignObj.PostPos.Y, 260, 1, ease.InBack)
				data.LeaveTimer = timing.New(1)
				data.LeaveStep = 4
			}
		case 2:
			if data.LeaveTimer.UpdateDone() {
				data.LeaveStep = 3
				data.LeaveTimer = timing.New(1)
			}
		case 3:
			if data.LeaveTimer.UpdateDone() {
				state.PushState(constants.GameOverStateKey)
			}
		case 4:
			if data.TruckTween != nil {
				y, done := data.TruckTween.Update(timing.DT)
				TrunkOffset(y)
				if done {
					data.TruckTween = nil
				}
			}
			if data.SignTween != nil {
				ys, done := data.SignTween.Update(timing.DT)
				data.SignObj.Pos.Y = ys
				if done {
					data.SignTween = nil
				}
			}
			if data.ScoreTween != nil {
				y, done := data.ScoreTween.Update(timing.DT)
				data.ScoreView.PortPos.Y = y
				if done {
					data.ScoreTween = nil
				}
			}
			if data.FadeTween == nil {
				data.FadeTween = gween.New(255., 0, 1, ease.Linear)
			}
			if data.LeaveTimer.UpdateDone() {
				state.SwitchState(constants.TransitionStateKey)
			}
		case 5:
			state.SwitchState(constants.MainMenuStateKey)
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

func LoadSellLabels() {
	if data.WareNameLabelOne == nil {
		data.WareNameLabelOne = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.09, 0, 0.)
		data.WareNameLabelOne.Obj.Layer = 30
		data.WareNameLabelOne.SetColor(pixel.ToRGBA(colornames.Black))
		data.WareNameLabelOne.SetPos(pixel.V(slotX, rightQueueY(0)-40))
		data.WareNameLabelOne.SetText(fmt.Sprintf("Sell Name of Thing\n$56"))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.WareNameLabelOne.Obj).
			AddComponent(myecs.Drawable, data.WareNameLabelOne).
			AddComponent(myecs.DrawTarget, data.GameView)
	}
	if data.WareNameLabelTwo == nil {
		data.WareNameLabelTwo = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.09, 0, 0.)
		data.WareNameLabelTwo.Obj.Layer = 30
		data.WareNameLabelTwo.SetColor(pixel.ToRGBA(colornames.Black))
		data.WareNameLabelTwo.SetPos(pixel.V(slotX, rightQueueY(1)-40))
		data.WareNameLabelTwo.SetText(fmt.Sprintf("Sell Name of Thing\n$56"))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.WareNameLabelTwo.Obj).
			AddComponent(myecs.Drawable, data.WareNameLabelTwo).
			AddComponent(myecs.DrawTarget, data.GameView)
	}
	if data.WareNameLabelThree == nil {
		data.WareNameLabelThree = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.09, 0, 0.)
		data.WareNameLabelThree.Obj.Layer = 30
		data.WareNameLabelThree.SetColor(pixel.ToRGBA(colornames.Black))
		data.WareNameLabelThree.SetPos(pixel.V(slotX, rightQueueY(2)-40))
		data.WareNameLabelThree.SetText(fmt.Sprintf("Sell Name of Thing\n$56"))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, data.WareNameLabelThree.Obj).
			AddComponent(myecs.Drawable, data.WareNameLabelThree).
			AddComponent(myecs.DrawTarget, data.GameView)
	}
}

func UpdateSellLabels() {
	data.WareNameLabelOne.Obj.Hidden = true
	data.WareNameLabelTwo.Obj.Hidden = true
	data.WareNameLabelThree.Obj.Hidden = true
	if len(data.SellWares) > 0 {
		WareOne := data.SellWares[0]
		if WareOne != nil {
			if WareOne.SellMe && !WareOne.Sold {
				data.WareNameLabelOne.SetText(fmt.Sprintf("Sell %s\nfor $%d", WareOne.Name, WareOne.Value))
				data.WareNameLabelOne.Obj.Hidden = false
			}
		}
	}
	if data.WareQueue[0] != nil && data.WareNameLabelOne.Obj.Hidden {
		WareOne := data.WareQueue[0]
		data.WareNameLabelOne.SetText(fmt.Sprintf("Pick Up %s", WareOne.Name))
		data.WareNameLabelOne.Obj.Hidden = false
	}
	if len(data.SellWares) > 1 {
		WareTwo := data.SellWares[1]
		if WareTwo != nil {
			if WareTwo.SellMe && !WareTwo.Sold {
				data.WareNameLabelTwo.SetText(fmt.Sprintf("Sell %s\nfor $%d", WareTwo.Name, WareTwo.Value))
				data.WareNameLabelTwo.Obj.Hidden = false
			}
		}
	}
	if data.WareQueue[1] != nil && data.WareNameLabelTwo.Obj.Hidden {
		WareTwo := data.WareQueue[1]
		data.WareNameLabelTwo.SetText(fmt.Sprintf("Pick Up %s", WareTwo.Name))
		data.WareNameLabelTwo.Obj.Hidden = false
	}
	if len(data.SellWares) > 2 {
		WareThree := data.SellWares[2]
		if WareThree != nil {
			if WareThree.SellMe && !WareThree.Sold {
				data.WareNameLabelThree.SetText(fmt.Sprintf("Sell %s\nfor $%d", WareThree.Name, WareThree.Value))
				data.WareNameLabelThree.Obj.Hidden = false
			}
		}
	}
	if data.WareQueue[2] != nil && data.WareNameLabelThree.Obj.Hidden {
		WareThree := data.WareQueue[2]
		data.WareNameLabelThree.SetText(fmt.Sprintf("Pick Up %s", WareThree.Name))
		data.WareNameLabelThree.Obj.Hidden = false
	}
	data.WareNameLabelOne.Obj.SetRect(pixel.R(0, 0, data.WareNameLabelOne.Width+5, data.WareNameLabelOne.Height+2))
	data.WareNameLabelOne.Obj.Rect = data.WareNameLabelOne.Obj.Rect.Moved(pixel.V(0, -5))
	data.WareNameLabelTwo.Obj.SetRect(pixel.R(0, 0, data.WareNameLabelTwo.Width+5, data.WareNameLabelTwo.Height+2))
	data.WareNameLabelTwo.Obj.Rect = data.WareNameLabelTwo.Obj.Rect.Moved(pixel.V(0, -5))
	data.WareNameLabelThree.Obj.SetRect(pixel.R(0, 0, data.WareNameLabelThree.Width+5, data.WareNameLabelThree.Height+2))
	data.WareNameLabelThree.Obj.Rect = data.WareNameLabelThree.Obj.Rect.Moved(pixel.V(0, -5))
}

func TruckReturn() {
	data.CurrentTruck.TileObject.Offset.Y = 0
	for _, obj := range data.TrunkObjects {
		obj.Offset.Y = 0
	}
	for _, result := range myecs.Manager.Query(myecs.IsWare) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ware, okW := result.Components[myecs.Ware].(*data.Ware)
		if okO && okW {
			if ware.TIndex >= 0 {
				obj.Offset.Y = 0
			}
		}
	}
}

func TruckReset() {
	for _, e := range data.TrunkEntities {
		myecs.Manager.DisposeEntity(e)
	}
	data.TrunkEntities = []*ecs.Entity{}
	data.TrunkObjects = []*object.Object{}
}

func TrunkOffset(y float64) {
	data.CurrentTruck.TileObject.Offset.Y = y
	for _, obj := range data.TrunkObjects {
		obj.Offset.Y = y
	}
	for _, result := range myecs.Manager.Query(myecs.IsWare) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		ware, okW := result.Components[myecs.Ware].(*data.Ware)
		if okO && okW {
			if ware.TIndex >= 0 {
				obj.Offset.Y = y
			}
		}
	}
}

func DrawLabelBG(win *pixelgl.Window) {
	data.IMDraw.Clear()
	data.IMDraw.Color = constants.UIBGColor
	if !data.WareNameLabelOne.Obj.Hidden {
		data.IMDraw.Push(data.WareNameLabelOne.Obj.Pos.Add(data.WareNameLabelOne.Obj.Rect.Min))
		data.IMDraw.Push(data.WareNameLabelOne.Obj.Pos.Add(pixel.V(data.WareNameLabelOne.Obj.Rect.Min.X, data.WareNameLabelOne.Obj.Rect.Max.Y)))
		data.IMDraw.Push(data.WareNameLabelOne.Obj.Pos.Add(data.WareNameLabelOne.Obj.Rect.Max))
		data.IMDraw.Push(data.WareNameLabelOne.Obj.Pos.Add(pixel.V(data.WareNameLabelOne.Obj.Rect.Max.X, data.WareNameLabelOne.Obj.Rect.Min.Y)))
		data.IMDraw.Polygon(0)
	}
	if !data.WareNameLabelTwo.Obj.Hidden {
		data.IMDraw.Push(data.WareNameLabelTwo.Obj.Pos.Add(data.WareNameLabelTwo.Obj.Rect.Min))
		data.IMDraw.Push(data.WareNameLabelTwo.Obj.Pos.Add(pixel.V(data.WareNameLabelTwo.Obj.Rect.Min.X, data.WareNameLabelTwo.Obj.Rect.Max.Y)))
		data.IMDraw.Push(data.WareNameLabelTwo.Obj.Pos.Add(data.WareNameLabelTwo.Obj.Rect.Max))
		data.IMDraw.Push(data.WareNameLabelTwo.Obj.Pos.Add(pixel.V(data.WareNameLabelTwo.Obj.Rect.Max.X, data.WareNameLabelTwo.Obj.Rect.Min.Y)))
		data.IMDraw.Polygon(0)
	}
	if !data.WareNameLabelThree.Obj.Hidden {
		data.IMDraw.Push(data.WareNameLabelThree.Obj.Pos.Add(data.WareNameLabelThree.Obj.Rect.Min))
		data.IMDraw.Push(data.WareNameLabelThree.Obj.Pos.Add(pixel.V(data.WareNameLabelThree.Obj.Rect.Min.X, data.WareNameLabelThree.Obj.Rect.Max.Y)))
		data.IMDraw.Push(data.WareNameLabelThree.Obj.Pos.Add(data.WareNameLabelThree.Obj.Rect.Max))
		data.IMDraw.Push(data.WareNameLabelThree.Obj.Pos.Add(pixel.V(data.WareNameLabelThree.Obj.Rect.Max.X, data.WareNameLabelThree.Obj.Rect.Min.Y)))
		data.IMDraw.Polygon(0)
	}
	data.IMDraw.Draw(data.GameView.Canvas)
}
