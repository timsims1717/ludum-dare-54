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
			ShapeKey:  constants.TinyBox,
			SpriteKey: constants.Green,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Yellow,
			ShapeKey:  constants.ShortLogOver,
			SpriteKey: constants.Yellow,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Blue,
			ShapeKey:  constants.ShortLogUp,
			SpriteKey: constants.Blue,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Purple,
			ShapeKey:  constants.MediumBox,
			SpriteKey: constants.Purple,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Turquoise,
			ShapeKey:  constants.MediumLogUp,
			SpriteKey: constants.Turquoise,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Brown,
			ShapeKey:  constants.LargeUpBox,
			SpriteKey: constants.Brown,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Indigo,
			ShapeKey:  constants.LongLogUp,
			SpriteKey: constants.Indigo,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Red,
			ShapeKey:  constants.MediumLogOver,
			SpriteKey: constants.Red,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Gold,
			ShapeKey:  constants.LongLogOver,
			SpriteKey: constants.Gold,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.White,
			ShapeKey:  constants.MediumUpBox,
			SpriteKey: constants.White,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Orange,
			ShapeKey:  constants.LargeLongBox,
			SpriteKey: constants.Orange,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Gray,
			ShapeKey:  constants.MediumLongBox,
			SpriteKey: constants.Gray,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Apple,
			ShapeKey:  constants.TinyBox,
			SpriteKey: constants.Apple,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      constants.Banana,
			ShapeKey:  constants.ShortLogOver,
			SpriteKey: constants.Banana,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		},
	}
)
