package data

import (
	"github.com/bytearena/ecs"
	"github.com/thoas/go-funk"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/typeface"
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

var (
	WareNameLabelOne   *typeface.Text
	WareNameLabelTwo   *typeface.Text
	WareNameLabelThree *typeface.Text
)

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
		return x.Active && NotInQueue(x)
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
			Name:      "Apple",
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.Apple,
			Value:     2 + constants.GlobalSeededRandom.Intn(10),
			Active:    true,
		}, {
			Name:      "Wanna Banana?",
			ShapeKey:  constants.SizeTwobyOne,
			SpriteKey: constants.Banana,
			Value:     2 + constants.GlobalSeededRandom.Intn(10),
			Active:    true,
		}, {
			Name:      "Memory Leeks",
			ShapeKey:  constants.SizeThreebyOne,
			SpriteKey: constants.Leaks,
			Value:     2 + constants.GlobalSeededRandom.Intn(10),
			Active:    true,
		}, {
			Name:      "Couple of Balls",
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.GolfBalls,
			Value:     5 + constants.GlobalSeededRandom.Intn(5),
			Active:    true,
		}, {
			Name:      "Club of Golf",
			ShapeKey:  constants.SizeFourbyOne,
			SpriteKey: constants.GolfClub,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Bag of Golf",
			ShapeKey:  constants.SizeFourbyTwo,
			SpriteKey: constants.GolfBag,
			Value:     100 + constants.GlobalSeededRandom.Intn(100),
			Active:    true,
		}, {
			Name:      "Hoover",
			ShapeKey:  constants.SizeTwobyFourL,
			SpriteKey: constants.Vacuum,
			Value:     150 + constants.GlobalSeededRandom.Intn(100),
			Active:    true,
		}, {
			Name:      "Sweeper",
			ShapeKey:  constants.SizeOnebyFour,
			SpriteKey: constants.Broom,
			Value:     25 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Ee Bagette",
			ShapeKey:  constants.SizeOnebyThree,
			SpriteKey: constants.FrenchBread,
			Value:     5 + constants.GlobalSeededRandom.Intn(25),
			Active:    true,
		}, {
			Name:      "Just a Dustpan",
			ShapeKey:  constants.SizeTwobyTwo,
			SpriteKey: constants.Dustpan,
			Value:     5 + constants.GlobalSeededRandom.Intn(20),
			Active:    true,
		}, {
			Name:      "Crystal Ball",
			ShapeKey:  constants.SizeTwobyTwo,
			SpriteKey: constants.CrystalBall,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Suspicious Stone",
			ShapeKey:  constants.SizeOnebyOne,
			SpriteKey: constants.JustARock,
			Value:     0,
			Active:    true,
		}, {
			Name:      "Zee Snake Oil",
			ShapeKey:  constants.SizeOnebyTwo,
			SpriteKey: constants.SnakeOil,
			Value:     50 + constants.GlobalSeededRandom.Intn(200),
			Active:    true,
		}, {
			Name:      "Ordinary Hat",
			ShapeKey:  constants.SizeThreebyTwo,
			SpriteKey: constants.TopHat,
			Value:     10 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Live Anima... Just Luggage",
			ShapeKey:  constants.SizeFourbyThree,
			SpriteKey: constants.SteamerTrunk,
			Value:     50 + constants.GlobalSeededRandom.Intn(1000),
			Active:    true,
		}, {
			Name:      "Books of Knowledge",
			ShapeKey:  constants.SizeTwobyThree,
			SpriteKey: constants.Encyclopedias,
			Value:     100 + constants.GlobalSeededRandom.Intn(250),
			Active:    true,
		}, {
			Name:      "Out of the Frying Pan",
			ShapeKey:  constants.SizeTwobyThreeP,
			SpriteKey: constants.FryingPan,
			Value:     20 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Boom goes the Dynamite",
			ShapeKey:  constants.SizeOnebyTwo,
			SpriteKey: constants.Dynamite,
			Value:     50 + constants.GlobalSeededRandom.Intn(100),
			Active:    true,
		}, {
			Name:      "Globe",
			ShapeKey:  constants.SizeTwobyTwo,
			SpriteKey: constants.Globe,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Whiskey",
			ShapeKey:  constants.SizeOnebyTwo,
			SpriteKey: constants.Whiskey,
			Value:     50 + constants.GlobalSeededRandom.Intn(100),
			Active:    true,
		}, {
			Name:      "Sarsaparilla",
			ShapeKey:  constants.SizeOnebyTwo,
			SpriteKey: constants.Sarsaparilla,
			Value:     0 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Digger",
			ShapeKey:  constants.SizeTwobyOne,
			SpriteKey: constants.Shovel,
			Value:     10 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Rifle",
			ShapeKey:  constants.SizeFourbyOne,
			SpriteKey: constants.Rifle,
			Value:     50 + constants.GlobalSeededRandom.Intn(20),
			Active:    true,
		}, {
			Name:      "Claranet",
			ShapeKey:  constants.SizeThreebyOne,
			SpriteKey: constants.Claranet,
			Value:     50 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "Cowboy Hat",
			ShapeKey:  constants.SizeThreebyTwoUpsideDownT,
			SpriteKey: constants.CowboyHat,
			Value:     10 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		}, {
			Name:      "My Turtle",
			ShapeKey:  constants.SizeTwobyOne,
			SpriteKey: constants.Turtle,
			Value:     35 + constants.GlobalSeededRandom.Intn(50),
			Active:    true,
		},
	}
)
