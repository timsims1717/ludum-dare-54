package data

var (
	Truck *truck
)

type truck struct {
	Tiles  [][][]bool
	Width  int
	Depth  int
	Height int
}

func NewTruck(w, d, h int) {
	Truck = &truck{}
	Truck.Width = w
	Truck.Depth = d
	Truck.Height = h
	for z := 0; z < h; z++ {
		Truck.Tiles = append(Truck.Tiles, [][]bool{})
		for y := 0; y < d; y++ {
			Truck.Tiles[z] = append(Truck.Tiles[z], []bool{})
			for x := 0; x < w; x++ {
				Truck.Tiles[z][y] = append(Truck.Tiles[z][y], false)
			}
		}
	}
}
