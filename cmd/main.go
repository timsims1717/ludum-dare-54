package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/states"
	"ludum-dare-54/pkg/debug"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/options"
	"ludum-dare-54/pkg/sfx"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/typeface"
	"ludum-dare-54/pkg/viewport"
	"ludum-dare-54/pkg/world"
)

func run() {
	world.SetTileSize(32)
	cfg := pixelgl.WindowConfig{
		Title:  constants.RandomTitle(),
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

	options.VSync = true

	mainFont, err := typeface.LoadTTF("assets/eurosti.ttf", 200.)
	typeface.Atlases["main"] = text.NewAtlas(mainFont, text.ASCII)

	state.Register(constants.MainMenuStateKey, state.New(states.MainMenuState))
	state.Register(constants.PackingStateKey, state.New(states.PackingState))
	state.Register(constants.TransitionStateKey, state.New(states.TransitionState))
	state.Register(constants.PauseStateKey, state.New(states.PauseState))
	state.Register(constants.GameOverStateKey, state.New(states.GameOverState))

	sfx.MusicPlayer.RegisterMusicTrack("assets/snakeoil.wav", "snakeoil")
	sfx.MusicPlayer.NewSet("snakeoil", []string{"snakeoil"}, sfx.Repeat, 0., 2.)
	sfx.SetMusicVolume(75)
	sfx.SetSoundVolume(75)

	sfx.SoundPlayer.RegisterSound("assets/buttonpress.wav", "buttonpress")
	sfx.SoundPlayer.RegisterSound("assets/pickup.wav", "pickup")
	sfx.SoundPlayer.RegisterSound("assets/place.wav", "place")
	sfx.SoundPlayer.RegisterSound("assets/place2.wav", "place2")

	testSheet, err := img.LoadSpriteSheet("assets/test1.json")
	if err != nil {
		panic(err)
	}
	img.AddBatcher(constants.TestBatch, testSheet, true, true)

	debug.Initialize(&viewport.MainCamera.PostCamPos)
	debug.Text = true

	data.SetTotalWareSize()

	win.SetColorMask(colornames.Black)
	win.Show()
	sfx.MusicPlayer.PlayMusic("snakeoil")
	timing.Reset()
	for !win.Closed() {
		timing.Update()
		debug.Clear()
		data.DebugInput.Update(win, viewport.MainCamera.Mat)

		options.WindowUpdate(win)
		if options.Updated {
			viewport.MainCamera.CamPos = pixel.V(viewport.MainCamera.Rect.W()*0.5, viewport.MainCamera.Rect.H()*0.5)
		}

		if data.DebugInput.Get("fullscreen").JustPressed() {
			options.FullScreen = !options.FullScreen
		}
		if data.DebugInput.Get("debugText").JustPressed() {
			debug.Text = !debug.Text
		}

		state.Update(win)
		viewport.MainCamera.Update()
		state.Draw(win)

		//win.SetSmooth(false)
		debug.Draw(win)
		//win.SetSmooth(true)
		//win.SetSmooth(options.BilinearFilter)

		sfx.MusicPlayer.Update()
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
