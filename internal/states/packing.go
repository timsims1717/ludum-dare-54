package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/internal/systems"
	"ludum-dare-54/pkg/debug"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/state"
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
	viewport.MainCamera.CamPos = pixel.ZV
	viewport.MainCamera.CamPos.X += w * 0.5 * world.TileSize
	viewport.MainCamera.CamPos.Y += d * 0.5 * world.TileSize
}

func (s *packingState) Update(win *pixelgl.Window) {
	debug.AddText("Packing State")
	data.GameInput.Update(win, viewport.MainCamera.Mat)
	debug.AddIntCoords("World", int(data.GameInput.World.X), int(data.GameInput.World.Y))

	if data.DebugInput.Get("debugSP").JustPressed() {
		viewport.MainCamera.ZoomIn(1.)
	} else if data.DebugInput.Get("debugSM").JustPressed() {
		viewport.MainCamera.ZoomIn(-1.)
	}

	systems.FunctionSystem()
	// more systems
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	systems.TemporarySystem()
	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *packingState) Draw(win *pixelgl.Window) {
	systems.DrawSystem(win, 1)
	img.Batchers[data.TestBatch].Draw(viewport.MainCamera.Canvas)
}

func (s *packingState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
