package systems

import (
	"fmt"
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

func CreateTruck() {
	if data.CurrentTruck == nil {
		data.AvailableTrucks[string(data.PickedTruckKey)].CopyTruck()
	}
	// build the truck
	for yt := -1; yt < data.CurrentTruck.Depth+1; yt++ {
		for xt := -1; xt < data.CurrentTruck.Width+1; xt++ {
			str := ""
			if yt == -1 {
				if xt == -1 {
					str = "fender_rl"
				} else if xt == 0 {
					str = "tire_l"
				} else if xt == data.CurrentTruck.Width-1 {
					str = "tire_r"
				} else if xt == data.CurrentTruck.Width {
					str = "fender_rr"
				} else {
					str = "bottom"
				}
			} else if yt == data.CurrentTruck.Depth {
				if xt == -1 {
					str = "fender_fl"
				} else if xt == 0 {
					str = "driver"
				} else if xt == 1 {
					str = ""
				} else if xt == data.CurrentTruck.Width-2 {
					str = "passenger"
				} else if xt == data.CurrentTruck.Width-1 {
					str = ""
				} else if xt == data.CurrentTruck.Width {
					str = "fender_fr"
				} else {
					str = "center"
				}
			} else {
				if xt == -1 {
					str = "side_l"
				} else if xt == data.CurrentTruck.Width {
					str = "side_r"
				}
			}
			if str != "" {
				offset := pixel.V(float64(xt)*world.TileSize, float64(yt)*world.TileSize)
				if yt == data.CurrentTruck.Depth {
					if xt == 0 || xt == data.CurrentTruck.Width-2 {
						offset.X += world.TileSize * 0.5
						offset.Y += world.TileSize * 0.5
					} else if xt > 0 && xt < data.CurrentTruck.Width-2 {
						offset.Y += world.TileSize * 0.5
					}
				}
				spr := img.NewOffsetSprite(fmt.Sprintf("%s_%s", data.CurrentTruck.SpriteKey, str), constants.TestBatch, offset)
				data.CurrentTruck.TileMap = append(data.CurrentTruck.TileMap, spr)
			}
		}
	}
	data.CurrentTruck.TileObject = object.New()
	data.CurrentTruck.TileObject.Layer = -1
	data.CurrentTruck.TileEntity = myecs.Manager.NewEntity()
	data.CurrentTruck.TileEntity.AddComponent(myecs.Object, data.CurrentTruck.TileObject).
		AddComponent(myecs.Drawable, data.CurrentTruck.TileMap)
	// build the tiles
	for yt := 0; yt < data.CurrentTruck.Depth; yt++ {
		for xt := 0; xt < data.CurrentTruck.Width; xt++ {
			obj := object.New().WithID("slot")
			obj.Pos = pixel.V(float64(xt)*world.TileSize, float64(yt)*world.TileSize)
			obj.SetRect(pixel.R(0, 0, world.TileSize, world.TileSize))
			obj.Layer = 0
			e := myecs.Manager.NewEntity()
			e.AddComponent(myecs.Drawable, img.NewSprite(fmt.Sprintf("%s_tile", data.CurrentTruck.SpriteKey), constants.TestBatch)).
				AddComponent(myecs.Object, obj).
				AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.GameView, func(hvc *data.HoverClick) {
					if hvc.Hover && !claimed {
						claimed = true
						x, y := world.WorldToMapAdj(hvc.Pos.X, hvc.Pos.Y)
						debug.AddIntCoords("Hovered Over", x, y)
						if data.HeldWare != nil {
							shadowObj := object.New()
							// change the pos of items w/even widths and depths
							shadowObj.Pos = AdjustPosInTrunk(hvc.Pos, obj.Pos, data.HeldWare.Shape)
							shadowObj.Layer = 15
							shadowImg := img.NewSprite(data.HeldWare.SpriteKey, constants.TestBatch)
							legal, _ := TrunkHover(world.Coords{X: x, Y: y}, data.HeldWare.Shape)
							if legal {
								shadowImg.Color = pixel.ToRGBA(color.RGBA{
									R: 255,
									G: 255,
									B: 255,
									A: 150,
								})
							} else {
								shadowImg.Color = pixel.ToRGBA(color.RGBA{
									R: 255,
									G: 0,
									B: 0,
									A: 150,
								})
							}
							shadow := myecs.Manager.NewEntity()
							shadow.AddComponent(myecs.Drawable, shadowImg).
								AddComponent(myecs.Object, shadowObj).
								AddComponent(myecs.Temp, myecs.ClearFlag(true))
						}
					}
				}))
			data.TrunkObjects = append(data.TrunkObjects, obj)
			data.TrunkEntities = append(data.TrunkEntities, e)
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

func TrunkHover(orig world.Coords, shape []world.Coords) (bool, int) {
	mShape := constants.GetMovedCoords(orig, shape)
	for z := 0; z < data.CurrentTruck.Height; z++ {
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
	if c.X >= data.CurrentTruck.Width || c.X < 0 || c.Y >= data.CurrentTruck.Depth || c.Y < 0 || d >= data.CurrentTruck.Height || d < 0 {
		return false
	}
	//Check if something else is occupying the space or a space above
	for z := d; z < data.CurrentTruck.Height; z++ {
		if data.CurrentTruck.Trunk[z][c.Y][c.X] {
			return false
		}
	}
	return true
}

func PlaceInTrunk(orig world.Coords, ware *data.Ware) (bool, int) {
	mShape := constants.GetMovedCoords(orig, ware.Shape)
	for z := 0; z < data.CurrentTruck.Height; z++ {
		legal := true
		for _, c := range mShape {
			if !LegalTrunkCoords(c, z) {
				legal = false
				break
			}
		}
		if legal {
			for _, c := range mShape {
				data.CurrentTruck.Trunk[z][c.Y][c.X] = true
				ware.TrunkC = append(ware.TrunkC, c)
			}
			ware.TrunkZ = z
			ware.TIndex = len(data.CurrentTruck.Wares)
			data.CurrentTruck.Wares = append(data.CurrentTruck.Wares, ware)
			UpdateTrunk()
			return true, z
		}
	}
	return false, 0
}

func UpdateTrunk() {
	data.CurrentTruck.FilledSpace = 0
	data.CurrentTruck.CurrHeight = 0
	for i, ware := range data.CurrentTruck.Wares {
		data.CurrentTruck.FilledSpace += len(ware.Shape)
		ware.TIndex = i
		ware.Buried = false
		ware.Sprite.Color = pixel.RGB(1, 1, 1)
		if ware.TrunkZ+1 > data.CurrentTruck.CurrHeight {
			data.CurrentTruck.CurrHeight = ware.TrunkZ + 1
		}
		if ware.TrunkZ == data.CurrentTruck.Height-1 {
			continue
		}
		for _, c := range ware.TrunkC {
			if !LegalTrunkCoords(c, ware.TrunkZ+1) {
				ware.Buried = true
				ware.Sprite.Color = pixel.RGB(0.9, 0.9, 0.9)
			}
		}
	}
	data.CurrentTruck.PercentFilled = int((float64(data.CurrentTruck.FilledSpace) / float64(data.CurrentTruck.TotalSpace)) * 100)
}
