package constants

type Sprites int

const (
	UndefinedSprite = iota
	Green
	Yellow
	Blue
	Purple
	Turquoise
	Brown
	Indigo
	Red
	Gold
	White
	Orange
	Gray
)

func (t Sprites) String() string {
	switch t {
	case Green:
		return "green"
	case Yellow:
		return "yellow"
	case Blue:
		return "blue"
	case Purple:
		return "purple"
	case Turquoise:
		return "turquoise"
	case Brown:
		return "brown"
	case Indigo:
		return "indigo"
	case Red:
		return "red"
	case Gold:
		return "gold"
	case White:
		return "white"
	case Orange:
		return "orange"
	case Gray:
		return "gray"
	case UndefinedSprite:
		return "Undefined"
	}
	return ""
}
