package systems

import (
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/util"
	"ludum-dare-54/pkg/world"
	"math"
)

func QueueSystem() {
	for i := 0; i < 8; i++ {
		if data.ItemQueue[i] == nil {
			found := false
			for j := i + 1; j < 8; j++ {
				if data.ItemQueue[j] != nil {
					item := data.ItemQueue[j]
					ip := object.NewInterpolation(object.InterpolateY).
						AddGween(item.Object.Pos.Y, rightQueueY(i), 0.4, ease.OutCubic)
					item.Entity.AddComponent(myecs.Interpolation, ip)
					item.QueueIndex = i
					data.ItemQueue[i], data.ItemQueue[j] = item, nil
					found = true
					break
				}
			}
			if !found {
				localWare := data.Wares[constants.GlobalSeededRandom.Intn(len(data.Wares))].CopyWare()
				localWare.QueueIndex = i
				obj := object.New().WithID("item-in-queue")
				obj.Pos = pixel.V(slotX, rightQueueY(i))
				obj.Layer = 15
				spr := img.Batchers[constants.TestBatch].Sprites[localWare.SpriteKey]
				obj.SetRect(spr.Frame())
				sca := slotSize * 0.9 / math.Max(obj.Rect.W(), obj.Rect.H())
				obj.Sca = pixel.V(sca, sca)
				localWare.Object = obj
				localWare.Entity = myecs.Manager.NewEntity()
				localWare.Entity.AddComponent(myecs.Object, localWare.Object).
					AddComponent(myecs.Drawable, localWare.Sprite).
					AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.GameView, func(hvc *data.HoverClick) {
						if data.HeldItem != nil && data.HeldItem.Object.ID == localWare.Object.ID {
							// Drag system takes care of the movement of the item
							if !localWare.Entity.HasComponent(myecs.Drag) {
								// the item must have just lost its drag component
								x, y := world.WorldToMapAdj(localWare.Object.Pos.X, localWare.Object.Pos.Y)
								x2, y2 := world.WorldToMap(localWare.Object.Pos.X, localWare.Object.Pos.Y)
								if world.Width(localWare.Shape)%2 == 0 {
									x = x2
								}
								if world.Height(localWare.Shape)%2 == 0 {
									y = y2
								}
								legal, layer := PlaceInTrunk(world.Coords{X: x, Y: y}, localWare)
								if legal {
									localWare.Object.Layer = layer + 1
									localWare.Object.Pos = world.MapToWorld(world.Coords{X: x, Y: y})
									if world.Width(localWare.Shape)%2 == 0 {
										localWare.Object.Pos.X += world.TileSize * 0.5
									}
									if world.Height(localWare.Shape)%2 == 0 {
										localWare.Object.Pos.Y += world.TileSize * 0.5
									}
									localWare.Object.Pos.Y += float64(localWare.TrunkZ) * 5.
								} else {
									// move to the empty part of the screen
									pos := GetNearestPos(localWare.Object.Pos, localWare.Object.Rect)
									ips := []*object.Interpolation{
										object.NewInterpolation(object.InterpolateY).
											AddGween(localWare.Object.Pos.Y, pos.Y, 0.4, ease.OutCubic),
										object.NewInterpolation(object.InterpolateX).
											AddGween(localWare.Object.Pos.X, pos.X, 0.4, ease.OutCubic),
									}
									localWare.Entity.AddComponent(myecs.Interpolation, ips)
								}
								data.HeldItem = nil
								hvc.Input.Get("click").Consume()
							}
						} else if hvc.Hover && hvc.Input.Get("click").JustPressed() && !localWare.Buried &&
							data.HeldItem == nil && !localWare.Entity.HasComponent(myecs.Interpolation) {
							localWare.Entity.AddComponent(myecs.Drag, &data.DragTimer{
								Timer: timing.New(0.2),
							})
							data.HeldItem = localWare
							localWare.Object.Layer = 20
							if localWare.QueueIndex > -1 {
								data.ItemQueue[localWare.QueueIndex] = nil
								localWare.QueueIndex = -1
								localWare.Object.Sca = pixel.V(1, 1)
							}
							if localWare.TIndex > -1 {
								for _, c := range localWare.TrunkC {
									data.CurrentTruck.Trunk[localWare.TrunkZ][c.Y][c.X] = false
								}
								localWare.TrunkC = []world.Coords{}
								if len(data.CurrentTruck.Wares) > 1 {
									data.CurrentTruck.Wares = append(data.CurrentTruck.Wares[:localWare.TIndex], data.CurrentTruck.Wares[localWare.TIndex+1:]...)
								} else {
									data.CurrentTruck.Wares = []*data.Ware{}
								}
								UpdateTrunk()
								localWare.TIndex = -1
							}
						}
					}))
				data.ItemQueue[i] = localWare
			}
		}
	}
}

var (
	bottomSlot = -40.
	slotSize   = 64.
	slotX      = 240.
)

func rightQueueY(i int) float64 {
	return -bottomSlot + (float64(i) * slotSize)
}

func GetNearestPos(pos pixel.Vec, r pixel.Rect) pixel.Vec {
	nPos := pos
	if !data.BottomDrop.Contains(pos) && !data.LeftDrop.Contains(pos) {
		bdx := 0.
		if pos.X > data.BottomDrop.Max.X {
			bdx = pos.X - data.BottomDrop.Max.X
		} else if pos.X < data.BottomDrop.Min.X {
			bdx = pos.X - data.BottomDrop.Min.X
		}
		bdy := 0.
		if pos.Y > data.BottomDrop.Max.Y {
			bdy = pos.Y - data.BottomDrop.Max.Y
		} else if pos.Y < data.BottomDrop.Min.Y {
			bdy = pos.Y - data.BottomDrop.Min.Y
		}
		ldx := 0.
		if pos.X > data.LeftDrop.Max.X {
			ldx = pos.X - data.LeftDrop.Max.X
		} else if pos.X < data.LeftDrop.Min.X {
			ldx = pos.X - data.LeftDrop.Min.X
		}
		ldy := 0.
		if pos.Y > data.LeftDrop.Max.Y {
			ldy = pos.Y - data.LeftDrop.Max.Y
		} else if pos.Y < data.LeftDrop.Min.Y {
			ldy = pos.Y - data.LeftDrop.Min.Y
		}
		bd := math.Sqrt(bdx*bdx + bdy*bdy)
		ld := math.Sqrt(ldx*ldx + ldy*ldy)
		if bd > ld {
			if ldx > 0 {
				nPos.X = data.LeftDrop.Max.X
			} else if ldx < 0 {
				nPos.X = data.LeftDrop.Min.X
			}
			if ldy > 0 {
				nPos.Y = data.LeftDrop.Max.Y
			} else if ldy < 0 {
				nPos.Y = data.LeftDrop.Min.Y
			}
		} else {
			if bdx > 0 {
				nPos.X = data.BottomDrop.Max.X
			} else if bdx < 0 {
				nPos.X = data.BottomDrop.Min.X
			}
			if bdy > 0 {
				nPos.Y = data.BottomDrop.Max.Y
			} else if bdy < 0 {
				nPos.Y = data.BottomDrop.Min.Y
			}
		}
	}
	if data.BottomDrop.Contains(nPos) {
		return util.ConstrainR(nPos, data.BottomDrop.Center(), r, data.BottomDrop)
	}
	if data.LeftDrop.Contains(nPos) {
		return util.ConstrainR(nPos, data.LeftDrop.Center(), r, data.LeftDrop)
	}
	return pixel.V(-100, -100)
}

func LeavePackingSystem() {
	if data.LeavePacking {

	}
}