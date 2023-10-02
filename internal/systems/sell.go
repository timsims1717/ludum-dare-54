package systems

import (
	"github.com/faiface/pixel"
	"image/color"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/util"
	"math"
)

func SellInit() {
	data.CurrentScore.DeliveryCount++
	for _, ware := range data.SellWares {
		myecs.Manager.DisposeEntity(ware.Entity)
	}
	data.SellWares = []*data.Ware{}
	numWares := util.Min(util.Min(1+data.CurrentScore.DeliveryCount/3, 3), len(data.CurrentTruck.Wares))
	data.BuyWares = numWares + constants.GlobalSeededRandom.Intn(2) + data.CurrentScore.DeliveryCount/4
	var used []int
	for i := 0; i < numWares; i++ {
		var sellWare *data.Ware
		if len(data.AbandonedWares) > 0 && constants.GlobalSeededRandom.Intn(10) == 0 {
			sellWare = data.GetFromAbandoned()
		} else {
			r := -1
			for r == -1 || util.Contains(r, used) {
				a := constants.GlobalSeededRandom.Intn(len(data.CurrentTruck.Wares))
				b := constants.GlobalSeededRandom.Intn(len(data.CurrentTruck.Wares))
				c := constants.GlobalSeededRandom.Intn(len(data.CurrentTruck.Wares))
				r = util.Min(a, util.Min(b, c))
			}
			used = append(used, r)
			sellWare = data.CurrentTruck.Wares[r].CopyWare()
		}
		sellWare.QueueIndex = i
		sellWare.SellMe = true
		obj := object.New().WithID("sell-ware")
		obj.Pos = pixel.V(slotX, rightQueueY(sellWare.QueueIndex))
		obj.Layer = 15
		spr := img.Batchers[constants.TestBatch].Sprites[sellWare.SpriteKey]
		obj.SetRect(spr.Frame())
		sca := slotSize * 0.9 / math.Max(obj.Rect.W(), obj.Rect.H())
		obj.Sca = pixel.V(sca, sca)
		sellWare.Sprite.Color = pixel.ToRGBA(color.RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 150,
		})
		sellWare.Object = obj
		sellWare.Entity = myecs.Manager.NewEntity()
		sellWare.Entity.AddComponent(myecs.Object, sellWare.Object).
			AddComponent(myecs.Drawable, sellWare.Sprite).
			AddComponent(myecs.Ware, sellWare).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.GameView, func(hvc *data.HoverClick) {
				if sellWare.Sold {

				} else {
					if !sellWare.Entity.HasComponent(myecs.Interpolation) {
						ip := object.NewInterpolation(object.InterpolateY).
							AddGween(sellWare.Object.Pos.Y, rightQueueY(sellWare.QueueIndex)+14, 0.8, ease.OutQuad).
							AddGween(rightQueueY(sellWare.QueueIndex)+14, rightQueueY(sellWare.QueueIndex), 0.8, ease.OutQuad)
						sellWare.Entity.AddComponent(myecs.Interpolation, ip)
					}
				}
			}))
		data.SellWares = append(data.SellWares, sellWare)
	}
}

func SellWare(ware *data.Ware) {
	ip := object.NewInterpolation(object.InterpolateX).
		AddGween(ware.Object.Pos.X, 500, 2., ease.InQuad)
	ware.Entity.AddComponent(myecs.Interpolation, ip)
	ware.Object.Pos.Y = rightQueueY(ware.QueueIndex)
	ware.Sprite.Color = util.White
	data.CurrentScore.Cash += ware.Value
	data.CurrentScore.SuccessfulDeliveries++
	ware.Sold = true
	PickupWare(ware.QueueIndex)
}

func PickupWare(i int) {
	if data.BuyWares > 0 {
		e := myecs.Manager.NewEntity()
		e.AddComponent(myecs.Update, data.NewTimerFunc(func() bool {
			CreateWareInQueue(i)
			myecs.Manager.DisposeEntity(e)
			return false
		}, 1))
		data.BuyWares--
	}
}
