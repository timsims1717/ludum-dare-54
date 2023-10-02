package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/sfx"
	"ludum-dare-54/pkg/typeface"
)

// Main Menu Items
var (
	MenuPlay     *typeface.Text
	MenuOptions  *typeface.Text
	MenuTutorial *typeface.Text
	MenuQuit     *typeface.Text
)

var MenuItems []*typeface.Text

func InitMenuItems(win *pixelgl.Window) {
	if MenuPlay == nil {
		MenuPlay = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
		MenuPlay.Obj.Layer = 50
		MenuPlay.SetPos(pixel.V(0, 100.))
		MenuPlay.SetColor(constants.BaseUIText)
		MenuPlay.SetText("Play")
		MenuPlay.Obj.SetRect(pixel.R(0, 0, 200, 100))
		MenuPlay.Obj.Rect = MenuPlay.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, MenuPlay.Obj).
			AddComponent(myecs.Drawable, MenuPlay).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					MenuPlay.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						ShowCarMenu()
						click.Consume()
					}
				} else {
					MenuPlay.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MenuPlay)
	}
	if MenuOptions == nil {
		MenuOptions = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
		MenuOptions.Obj.Layer = 50
		MenuOptions.SetPos(pixel.V(0, -15.))
		MenuOptions.SetColor(constants.BaseUIText)
		MenuOptions.SetText("Options")
		MenuOptions.Obj.SetRect(pixel.R(0, 0, 275, 100))
		MenuOptions.Obj.Rect = MenuOptions.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, MenuOptions.Obj).
			AddComponent(myecs.Drawable, MenuOptions).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					MenuOptions.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						ShowOptionsMenu()
						click.Consume()
					}
				} else {
					MenuOptions.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MenuOptions)
	}
	if MenuQuit == nil {
		MenuQuit = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
		MenuQuit.Obj.Layer = 50
		MenuQuit.SetPos(pixel.V(0, -130.))
		MenuQuit.SetColor(constants.BaseUIText)
		MenuQuit.SetText("Quit")
		MenuQuit.Obj.SetRect(pixel.R(0, 0, 200, 100))
		MenuQuit.Obj.Rect = MenuQuit.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, MenuQuit.Obj).
			AddComponent(myecs.Drawable, MenuQuit).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					MenuQuit.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						click.Consume()
						win.SetClosed(true)
					}
				} else {
					MenuQuit.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MenuQuit)
	}

	InitOptionsMenu()
	InitCarMenu()
	InitDifficultyMenu()
	InitPauseMenu()
	InitGameOverMenu()
}

func ShowMainMenu() {
	data.Starting = false
	HideAllMenus()
	MenuPlay.Obj.Hidden = false
	MenuOptions.Obj.Hidden = false
	MenuQuit.Obj.Pos.Y = -130
	MenuQuit.Obj.Hidden = false
}

func HideAllMenus() {
	for _, item := range MenuItems {
		if item != nil {
			item.Obj.Hidden = true
		}
	}
}

func DrawMenuBG() {
	for _, item := range MenuItems {
		if item != nil && !item.Obj.Hidden {
			data.IMDraw.Clear()
			data.IMDraw.Color = constants.UIBGColor
			data.IMDraw.Push(item.Obj.Pos.Add(item.Obj.Rect.Min))
			data.IMDraw.Push(item.Obj.Pos.Add(pixel.V(item.Obj.Rect.Min.X, item.Obj.Rect.Max.Y)))
			data.IMDraw.Push(item.Obj.Pos.Add(item.Obj.Rect.Max))
			data.IMDraw.Push(item.Obj.Pos.Add(pixel.V(item.Obj.Rect.Max.X, item.Obj.Rect.Min.Y)))
			data.IMDraw.Polygon(0)
			data.IMDraw.Draw(data.MenuView.Canvas)
		}
	}
}
