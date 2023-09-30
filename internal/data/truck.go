package data

var (
	Truck *truck
)

type truck struct {
	Trunk  [][][]bool
	Wares  []*Item
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
		Truck.Trunk = append(Truck.Trunk, [][]bool{})
		for y := 0; y < d; y++ {
			Truck.Trunk[z] = append(Truck.Trunk[z], []bool{})
			for x := 0; x < w; x++ {
				Truck.Trunk[z][y] = append(Truck.Trunk[z][y], false)
			}
		}
	}
}
