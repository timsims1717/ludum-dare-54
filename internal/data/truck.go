package data

import "ludum-dare-54/internal/constants"

var (
	CurrentTruck *Truck
)

type Truck struct {
	Trunk  [][][]bool
	Wares  []*Ware
	Width  int
	Depth  int
	Height int
}

func NewTruck(w, d, h int) {
	CurrentTruck = &Truck{}
	CurrentTruck.Width = w
	CurrentTruck.Depth = d
	CurrentTruck.Height = h
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
		Width:  t.Width,
		Height: t.Height,
		Depth:  t.Depth,
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
	AvalibleTrucks = map[string]*Truck{
		constants.SmartCar: {
			Width:  5,
			Height: 3,
			Depth:  3,
		}, constants.Minivan: {
			Width:  5,
			Height: 4,
			Depth:  5,
		}, constants.CargoVan: {
			Width:  5,
			Height: 5,
			Depth:  5,
		}, constants.SemiTruck: {
			Width:  10,
			Height: 5,
			Depth:  5,
		},
	}
)
