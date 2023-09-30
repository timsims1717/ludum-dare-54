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
	SpriteKey  constants.Sprites
	Value      int
	Name       string
	QueueIndex int
	Object     *object.Object
	Entity     *ecs.Entity
	Sprite     *img.Sprite
	Shape      []world.Coords
}

func (w *Ware) CopyWare() *Ware {
	return &Ware{
		Name:      w.Name,
		ShapeKey:  w.ShapeKey,
		SpriteKey: w.SpriteKey,
		Value:     w.Value,
		Sprite:    img.NewSprite(w.SpriteKey.String(), constants.TestBatch),
		Shape:     constants.TrunkominoTypes[w.ShapeKey],
	}
}

var (
	Wares = []*Ware{
		{
			Name:      "Green",
			ShapeKey:  constants.TinyBox,
			SpriteKey: constants.Green,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Yellow",
			ShapeKey:  constants.ShortLogOver,
			SpriteKey: constants.Yellow,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Blue",
			ShapeKey:  constants.ShortLogUp,
			SpriteKey: constants.Blue,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Purple",
			ShapeKey:  constants.MediumBox,
			SpriteKey: constants.Purple,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Turquoise",
			ShapeKey:  constants.MediumLogUp,
			SpriteKey: constants.Turquoise,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Brown",
			ShapeKey:  constants.LargeUpBox,
			SpriteKey: constants.Brown,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Indigo",
			ShapeKey:  constants.LongLogUp,
			SpriteKey: constants.Indigo,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Red",
			ShapeKey:  constants.MediumLogOver,
			SpriteKey: constants.Red,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Gold",
			ShapeKey:  constants.LongLogOver,
			SpriteKey: constants.Gold,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "White",
			ShapeKey:  constants.MediumUpBox,
			SpriteKey: constants.White,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Orange",
			ShapeKey:  constants.LargeLongBox,
			SpriteKey: constants.Orange,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		}, {
			Name:      "Gray",
			ShapeKey:  constants.MediumLongBox,
			SpriteKey: constants.Gray,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
		},
	}
)
