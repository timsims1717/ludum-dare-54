package systems

import (
	"github.com/bytearena/ecs"
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/util"
)

func NewBackground() {
	for _, e := range data.BackgroundItems {
		myecs.Manager.DisposeEntity(e)
	}
	data.BackgroundItems = []*ecs.Entity{}
	if data.CurrentTruck != nil {
		CreateRuts()
	}
	CreateDynamicBackground()
	CreateTables()
}

func CreateTables() {
	for i := 3; i >= 0; i-- {
		obj := object.New()
		obj.Pos = pixel.V(slotX, rightQueueY(i)-15.)
		obj.Layer = -2
		spr := img.NewSprite("ware_table", constants.TestBatch)
		e := myecs.Manager.NewEntity().
			AddComponent(myecs.Object, obj).
			AddComponent(myecs.Drawable, spr)
		data.BackgroundItems = append(data.BackgroundItems, e)
	}
}

func CreateRuts() {
	for i := 0; i < 9; i++ {
		objl := object.New()
		objl.Pos = pixel.V(0, float64(i-2)*32+16)
		if data.CurrentTruck.SpriteKey == "wagon" {
			objl.Pos.X -= 32
		}
		objl.Layer = -2
		objr := object.New()
		objr.Pos = pixel.V(float64((data.CurrentTruck.Width-1)*32), float64(i-2)*32+16)
		if data.CurrentTruck.SpriteKey == "wagon" {
			objr.Pos.X += 32
		}
		objr.Layer = -2
		str := "path_middle"
		if i == 0 {
			str = "path_bottom"
		}
		spr := img.NewSprite(str, constants.TestBatch)
		el := myecs.Manager.NewEntity().
			AddComponent(myecs.Object, objl).
			AddComponent(myecs.Drawable, spr)
		data.BackgroundItems = append(data.BackgroundItems, el)
		er := myecs.Manager.NewEntity().
			AddComponent(myecs.Object, objr).
			AddComponent(myecs.Drawable, spr)
		data.BackgroundItems = append(data.BackgroundItems, er)
	}
}

func CreateDynamicBackground() {
	var allPos []pixel.Vec
	c := constants.GlobalSeededRandom.Intn(20) + 20
outer:
	for i := 0; i < c; i++ {
		x := float64(constants.GlobalSeededRandom.Intn(650) - 300)
		y := float64(constants.GlobalSeededRandom.Intn(380) - 120)
		p := pixel.V(x, y)
		for _, pos := range allPos {
			if util.Magnitude(pos.Sub(p)) < 32 {
				i--
				continue outer
			}
		}
		allPos = append(allPos, p)
		obj := object.New()
		obj.Pos = p
		obj.Layer = -2
		str := "clover"
		switch constants.GlobalSeededRandom.Intn(3) {
		case 0:
			str = "bush"
		case 1:
			str = "grass"
		}
		spr := img.NewSprite(str, constants.TestBatch)
		e := myecs.Manager.NewEntity().
			AddComponent(myecs.Object, obj).
			AddComponent(myecs.Drawable, spr)
		data.BackgroundItems = append(data.BackgroundItems, e)
	}
}
