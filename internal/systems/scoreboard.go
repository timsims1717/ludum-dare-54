package systems

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/typeface"
)

func ScoreboardInit() {
	data.LeftCount = typeface.New("main", typeface.NewAlign(typeface.Left, typeface.Top), 1.2, 0.15, 300., 0.)
	data.LeftCount.Obj.Layer = 30
	data.LeftCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.LeftCount.SetPos(pixel.V(-data.ScoreView.Rect.W()*0.5+8., data.ScoreView.Rect.H()-8.))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.LeftCount.Obj).
		AddComponent(myecs.Drawable, data.LeftCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.RightCount = typeface.New("main", typeface.NewAlign(typeface.Right, typeface.Top), 1.2, 0.15, 300., 0.)
	data.RightCount.Obj.Layer = 30
	data.RightCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.RightCount.SetPos(pixel.V(data.ScoreView.Rect.W()*0.5-8., data.ScoreView.Rect.H()-8.))
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.RightCount.Obj).
		AddComponent(myecs.Drawable, data.RightCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.PercCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.35, 300., 0.)
	data.PercCount.Obj.Layer = 30
	data.PercCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.PercCount.SetPos(pixel.V(0., 65.))
	data.PercCount.SetText("44% Full")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.PercCount.Obj).
		AddComponent(myecs.Drawable, data.PercCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)

	data.TimerCount = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.35, 300., 0.)
	data.TimerCount.Obj.Layer = 30
	data.TimerCount.SetColor(pixel.ToRGBA(colornames.Black))
	data.TimerCount.SetPos(pixel.V(0., 65.))
	data.TimerCount.SetText("16.012")
	myecs.Manager.NewEntity().
		AddComponent(myecs.Object, data.TimerCount.Obj).
		AddComponent(myecs.Drawable, data.TimerCount).
		AddComponent(myecs.DrawTarget, data.ScoreView)
}
