package constants

import "ludum-dare-54/pkg/world"

var (
	//Trunkominos
	TrunkominoTypes = map[TrunkominoType][]world.Coords{
		SizeOnebyOne:   {{0, 0}},
		SizeTwobyOne:   {{0, 0}, {1, 0}},
		SizeOnebyTwo:   {{0, 0}, {0, 1}},
		SizeThreebyOne: {{0, 0}, {1, 0}, {2, 0}},
		SizeOnebyThree: {{0, 0}, {0, 1}, {0, 2}},
		SizeFourbyOne:  {{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		SizeOnebyFour:  {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		SizeTwobyTwo:   {{0, 0}, {1, 0}, {0, 1}, {1, 1}},
		SizeThreebyTwo: {{0, 0}, {1, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}},
		SizeFourbyTwo:  {{0, 0}, {1, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}, {3, 0}, {3, 1}},
		SizeTwobyThree: {{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}, {1, 2}},
		SizeTwobyFour:  {{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}, {1, 2}, {0, 3}, {1, 3}},
		SizeTwobyFourL: {{0, 0}, {1, 0}, {0, 1}, {0, 2}, {0, 3}},
	}
)

type TrunkominoType string

const (
	UndefinedTrunkominoType = "Undefined"
	SizeOnebyOne            = "SizeOnebyOne"
	SizeTwobyOne            = "SizeTwobyOne"
	SizeOnebyTwo            = "SizeOnebyTwo"
	SizeThreebyOne          = "SizeThreebyOne"
	SizeOnebyThree          = "SizeOnebyThree"
	SizeFourbyOne           = "SizeFourbyOne"
	SizeOnebyFour           = "SizeOnebyFour"
	SizeTwobyTwo            = "SizeTwobyTwo"
	SizeThreebyTwo          = "SizeThreebyTwo"
	SizeFourbyTwo           = "SizeFourbyTwo"
	SizeTwobyThree          = "SizeTwobyThree"
	SizeTwobyFour           = "SizeTwobyFour"
	SizeTwobyFourL          = "SizeTwobyFourL"
)

// GetMovedCoords assumes the center of even numbers is moved to the left and
// down.
func GetMovedCoords(center world.Coords, a []world.Coords) []world.Coords {
	w := world.Width(a)
	h := world.Height(a)
	x := -((w - 1) / 2)
	y := -((h - 1) / 2)
	n := make([]world.Coords, len(a))
	for i := 0; i < len(a); i++ {
		n[i] = a[i]
		n[i].X += x + center.X
		n[i].Y += y + center.Y
	}
	return n
}
