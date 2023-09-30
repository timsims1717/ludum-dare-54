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
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/viewport"
	"ludum-dare-54/pkg/world"
)

type packingState struct {
	*state.AbstractState
}

func (s *packingState) Unload() {

}

func (s *packingState) Load() {
	w := 5.
	d := 3.
	h := 3.
	systems.CreateTruck(w, d, h)
	data.GameView = viewport.New(nil)
	data.GameView.SetRect(pixel.R(0, 0, 640, 360))
	data.GameView.PortPos = viewport.MainCamera.PostCamPos
	data.GameView.CamPos = pixel.ZV
	data.GameView.CamPos.X += (w - 1) * 0.5 * world.TileSize
	data.GameView.CamPos.Y += (d - 1) * 0.5 * world.TileSize
	UpdateViews()
}

func (s *packingState) Update(win *pixelgl.Window) {
	debug.AddText("Packing State")
	data.GameInput.Update(win, viewport.MainCamera.Mat)
	debug.AddIntCoords("World", int(data.GameInput.World.X), int(data.GameInput.World.Y))
	inPos := data.GameView.ProjectWorld(data.GameInput.World)
	debug.AddIntCoords("GameView World", int(inPos.X), int(inPos.Y))

	if data.DebugInput.Get("debugSP").JustPressed() {
		data.GameView.ZoomIn(1.)
	} else if data.DebugInput.Get("debugSM").JustPressed() {
		data.GameView.ZoomIn(-1.)
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
	// more systems
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.GameView.Update()

	systems.TemporarySystem()
	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *packingState) Draw(win *pixelgl.Window) {
	data.GameView.Canvas.Clear(colornames.Green)
	systems.DrawSystem(win, 1)
	img.Batchers[data.TestBatch].Draw(data.GameView.Canvas)
	img.Clear()
	data.GameView.Canvas.Draw(win, data.GameView.Mat)
}

func (s *packingState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func UpdateViews() {
	data.GameView.PortSize.X = viewport.MainCamera.Rect.W() / data.GameView.Rect.W()
	data.GameView.PortSize.Y = viewport.MainCamera.Rect.H() / data.GameView.Rect.H()
}
