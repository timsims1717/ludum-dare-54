package data

import (
	"github.com/bytearena/ecs"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
)

var (
	CurrentTruck   *Truck
	PickedTruckKey constants.TruckTypes
)

type Truck struct {
	Trunk         [][][]bool
	Wares         []*Ware
	Width         int
	Depth         int
	Height        int
	TotalSpace    int
	FilledSpace   int
	PercentFilled int
	TruckLabel    string
	SpriteKey     string
	CurrHeight    int
	TileMap       []*img.Sprite
	TileObject    *object.Object
	TileEntity    *ecs.Entity
}

func NewTruck(w, d, h int) {
	CurrentTruck = &Truck{}
	CurrentTruck.Width = w
	CurrentTruck.Depth = d
	CurrentTruck.Height = h
	CurrentTruck.TotalSpace = w * d * h
	CurrentTruck.TruckLabel = constants.UndefinedTruckType
	for z := 0; z < h; z++ {
		CurrentTruck.Trunk = append(CurrentTruck.Trunk, [][]bool{})
		for y := 0; y < d; y++ {
			CurrentTruck.Trunk[z] = append(CurrentTruck.Trunk[z], []bool{})
			for x := 0; x < w; x++ {
				CurrentTruck.Trunk[z][y] = append(CurrentTruck.Trunk[z][y], false)
			}
		}
	}
}

func (t *Truck) CopyTruck() {
	CurrentTruck = &Truck{
		Width:      t.Width,
		Height:     t.Height,
		Depth:      t.Depth,
		TotalSpace: t.Width * t.Height * t.Depth,
		TruckLabel: t.TruckLabel,
	}
	for z := 0; z < t.Height; z++ {
		CurrentTruck.Trunk = append(CurrentTruck.Trunk, [][]bool{})
		for y := 0; y < t.Depth; y++ {
			CurrentTruck.Trunk[z] = append(CurrentTruck.Trunk[z], []bool{})
			for x := 0; x < t.Width; x++ {
				CurrentTruck.Trunk[z][y] = append(CurrentTruck.Trunk[z][y], false)
			}
		}
	}
}

var (
	AvailableTrucks = map[string]*Truck{
		constants.SmartCar: {
			Width:      5,
			Depth:      4,
			Height:     3,
			TruckLabel: "SMART CAR",
			SpriteKey:  "smart",
		}, constants.Minivan: {
			Width:      5,
			Depth:      5,
			Height:     4,
			TruckLabel: "MINIVAN",
			SpriteKey:  "mini",
		}, constants.CargoVan: {
			Width:      5,
			Depth:      6,
			Height:     5,
			TruckLabel: "CARGO VAN",
			SpriteKey:  "cargo",
		}, constants.SemiTruck: {
			Width:      5,
			Depth:      7,
			Height:     6,
			TruckLabel: "BOX TRUCK",
			SpriteKey:  "box",
		}, constants.Wagon: {
			Width:      4,
			Depth:      7,
			Height:     5,
			TruckLabel: "CONASTOGA",
			SpriteKey:  "wagon",
		},
	}
)
