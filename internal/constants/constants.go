package constants

import (
	"fmt"
	"ludum-dare-54/pkg/world"
	"math/rand"
	"time"
)

const (
	TestBatch = "test_batch"
	Title     = "Well that didn't work"
	Release   = 0
	Version   = 1
	Build     = 20230428

	//Initial Conditions
	InitialTrunkTargetFill = 0.5
	NumberofStrikes        = 3

	//Trunkominos
	NormalizedTrunkominoType = map[[]world.Coords]TrunkominoType{
		[]world.Coords{{0, 0}}:                                                         TinyBox,
		[]world.Coords{{0, 0}, {1, 0}}:                                                 ShortLogDown,
		[]world.Coords{{0, 0}, {0, 1}}:                                                 ShortLogUp,
		[]world.Coords{{0, 0}, {1, 0}, {2, 0}}:                                         MediumLogDown,
		[]world.Coords{{0, 0}, {0, 1}, {0, 2}}:                                         MediumLogUp,
		[]world.Coords{{0, 0}, {1, 0}, {2, 0}, {3, 0}}:                                 LongLogDown,
		[]world.Coords{{0, 0}, {0, 1}, {0, 2}, {0, 3}}:                                 LongLogUp,
		[]world.Coords{{0, 0}, {1, 0}, {0, 1}, {1, 1}}:                                 MediumSquareBox,
		[]world.Coords{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}}:                 MediumLongBox,
		[]world.Coords{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {2, 0}, {2, 1}, {3, 0}, {3, 1}}: LargeLongBox,
		[]world.Coords{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}, {1, 2}}:                 MediumUpBox,
		[]world.Coords{{0, 0}, {1, 0}, {0, 1}, {1, 1}, {0, 2}, {1, 2}, {0, 3}, {1, 3}}: LargeUpBox,
	}
)

var (
	TitleText          = "LD54"
	GlobalSeededRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	TitleVariants      = []string{
		"Is this a Traveling Salesman Problem or a Packing Problem",
		"42 Days no Driving Violations",
		"Fueled by Chicken Wings and Gin",
		"First Stop after the Tetronimo Factory",
		"We need more sales! Everyone's working weekends!",
		"Your truck just barfed all over the loading dock, clean it up!",
		"We killed a bug, YAY!!!",
		"Don't forget to spend your hard earned dollars at the company store!",
		"My nose feels bigger",
		"Why do we even have that lever!?!",
		"Ain’t no fellow who regretted giving it one extra shake, but you can bet every guy has regretted giving one too few.",
	}
)

func RandomTitle() string {
	TitleText = TitleVariants[GlobalSeededRandom.Intn(len(TitleVariants))]
	return TitleText
}

type FailCondition int

const (
	RunOutOfMoney = iota
	TooManyMisses
)

func (fc FailCondition) String() string {
	switch fc {
	case RunOutOfMoney:
		return "Game Over.\nHave run out of money, you better give up and go home"
	case TooManyMisses:
		return fmt.Sprintf("Game Over.\nYou missed sales a %d vendors", NumberofStrikes)
	}
	return ""
}

type TrunkominoType int

const (
	UndefinedTrunkominoType = iota
	TinyBox
	ShortLogDown
	ShortLogUp
	MediumLogDown
	MediumLogUp
	LongLogDown
	LongLogUp
	MediumSquareBox
	MediumLongBox
	LargeLongBox
	MediumUpBox
	LargeUpBox
)

func (t TrunkominoType) String() string {
	switch t {
	case TinyBox:
		return "Tiny Box: 1x1"
	case ShortLogDown:
		return "ShortLogDown: 1x2"
	case ShortLogUp:
		return "ShortLogUp: 2x1"
	case MediumLogDown:
		return "MediumLogDown: 1x3"
	case MediumLogUp:
		return "MediumLogUp: 3x1"
	case LongLogDown:
		return "LongLogDown: 1x4"
	case LongLogUp:
		return "LongLogUp: 4x1"
	case MediumSquareBox:
		return "MediumSquareBox: 2x2"
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
