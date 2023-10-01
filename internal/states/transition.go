package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/internal/systems"
	"ludum-dare-54/pkg/debug"
	"ludum-dare-54/pkg/options"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/viewport"
)

type transitionState struct {
	*state.AbstractState
}

func (s *transitionState) Unload() {

}

func (s *transitionState) Load() {
	if data.GameView == nil {
		data.GameView = viewport.New(nil)
		data.GameView.SetRect(pixel.R(0, 0, 640, 360))
	}
	data.GameView.CamPos = pixel.ZV
}

func (s *transitionState) Update(win *pixelgl.Window) {
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

	systems.FunctionSystem()
	// custom systems

	// object systems
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.GameView.Update()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *transitionState) Draw(win *pixelgl.Window) {
	data.GameView.Canvas.Clear(colornames.Red)

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
