package systems

import (
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
	data.LeftCount = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftCount.Obj.Layer = 30
	data.LeftCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftCount.SetPos(pixel.V(-140, 170))
	data.LeftCount.SetText("Wares: 18\n$50.00")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftCount.Obj).
		AddComponent(myecs.Drawable, data.LeftCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightCount = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightCount.Obj.Layer = 30
	data.RightCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightCount.SetPos(pixel.V(140, 170))
	data.RightCount.SetText("Deliveries\n0 Complete\n0 Missed")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightCount.Obj).
		AddComponent(myecs.Drawable, data.RightCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.PercCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
	data.PercCount.Obj.Layer = 30
	data.PercCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.PercCount.SetPos(pixel.V(0., 25.))
	data.PercCount.SetText("44% Full")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.PercCount.Obj).
		AddComponent(myecs.Drawable, data.PercCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.TimerCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 300., 0.)
	data.TimerCount.Obj.Layer = 30
	data.TimerCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.TimerCount.SetPos(pixel.V(0., 25.))
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
