package constants

import "ludum-dare-54/pkg/world"

var (
	//Trunkominos
	TrunkominoTypes = map[TrunkominoType][]world.Coords{
		TinyBox:       {{0, 0}},
		ShortLogOver:  {{0, 0}, {1, 0}},
		ShortLogUp:    {{0, 0}, {0, 1}},
		MediumLogOver: {{0, 0}, {1, 0}, {2, 0}},
		MediumLogUp:   {{0, 0}, {0, 1}, {0, 2}},
		LongLogOver:   {{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		LongLogUp:     {{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		MediumBox:     {{0, 0}, {1, 0}, {0, 1}, {1, 1}},
		MediumLongBox: {{0, 0}, {1, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}},
		LargeLongBox:  {{0, 0}, {1, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}, {3, 0}, {3, 1}},
		MediumUpBox:   {{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}, {1, 2}},
		LargeUpBox:    {{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}, {1, 2}, {0, 3}, {1, 3}},
	}
)

type TrunkominoType string

const (
	UndefinedTrunkominoType = "Undefined"
	TinyBox                 = "TinyBox"
	ShortLogOver            = "ShortLogOver"
	ShortLogUp              = "ShortLogUp"
	MediumLogOver           = "MediumLogOver"
	MediumLogUp             = "MediumLogUp"
	LongLogOver             = "LongLogOver"
	LongLogUp               = "LongLogUp"
	MediumBox               = "MediumBox"
	MediumLongBox           = "MediumLongBox"
	LargeLongBox            = "LargeLongBox"
	MediumUpBox             = "MediumUpBox"
	LargeUpBox              = "LargeUpBox"
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
