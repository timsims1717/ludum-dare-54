package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/internal/systems"
	"ludum-dare-54/pkg/debug"
	"ludum-dare-54/pkg/options"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/viewport"
)

type gameoverState struct {
	*state.AbstractState
}

func (s *gameoverState) Unload(win *pixelgl.Window) {
	data.FadeTween = nil
	data.LeaveTransition = false
	data.TransitionTimer = nil
	data.TransitionStep = 0
	systems.HideAllMenus()
	data.LeavePacking = false
}

func (s *gameoverState) Load(win *pixelgl.Window) {
	data.Starting = false
	if data.GameView == nil {
		data.GameView = viewport.New(nil)
		data.GameView.SetRect(pixel.R(0, 0, 640, 360))
	}
	systems.HideBigMessage()
	systems.ShowGameOverMenu()
	s.UpdateViews()
	data.LeavePacking = false
	systems.UpdateGameOverStats()
}

func (s *gameoverState) Update(win *pixelgl.Window) {
	if data.FadeTween != nil {
		c, done := data.FadeTween.Update(timing.DT)
		viewport.MainCamera.Mask.R = uint8(c)
		viewport.MainCamera.Mask.G = uint8(c)
		viewport.MainCamera.Mask.B = uint8(c)
		if done {
			data.FadeTween = nil
		}
	}

	debug.AddText("Game Over State")
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
	systems.LeaveGameOverSystem()
	// object systems
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.MenuView.Update()
	data.GameView.Update()
	data.ScoreView.Update()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *gameoverState) Draw(win *pixelgl.Window) {
	data.MenuView.Canvas.Clear(color.RGBA{})
	systems.DrawMenuBG()
	systems.DrawSystem(win, 50)
	data.MenuView.Draw(win)

	systems.TemporarySystem()
	if options.Updated {
		s.UpdateViews()
	}
}

func (s *gameoverState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *gameoverState) UpdateViews() {
	data.GameView.PortPos = viewport.MainCamera.PostCamPos
	data.GameView.PortSize.X = viewport.MainCamera.Rect.W() / data.GameView.Rect.W()
	data.GameView.PortSize.Y = viewport.MainCamera.Rect.H() / data.GameView.Rect.H()

	data.MenuView.PortPos = viewport.MainCamera.PostCamPos
	data.MenuView.SetRect(pixel.R(0, 0, viewport.MainCamera.Rect.W(), viewport.MainCamera.Rect.H()))
	data.MenuView.CamPos = pixel.ZV
}
