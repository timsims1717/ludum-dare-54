package world

import (
	"github.com/faiface/pixel"
)

var (
	TileSize  float64
	ITileSize int
	Origin    = Coords{
		X: 0,
		Y: 0,
	}
	TileRect pixel.Rect
)

func SetTileSize(s float64) {
	TileSize = s
	TileRect = pixel.R(0, 0, s, s)
	ITileSize = int(s)
}

func MapToWorld(a Coords) pixel.Vec {
	return pixel.V(float64(a.X)*TileSize, float64(a.Y)*TileSize)
}

func WorldToMap(x, y float64) (int, int) {
	rx := int(x / TileSize)
	ry := int(y / TileSize)
	if x < 0 {
		rx -= 1
	}
	if y < 0 {
		ry -= 1
	}
	return rx, ry
}

func WorldToMapAdj(x, y float64) (int, int) {
	ix := x + TileSize*0.5
	iy := y + TileSize*0.5
	rx := int(ix / TileSize)
	ry := int(iy / TileSize)
	if ix < 0 {
		rx -= 1
	}
	if iy < 0 {
		ry -= 1
	}
	return rx, ry
}
