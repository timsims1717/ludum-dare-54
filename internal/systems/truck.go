package systems

import (
	"github.com/faiface/pixel"
	"image/color"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/debug"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/world"
)

var claimed = false

func CreateTruck(w, d, h float64) {
	if data.Truck == nil {
		data.NewTruck(int(w), int(d), int(h))
	}
	for yt := 0; yt < data.Truck.Depth; yt++ {
		for xt := 0; xt < data.Truck.Width; xt++ {
			obj := object.New().WithID("slot")
			obj.Pos = pixel.V(float64(xt)*world.TileSize, float64(yt)*world.TileSize)
			obj.SetRect(pixel.R(0, 0, world.TileSize, world.TileSize))
			obj.Layer = 0
			e := myecs.Manager.NewEntity()
			e.AddComponent(myecs.Drawable, img.NewSprite("square", constants.TestBatch)).
				AddComponent(myecs.Object, obj).
				AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.GameView, func(hvc *data.HoverClick) {
					if hvc.Hover && !claimed {
						claimed = true
						x, y := world.WorldToMapAdj(hvc.Pos.X, hvc.Pos.Y)
						debug.AddIntCoords("Hovered Over", x, y)
						if data.HeldItem != nil {
							shadowObj := object.New()
							// change the pos of items w/even widths and depths
							shadowObj.Pos = AdjustPosInTrunk(hvc.Pos, obj.Pos, data.HeldItem.Shape)
							shadowObj.Layer = 2
							shadowImg := img.NewSprite(data.HeldItem.Key, constants.TestBatch)
							shadowImg.Color = pixel.ToRGBA(color.RGBA{
								R: 255,
								G: 255,
								B: 255,
								A: 150,
							})
							shadow := myecs.Manager.NewEntity()
							shadow.AddComponent(myecs.Drawable, shadowImg).
								AddComponent(myecs.Object, shadowObj).
								AddComponent(myecs.Temp, myecs.ClearFlag(true))
						}
					}
				}))
		}
	}
}

func TrunkClean() {
	claimed = false
}

func AdjustPosInTrunk(inPos, objPos pixel.Vec, shape []world.Coords) pixel.Vec {
	rPos := objPos
	if world.Width(shape)%2 == 0 {
		if inPos.X-objPos.X < 0 {
			rPos.X -= world.TileSize * 0.5
		} else {
			rPos.X += world.TileSize * 0.5
		}
	}
	if world.Height(shape)%2 == 0 {
		if inPos.Y-objPos.Y < 0 {
			rPos.Y -= world.TileSize * 0.5
		} else {
			rPos.Y += world.TileSize * 0.5
		}
	}
	return rPos
}

func TrunkPlacement(orig world.Coords, shape []world.Coords) (bool, int) {
	mShape := constants.GetMovedCoords(orig, shape)
	for z := 0; z < data.Truck.Height; z++ {
		legal := true
		for _, c := range mShape {
			if !LegalTrunkCoords(c, z) {
				legal = false
				break
			}
		}
		if legal {
			return true, z
		}
	}
	return false, 0
}

func LegalTrunkCoords(c world.Coords, d int) bool {
	//Check if inside the truck
	if c.X >= data.Truck.Width || c.X < 0 || c.Y >= data.Truck.Depth || c.Y < 0 || d >= data.Truck.Height || d < 0 {
		return false
	}
	//Check if something else is occupying the space or a space above
	for z := d; z < data.Truck.Height; z++ {
		if data.Truck.Trunk[z][c.Y][c.X] {
			return false
		}
	}
	return true
}

func PlaceInTrunk(orig world.Coords, item *data.Item) (bool, int) {
	mShape := constants.GetMovedCoords(orig, item.Shape)
	for z := 0; z < data.Truck.Height; z++ {
		legal := true
		for _, c := range mShape {
			if !LegalTrunkCoords(c, z) {
				legal = false
				break
			}
		}
		if legal {
			for _, c := range mShape {
				data.Truck.Trunk[z][c.Y][c.X] = true
				item.TrunkC = append(item.TrunkC, c)
			}
			item.TrunkZ = z
			item.TIndex = len(data.Truck.Wares)
			data.Truck.Wares = append(data.Truck.Wares, item)
			UpdateTrunk()
			return true, z
		}
	}
	return false, 0
}

func UpdateTrunk() {
	for i, ware := range data.Truck.Wares {
		ware.TIndex = i
		ware.Buried = false
		ware.Sprite.Color = pixel.RGB(1, 1, 1)
		if ware.TrunkZ == data.Truck.Height-1 {
			continue
		}
		for _, c := range ware.TrunkC {
			if !LegalTrunkCoords(c, ware.TrunkZ+1) {
				ware.Buried = true
				ware.Sprite.Color = pixel.RGB(0.9, 0.9, 0.9)
			}
		}
	}
}
