package systems

import (
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/typeface"
)

// Main Menu Items
var (
	MenuPlay     *data.MenuItem
	MenuTutorial *data.MenuItem
	MenuOptions  *data.MenuItem
	MenuQuit     *data.MenuItem
)

var MenuItems []*data.MenuItem

func InitMenuItems() {
	if MenuPlay == nil {
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(0, 100.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText("Play")
		txt.Obj.SetRect(pixel.R(0, 0, 200, 100))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -26))
		MenuPlay = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowCarMenu()
						click.Consume()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MenuPlay)
	}

	InitCarMenu()
	InitDifficultyMenu()
	ShowMainMenu()
}

func ShowMainMenu() {
	data.Starting = false
	HideAllMenus()
	MenuPlay.Text.Obj.Hidden = false
}

func HideAllMenus() {
	for _, item := range MenuItems {
		if item != nil {
			item.Text.Obj.Hidden = true
		}
	}
}

func DrawMenuBG() {
	for _, item := range MenuItems {
		if item != nil && !item.Text.Obj.Hidden {
			data.IMDraw.Clear()
			data.IMDraw.Color = constants.UIBGColor
			data.IMDraw.Push(item.Text.Obj.Pos.Add(item.Text.Obj.Rect.Min))
			data.IMDraw.Push(item.Text.Obj.Pos.Add(pixel.V(item.Text.Obj.Rect.Min.X, item.Text.Obj.Rect.Max.Y)))
			data.IMDraw.Push(item.Text.Obj.Pos.Add(item.Text.Obj.Rect.Max))
			data.IMDraw.Push(item.Text.Obj.Pos.Add(pixel.V(item.Text.Obj.Rect.Max.X, item.Text.Obj.Rect.Min.Y)))
			data.IMDraw.Polygon(0)
			data.IMDraw.Draw(data.MenuView.Canvas)
		}
	}
}
