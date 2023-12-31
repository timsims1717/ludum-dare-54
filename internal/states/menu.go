package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/internal/systems"
	"ludum-dare-54/pkg/debug"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/options"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/viewport"
)

type mainMenuState struct {
	*state.AbstractState
}

func (s *mainMenuState) Unload(win *pixelgl.Window) {
	data.FadeTween = nil
	data.LeaveTransition = false
	data.TransitionTimer = nil
	data.TransitionStep = 0
}

func (s *mainMenuState) Load(win *pixelgl.Window) {
	data.Starting = false
	data.FadeTween = gween.New(0., 255, 0.4, ease.Linear)
	if data.MenuView == nil {
		data.MenuView = viewport.New(nil)
		data.MenuView.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W(), viewport.MainCamera.Rect.H()))
	}
	if data.GameView == nil {
		data.GameView = viewport.New(nil)
		data.GameView.SetRect(pixel.R(0, 0, 640, 360))
	}
	data.MenuView.CamPos = pixel.ZV
	if data.IMDraw == nil {
		data.IMDraw = imdraw.New(nil)
	}
	systems.InitMenuItems(win)
	systems.ShowMainMenu()
	s.UpdateViews()
	systems.NewBackground()
}

func (s *mainMenuState) Update(win *pixelgl.Window) {
	if data.FadeTween != nil {
		c, done := data.FadeTween.Update(timing.DT)
		viewport.MainCamera.Mask.R = uint8(c)
		viewport.MainCamera.Mask.G = uint8(c)
		viewport.MainCamera.Mask.B = uint8(c)
		if done {
			data.FadeTween = nil
		}
	}

	debug.AddText("Main Menu State")
	data.GameInput.Update(win, viewport.MainCamera.Mat)
	debug.AddIntCoords("World", int(data.GameInput.World.X), int(data.GameInput.World.Y))
	inPos := data.MenuView.ProjectWorld(data.GameInput.World)
	debug.AddIntCoords("MenuView World", int(inPos.X), int(inPos.Y))

	if data.DebugInput.Get("debugSP").JustPressed() {
		data.MenuView.ZoomIn(1.)
	} else if data.DebugInput.Get("debugSM").JustPressed() {
		data.MenuView.ZoomIn(-1.)
	}
	if data.DebugInput.Get("camUp").Pressed() {
		data.MenuView.PortPos.Y += 100. * timing.DT
	} else if data.DebugInput.Get("camDown").Pressed() {
		data.MenuView.PortPos.Y -= 100. * timing.DT
	}
	if data.DebugInput.Get("camRight").Pressed() {
		data.MenuView.PortPos.X += 100. * timing.DT
	} else if data.DebugInput.Get("camLeft").Pressed() {
		data.MenuView.PortPos.X -= 100. * timing.DT
	}

	systems.FunctionSystem()
	// custom systems
	systems.StartSystem()
	// object systems
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.GameView.Update()
	data.MenuView.Update()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *mainMenuState) Draw(win *pixelgl.Window) {
	data.GameView.Canvas.Clear(constants.PackingColor)
	systems.DrawSystem(win, -2)
	img.Batchers[constants.TestBatch].Draw(data.GameView.Canvas)
	img.Clear()
	data.GameView.Draw(win)

	data.MenuView.Canvas.Clear(color.RGBA{})
	systems.DrawMenuBG()
	systems.DrawSystem(win, 50)
	data.MenuView.Draw(win)

	img.Clear()

	systems.TemporarySystem()
	if options.Updated {
		s.UpdateViews()
	}
}

func (s *mainMenuState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *mainMenuState) UpdateViews() {
	data.GameView.PortPos = viewport.MainCamera.PostCamPos
	data.GameView.PortSize.X = viewport.MainCamera.Rect.W() / data.GameView.Rect.W()
	data.GameView.PortSize.Y = data.GameView.PortSize.X

	data.MenuView.PortPos = viewport.MainCamera.PostCamPos
	data.MenuView.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W(), viewport.MainCamera.Rect.H()))
	data.MenuView.CamPos = pixel.ZV
}
