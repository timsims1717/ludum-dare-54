package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/internal/systems"
	"ludum-dare-54/pkg/debug"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/options"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/viewport"
)

type pauseState struct {
	*state.AbstractState
}

func (s *pauseState) Unload(win *pixelgl.Window) {
	data.LeaveTransition = false
	data.Paused = false
}

func (s *pauseState) Load(win *pixelgl.Window) {
	data.LeaveTransition = false
	data.Paused = true
	data.Starting = false
	data.MenuView.CamPos = pixel.ZV
	if data.IMDraw == nil {
		data.IMDraw = imdraw.New(nil)
	}
	systems.ShowPauseMenu()
	s.UpdateViews()
}

func (s *pauseState) Update(win *pixelgl.Window) {
	debug.AddText("Pause State")
	data.GameInput.Update(win, viewport.MainCamera.Mat)
	debug.AddIntCoords("World", int(data.GameInput.World.X), int(data.GameInput.World.Y))
	inPos := data.MenuView.ProjectWorld(data.GameInput.World)
	debug.AddIntCoords("MenuView World", int(inPos.X), int(inPos.Y))

	if data.GameInput.Get("pause").JustPressed() {
		state.PopState()
		data.GameInput.Get("pause").Consume()
	} else {
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
		// object systems
		systems.InterpolationSystem()
		systems.ParentSystem()
		systems.ObjectSystem()
	}
	data.MenuView.Update()
	data.ScoreView.Update()
	data.GameView.Update()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *pauseState) Draw(win *pixelgl.Window) {
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

func (s *pauseState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *pauseState) UpdateViews() {
	data.MenuView.PortPos = viewport.MainCamera.PostCamPos
	data.MenuView.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W(), viewport.MainCamera.Rect.H()))
	data.MenuView.CamPos = pixel.ZV
}
