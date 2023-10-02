package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/sfx"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/typeface"
)

// game over menu items
var (
	Statistics *typeface.Text
	ToMenu     *typeface.Text
)

func InitGameOverMenu() {
	if Statistics == nil {
		Statistics = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Top), 1.2, 0.18, 0, 0.)
		Statistics.Obj.Layer = 50
		Statistics.SetPos(pixel.V(0, 240.))
		Statistics.SetColor(constants.HoverUIText)
		Statistics.SetText("Booble: 45\njixrob\nIrjrjrjr")
		Statistics.Obj.SetRect(pixel.R(0, 0, 500, 420))
		Statistics.Obj.Rect = Statistics.Obj.Rect.Moved(pixel.V(0, -105))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, Statistics.Obj).
			AddComponent(myecs.Drawable, Statistics).
			AddComponent(myecs.DrawTarget, data.MenuView)
		MenuItems = append(MenuItems, Statistics)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(0, 300.))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText("Statistics")
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = Statistics.Obj.Hidden
				return false
			}))
	}
	if ToMenu == nil {
		ToMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		ToMenu.Obj.Layer = 50
		ToMenu.SetPos(pixel.V(0, -130.))
		ToMenu.SetColor(constants.BaseUIText)
		ToMenu.SetText("Main Menu")
		ToMenu.Obj.SetRect(pixel.R(0, 0, 400, 100))
		ToMenu.Obj.Rect = ToMenu.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, ToMenu.Obj).
			AddComponent(myecs.Drawable, ToMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					ToMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						click.Consume()
						data.LeavePacking = true
					}
				} else {
					ToMenu.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, ToMenu)
	}
}

func ShowGameOverMenu() {
	data.Starting = false
	HideAllMenus()
	Statistics.Obj.Hidden = false
	ToMenu.Obj.Hidden = false
}

func LeaveGameOverSystem() {
	if data.LeavePacking {
		if data.FadeTween == nil {
			data.FadeTween = gween.New(255., 0, 1, ease.Linear)
		}
		if data.LeaveTimer == nil {
			data.LeaveTimer = timing.New(1)
		}
		if data.LeaveTimer.UpdateDone() {
			state.PopState()
			data.LeaveStep = 5
		}
	}
}

func UpdateGameOverStats() {
	Statistics.SetText(fmt.Sprintf(`Truck: %s
Money Earned: $%d
Deliveries Complete: %d
Deliveries Missed: %d
Wares Abandoned: %d
Loaded Wares: %d
Stops Made: %d`,
		data.CurrentTruck.TruckLabel,
		data.CurrentScore.Cash,
		data.CurrentScore.DeliveryCount,
		data.CurrentScore.MissedDeliveries,
		data.CurrentScore.AbandonedWares,
		len(data.CurrentTruck.Wares),
		len(data.CartPositions)-1))
}
