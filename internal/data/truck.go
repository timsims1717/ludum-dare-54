package data

var (
	Truck *truck
)

type truck struct {
	Tiles [][][]bool
}

func NewTruck() {
	Truck = &truck{}
}
