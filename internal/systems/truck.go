package systems

import (
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
)

func CreateTruck(w, d, h float64) {
	if data.Truck == nil {
		data.NewTruck()
	}
	for y := 0; y < int(d); y++ {
		for x := 0; x < int(w); x++ {
			obj := object.New()
			obj.Pos = pixel.V(float64(x)*64, float64(y)*64)
			obj.Layer = 1
			e := myecs.Manager.NewEntity()
			e.AddComponent(myecs.Drawable, img.NewSprite("square", data.TestBatch)).
				AddComponent(myecs.Object, obj)
		}
	}
}
