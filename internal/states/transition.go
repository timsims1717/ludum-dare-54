package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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

type transitionState struct {
	*state.AbstractState
}

func (s *transitionState) Unload(win *pixelgl.Window) {
	data.FadeTween = nil
	data.LeaveTransition = false
	data.TransitionTimer = nil
	data.TransitionStep = 0
}

func (s *transitionState) Load(win *pixelgl.Window) {
	data.Starting = false
	data.FadeTween = gween.New(0., 255, 0.4, ease.Linear)
	if data.GameView == nil {
		data.GameView = viewport.New(nil)
		data.GameView.SetRect(pixel.R(0, 0, 640, 360))
	}
	data.GameView.CamPos = pixel.ZV
	s.UpdateViews()
	data.LeaveTransition = false
	systems.AddNewStall()
	systems.CreateMiniTruck()
}

func (s *transitionState) Update(win *pixelgl.Window) {
	if data.FadeTween != nil {
		c, done := data.FadeTween.Update(timing.DT)
		viewport.MainCamera.Mask.R = uint8(c)
		viewport.MainCamera.Mask.G = uint8(c)
		viewport.MainCamera.Mask.B = uint8(c)
		if done {
			data.FadeTween = nil
		}
	}

	debug.AddText("Transition State")
	data.GameInput.Update(win, viewport.MainCamera.Mat)
	debug.AddIntCoords("World", int(data.GameInput.World.X), int(data.GameInput.World.Y))
	inPos := data.GameView.ProjectWorld(data.GameInput.World)
	debug.AddIntCoords("GameView World", int(inPos.X), int(inPos.Y))

	if data.DebugInput.Get("debugSP").JustPressed() {
		data.ScoreView.ZoomIn(1.)
	} else if data.DebugInput.Get("debugSM").JustPressed() {
		data.ScoreView.ZoomIn(-1.)
	}
	if data.DebugInput.Get("camUp").Pressed() {
		data.GameView.CamPos.Y += 100. * timing.DT
	} else if data.DebugInput.Get("camDown").Pressed() {
		data.GameView.CamPos.Y -= 100. * timing.DT
	}
	if data.DebugInput.Get("camRight").Pressed() {
		data.GameView.CamPos.X += 100. * timing.DT
	} else if data.DebugInput.Get("camLeft").Pressed() {
		data.GameView.CamPos.X -= 100. * timing.DT
	}
	if data.DebugInput.Get("leave").Pressed() {
		data.LeaveTransition = true
	}

	systems.FunctionSystem()
	// custom systems
	systems.TransitionSystem()
	// object systems
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.GameView.CamPos = data.MiniTruckObj.Pos

	data.GameView.Update()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *transitionState) Draw(win *pixelgl.Window) {
	data.GameView.Canvas.Clear(constants.PackingColor)

	systems.DrawPaths()
	systems.DrawSystem(win, 39)
	systems.DrawSystem(win, 40)
	img.Batchers[constants.TestBatch].Draw(data.GameView.Canvas)
	data.GameView.Draw(win)

	img.Clear()
	systems.TemporarySystem()
	if options.Updated {
		s.UpdateViews()
	}
}

func (s *transitionState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *transitionState) UpdateViews() {
	data.GameView.PortPos = viewport.MainCamera.PostCamPos
	data.GameView.PortSize.X = viewport.MainCamera.Rect.W() / data.GameView.Rect.W()
	data.GameView.PortSize.Y = viewport.MainCamera.Rect.H() / data.GameView.Rect.H()
}
