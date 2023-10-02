package systems

import (
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

// Pause Menu Items
var (
	MenuResume  *typeface.Text
	MenuAbandon *typeface.Text
)

func InitPauseMenu() {
	if MenuResume == nil {
		MenuResume = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		MenuResume.Obj.Layer = 50
		MenuResume.SetPos(pixel.V(0, 100.))
		MenuResume.SetColor(constants.BaseUIText)
		MenuResume.SetText("Resume")
		MenuResume.Obj.SetRect(pixel.R(0, 0, 290, 100))
		MenuResume.Obj.Rect = MenuResume.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, MenuResume.Obj).
			AddComponent(myecs.Drawable, MenuResume).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					MenuResume.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						click.Consume()
						state.PopState()
					}
				} else {
					MenuResume.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MenuResume)
	}
	if MenuAbandon == nil {
		MenuAbandon = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		MenuAbandon.Obj.Layer = 50
		MenuAbandon.SetPos(pixel.V(0, -130.))
		MenuAbandon.SetColor(constants.BaseUIText)
		MenuAbandon.SetText("Abandon Game")
		MenuAbandon.Obj.SetRect(pixel.R(0, 0, 600, 100))
		MenuAbandon.Obj.Rect = MenuAbandon.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, MenuAbandon.Obj).
			AddComponent(myecs.Drawable, MenuAbandon).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					MenuAbandon.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						click.Consume()
						Abandon()
					}
				} else {
					MenuAbandon.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MenuAbandon)
	}
}

func ShowPauseMenu() {
	HideAllMenus()
	MenuResume.Obj.Hidden = false
	MenuOptions.Obj.Hidden = false
	MenuAbandon.Obj.Hidden = false
	MenuQuit.Obj.Pos.Y = -245
	MenuQuit.Obj.Hidden = false
}

func AbandonSystem() {
	if data.LeaveTransition {
		if data.FadeTween == nil {
			data.FadeTween = gween.New(255., 0, 1, ease.Linear)
		}
		if data.TransitionTimer == nil {
			data.TransitionTimer = timing.New(1)
		}
		if data.TransitionTimer.UpdateDone() {
			state.SwitchState(constants.MainMenuStateKey)
		}
	}
}

func Abandon() {
	data.LeaveTransition = true
}
