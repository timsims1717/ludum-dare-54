package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

var MenuItems = []*data.MenuItem{
	MenuPlay,
	MenuTutorial,
	MenuOptions,
	MenuQuit,
}

func InitMenuItems() {
	playText := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
	playText.Obj.Layer = 50
	playText.SetPos(pixel.V(0, 10.))
	playText.SetColor(pixel.ToRGBA(colornames.Black))
	playText.SetText("Play")
	playText.Obj.SetRect(pixel.R(0, 0, 50, 20))
	MenuPlay = &data.MenuItem{
		Text: playText,
		Func: nil,
	}
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, playText.Obj).
		AddComponent(myecs.Drawable, playText).
		AddComponent(myecs.DrawTarget, data.MenuView)
}

func DrawMenuBG(win *pixelgl.Window) {
	for _, item := range MenuItems {
		if item != nil && !item.Text.Obj.Hidden {
			data.MenuIMDraw.Clear()
			data.MenuIMDraw.Color = constants.UIBGColor
			data.MenuIMDraw.Push(item.Text.Obj.Pos.Add(item.Text.Obj.Rect.Min))
			data.MenuIMDraw.Push(item.Text.Obj.Pos.Add(pixel.V(item.Text.Obj.Rect.Min.X, item.Text.Obj.Rect.Max.Y)))
			data.MenuIMDraw.Push(item.Text.Obj.Pos.Add(item.Text.Obj.Rect.Max))
			data.MenuIMDraw.Push(item.Text.Obj.Pos.Add(pixel.V(item.Text.Obj.Rect.Max.X, item.Text.Obj.Rect.Min.Y)))
			data.MenuIMDraw.Draw(data.MenuView.Canvas)
		}
	}
	data.MenuView.Draw(win)
}
