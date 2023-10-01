package data

import (
	"github.com/bytearena/ecs"
	"github.com/thoas/go-funk"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/world"
)

type Ware struct {
	Active     bool
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
func GetTotalActiveWares() int {
	filteredWares := funk.Filter(Wares, func(x *Ware) bool {
		return x.Active
	}).([]*Ware)
	return len(filteredWares)
}

func GetRandomWare() *Ware {
	filteredWares := funk.Filter(Wares, func(x *Ware) bool {
		return x.Active
	}).([]*Ware)
	randWareItr := constants.GlobalSeededRandom.Intn(len(filteredWares))
	return &Ware{
		Name:      filteredWares[randWareItr].Name,
		ShapeKey:  filteredWares[randWareItr].ShapeKey,
		SpriteKey: filteredWares[randWareItr].SpriteKey,
		Value:     filteredWares[randWareItr].Value,
		Sprite:    img.NewSprite(filteredWares[randWareItr].SpriteKey, constants.TestBatch),
		Shape:     constants.TrunkominoTypes[filteredWares[randWareItr].ShapeKey],
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
			Active:    false,
		}, {
			Name:      constants.Yellow,
			ShapeKey:  constants.SizeTwobyOne,
			SpriteKey: constants.Yellow,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Blue,
			ShapeKey:  constants.SizeOnebyTwo,
			SpriteKey: constants.Blue,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Purple,
			ShapeKey:  constants.SizeTwobyTwo,
			SpriteKey: constants.Purple,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Turquoise,
			ShapeKey:  constants.SizeOnebyThree,
			SpriteKey: constants.Turquoise,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Brown,
			ShapeKey:  constants.SizeTwobyFour,
			SpriteKey: constants.Brown,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Indigo,
			ShapeKey:  constants.SizeOnebyFour,
			SpriteKey: constants.Indigo,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Red,
			ShapeKey:  constants.SizeThreebyOne,
			SpriteKey: constants.Red,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Gold,
			ShapeKey:  constants.SizeFourbyOne,
			SpriteKey: constants.Gold,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.White,
			ShapeKey:  constants.SizeTwobyThree,
			SpriteKey: constants.White,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Orange,
			ShapeKey:  constants.SizeFourbyTwo,
			SpriteKey: constants.Orange,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Gray,
			ShapeKey:  constants.SizeThreebyTwo,
			SpriteKey: constants.Gray,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    false,
		}, {
			Name:      constants.Apple,
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.Apple,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.Banana,
			ShapeKey:  constants.SizeTwobyOne,
			SpriteKey: constants.Banana,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.Leaks,
			ShapeKey:  constants.SizeThreebyOne,
			SpriteKey: constants.Leaks,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.GolfBalls,
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.GolfBalls,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.GolfClub,
			ShapeKey:  constants.SizeFourbyOne,
			SpriteKey: constants.GolfClub,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.GolfBag,
			ShapeKey:  constants.SizeFourbyTwo,
			SpriteKey: constants.GolfBag,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.Vacuum,
			ShapeKey:  constants.SizeTwobyFourL,
			SpriteKey: constants.Vacuum,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.Broom,
			ShapeKey:  constants.SizeOnebyFour,
			SpriteKey: constants.Broom,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.FrenchBread,
			ShapeKey:  constants.SizeOnebyThree,
			SpriteKey: constants.FrenchBread,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.Dustpan,
			ShapeKey:  constants.SizeTwobyTwo,
			SpriteKey: constants.Dustpan,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.CrystalBall,
			ShapeKey:  constants.SizeTwobyTwo,
			SpriteKey: constants.CrystalBall,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.JustARock,
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.JustARock,
			Value:     0 + constants.GlobalSeededRandom.Intn(10),
			Active:    true,
		}, {
			Name:      constants.SnakeOil,
			ShapeKey:  constants.SizeOnebyTwo,
			SpriteKey: constants.SnakeOil,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      constants.TopHat,
			ShapeKey:  constants.SizeThreebyTwo,
			SpriteKey: constants.TopHat,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		},
	}
)
