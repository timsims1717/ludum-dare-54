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
	"ludum-dare-54/pkg/typeface"
)

func ScoreboardInit() {
	data.LeftTitle = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftTitle.Obj.Layer = 30
	data.LeftTitle.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftTitle.SetPos(pixel.V(-140, 170))
	data.LeftTitle.SetText("DELIVERIES")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftTitle.Obj).
		AddComponent(myecs.Drawable, data.LeftTitle).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.LeftCompletes = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftCompletes.Obj.Layer = 30
	data.LeftCompletes.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftCompletes.SetPos(pixel.V(-140, 140))
	data.LeftCompletes.SetText(fmt.Sprintf("%d Complete", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftCompletes.Obj).
		AddComponent(myecs.Drawable, data.LeftCompletes).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.LeftMisseds = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftMisseds.Obj.Layer = 30
	data.LeftMisseds.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftMisseds.SetPos(pixel.V(-140, 110))
	data.LeftMisseds.SetText(fmt.Sprintf("%d Missed", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftMisseds.Obj).
		AddComponent(myecs.Drawable, data.LeftMisseds).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.LeftAbandoned = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftAbandoned.Obj.Layer = 30
	data.LeftAbandoned.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftAbandoned.SetPos(pixel.V(-140, 80))
	data.LeftAbandoned.SetText(fmt.Sprintf("%d Abandonded", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftAbandoned.Obj).
		AddComponent(myecs.Drawable, data.LeftAbandoned).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.LeftCash = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftCash.Obj.Layer = 30
	data.LeftCash.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftCash.SetPos(pixel.V(-140, 50))
	data.LeftCash.SetText(fmt.Sprintf("$%d", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftCash.Obj).
		AddComponent(myecs.Drawable, data.LeftCash).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightTitle = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightTitle.Obj.Layer = 30
	data.RightTitle.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightTitle.SetPos(pixel.V(140, 170))
	data.RightTitle.SetText(data.CurrentTruck.TruckLabel)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightTitle.Obj).
		AddComponent(myecs.Drawable, data.RightTitle).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightLoadedWares = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightLoadedWares.Obj.Layer = 30
	data.RightLoadedWares.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightLoadedWares.SetPos(pixel.V(140, 140))
	data.RightLoadedWares.SetText(fmt.Sprintf("Loaded Wares: 42"))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightLoadedWares.Obj).
		AddComponent(myecs.Drawable, data.RightLoadedWares).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightLoadHeight = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightLoadHeight.Obj.Layer = 30
	data.RightLoadHeight.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightLoadHeight.SetPos(pixel.V(140, 110))
	data.RightLoadHeight.SetText(fmt.Sprintf("Load Height: %d / %d", 0, data.CurrentTruck.Depth))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightLoadHeight.Obj).
		AddComponent(myecs.Drawable, data.RightLoadHeight).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightPercentFull = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightPercentFull.Obj.Layer = 30
	data.RightPercentFull.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightPercentFull.SetPos(pixel.V(140, 80))
	data.RightPercentFull.SetText(fmt.Sprintf("%d%% Full", 42))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightPercentFull.Obj).
		AddComponent(myecs.Drawable, data.RightPercentFull).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.PercCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
	data.PercCount.Obj.Layer = 30
	data.PercCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.PercCount.SetPos(pixel.V(0., -20))
	data.PercCount.SetText("44% Full")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.PercCount.Obj).
		AddComponent(myecs.Drawable, data.PercCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.TimerCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
	data.TimerCount.Obj.Layer = 30
	data.TimerCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.TimerCount.SetPos(pixel.V(0., -20))
	data.TimerCount.SetText("16.012")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.TimerCount.Obj).
		AddComponent(myecs.Drawable, data.TimerCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.ButtonText = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.24, 300., 0.)
	data.ButtonText.Obj.Layer = 30
	data.ButtonText.SetColor(pixel.ToRGBA(colornames.Black))
	data.ButtonText.SetPos(pixel.V(0., -100.))
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
	data.ButtonObj.Pos = pixel.V(0., -115.)
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.ButtonObj).
		AddComponent(myecs.Drawable, data.ButtonSpr).
		AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.ScoreView, func(hvc *data.HoverClick) {
			if !data.ButtonLock {
				click := hvc.Input.Get("click")
				data.ButtonSpr.Color = pixel.RGB(1, 1, 1)
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
