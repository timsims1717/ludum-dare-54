package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/states"
	"ludum-dare-54/pkg/debug"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/viewport"
	"ludum-dare-54/pkg/world"
)

func run() {
	world.SetTileSize(32)
	cfg := pixelgl.WindowConfig{
		Title:  "Gem Runner",
		Bounds: pixel.R(0, 0, 1600, 900),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	viewport.MainCamera = viewport.New(win.Canvas())
	viewport.MainCamera.SetRect(pixel.R(0, 0, 1600, 900))
	viewport.MainCamera.CamPos = pixel.V(1600*0.5, 900*0.5)

	state.Register(states.PackingStateKey, state.New(states.PackingState))

	testSheet, err := img.LoadSpriteSheet("assets/test1.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(constants.TestBatch, testSheet, true, true)

	debug.Initialize(&viewport.MainCamera.PostCamPos)
	debug.Text = true

	//systems.InitMainBorder()

	win.Show()
	timing.Reset()
	for !win.Closed() {
		timing.Update()
		debug.Clear()
		data.DebugInput.Update(win, viewport.MainCamera.Mat)

		//options.WindowUpdate(win)
		//if options.Updated {
		//	viewport.MainCamera.CamPos = pixel.V(viewport.MainCamera.Rect.W()*0.5, viewport.MainCamera.Rect.H()*0.5)
		//}

		state.Update(win)
		viewport.MainCamera.Update()
		state.Draw(win)

		//win.SetSmooth(false)
		debug.Draw(win)
		//win.SetSmooth(true)
		//win.SetSmooth(options.BilinearFilter)

		//sfx.MusicPlayer.Update()
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
