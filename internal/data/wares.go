package data

import (
	"github.com/bytearena/ecs"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/world"
)

type Ware struct {
	ShapeKey   constants.TrunkominoType
	SpriteKey  string
	Value      int
	Name       string
	QueueIndex int
	TIndex     int
	Object     *object.Object
	Entity     *ecs.Entity
	Sprite     *img.Sprite
	Shape      []world.Coords
	TrunkC     []world.Coords
	TrunkZ     int
	Buried     bool
	SellMe     bool
	Sold       bool
}

func (w *Ware) CopyWare() *Ware {
	return &Ware{
		Name:      w.Name,
		ShapeKey:  w.ShapeKey,
		SpriteKey: w.SpriteKey,
		Value:     w.Value,
		Sprite:    img.NewSprite(w.SpriteKey, constants.TestBatch),
		Shape:     constants.TrunkominoTypes[w.ShapeKey],
		TIndex:    -1,
	}
}

var (
	Wares = []*Ware{
		{
			Name:      constants.Green,
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.Green,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Yellow,
			ShapeKey:  constants.SizeTwobyOne,
			SpriteKey: constants.Yellow,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Blue,
			ShapeKey:  constants.SizeOnebyTwo,
			SpriteKey: constants.Blue,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Purple,
			ShapeKey:  constants.SizeTwobyTwo,
			SpriteKey: constants.Purple,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Turquoise,
			ShapeKey:  constants.SizeOnebyThree,
			SpriteKey: constants.Turquoise,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Brown,
			ShapeKey:  constants.SizeTwobyFour,
			SpriteKey: constants.Brown,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Indigo,
			ShapeKey:  constants.SizeOnebyFour,
			SpriteKey: constants.Indigo,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Red,
			ShapeKey:  constants.SizeThreebyOne,
			SpriteKey: constants.Red,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Gold,
			ShapeKey:  constants.SizeFourbyOne,
			SpriteKey: constants.Gold,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.White,
			ShapeKey:  constants.SizeTwobyThree,
			SpriteKey: constants.White,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Orange,
			ShapeKey:  constants.SizeFourbyTwo,
			SpriteKey: constants.Orange,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Gray,
			ShapeKey:  constants.SizeThreebyTwo,
			SpriteKey: constants.Gray,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Apple,
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.Apple,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Banana,
			ShapeKey:  constants.SizeTwobyOne,
			SpriteKey: constants.Banana,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Leaks,
			ShapeKey:  constants.SizeThreebyOne,
			SpriteKey: constants.Leaks,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.GolfBalls,
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.GolfBalls,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.GolfClub,
			ShapeKey:  constants.SizeFourbyOne,
			SpriteKey: constants.GolfClub,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.GolfBag,
			ShapeKey:  constants.SizeFourbyTwo,
			SpriteKey: constants.GolfBag,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Vacuum,
			ShapeKey:  constants.SizeTwobyFourL,
			SpriteKey: constants.Vacuum,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		},
	}
)
