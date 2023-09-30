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
	// Items
	Items = []string{
		"green",
		"yellow",
		"blue",
		"purple",
		"turquoise",
		"brown",
		"indigo",
		"red",
		"gold",
		"white",
		"orange",
		"gray",
	}
	ItemShapes = map[string]TrunkominoType{
		"green":     TinyBox,
		"yellow":    ShortLogOver,
		"blue":      ShortLogUp,
		"purple":    MediumBox,
		"turquoise": MediumLogUp,
		"brown":     LargeUpBox,
		"indigo":    LongLogUp,
		"red":       MediumLogOver,
		"gold":      LongLogOver,
		"white":     MediumUpBox,
		"orange":    LargeLongBox,
		"gray":      MediumLongBox,
	}
)

type TrunkominoType int

const (
	UndefinedTrunkominoType = iota
	TinyBox
	ShortLogOver
	ShortLogUp
	MediumLogOver
	MediumLogUp
	LongLogOver
	LongLogUp
	MediumBox
	MediumLongBox
	LargeLongBox
	MediumUpBox
	LargeUpBox
)

func (t TrunkominoType) String() string {
	switch t {
	case TinyBox:
		return "Tiny Box: 1x1"
	case ShortLogOver:
		return "ShortLogOver: 1x2"
	case ShortLogUp:
		return "ShortLogUp: 2x1"
	case MediumLogOver:
		return "MediumLogOver: 1x3"
	case MediumLogUp:
		return "MediumLogUp: 3x1"
	case LongLogOver:
		return "LongLogOver: 1x4"
	case LongLogUp:
		return "LongLogUp: 4x1"
	case MediumBox:
		return "MediumBox: 2x2"
	case MediumLongBox:
		return "MediumLongBox: 2x3"
	case LargeLongBox:
		return "LargeLongBox: 2x4"
	case MediumUpBox:
		return "MediumUpBox: 3x2"
	case LargeUpBox:
		return "LargeUpBox: 4x2"
	case UndefinedTrunkominoType:
		return "Undefined"
	}
	return ""
}

func GetShape(key string) []world.Coords {
	return TrunkominoTypes[ItemShapes[key]]
}

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
