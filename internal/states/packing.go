package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/internal/systems"
	"ludum-dare-54/pkg/debug"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/img"
	"ludum-dare-54/pkg/object"
	"ludum-dare-54/pkg/options"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/viewport"
	"ludum-dare-54/pkg/world"
	"math"
)

type packingState struct {
	*state.AbstractState
}

func (s *packingState) Unload() {
	data.FadeTween = nil
	data.LeavePacking = false
	data.LeaveStep = 0
	data.WareQueue = [8]*data.Ware{}
	for _, ware := range data.SellWares {
		myecs.Manager.DisposeEntity(ware.Entity)
	}
	data.SellWares = []*data.Ware{}
	for _, result := range myecs.Manager.Query(myecs.IsWare) {
		_, okO := result.Components[myecs.Object].(*object.Object)
		ware, okW := result.Components[myecs.Ware].(*data.Ware)
		if okO && okW {
			if ware.TIndex < 0 {
				myecs.Manager.DisposeEntity(result)
			}
		}
	}
}

func (s *packingState) Load() {
	data.FadeTween = gween.New(0., 255, 0.4, ease.Linear)
	if data.GameView == nil {
		data.GameView = viewport.New(nil)
		data.GameView.SetRect(pixel.R(0, 0, 640, 360))
	}
	if data.ScoreView == nil {
		data.ScoreView = viewport.New(nil)
		data.ScoreView.CamPos = pixel.V(0, data.ScoreView.Rect.H()*0.5)
	}
	if data.FirstLoad {
		data.IsTimer = false
		systems.CreateTruck()
		data.GameView.CamPos = pixel.ZV
		data.NewScore()
		data.SetDifficulty(constants.Easy)
		data.BottomDrop = pixel.R(-240, -130, 340, -40)
		data.LeftDrop = pixel.R(-240, -130, -40, 60)
		systems.ScoreboardInit()
	} else {
		systems.ScoreboardReset()
		data.DepartureTimer = timing.New(float64(data.CurrentDifficulty.TimeToSell))
		systems.SellInit()
	}
	data.GameView.CamPos.X += (float64(data.CurrentTruck.Width)-1)*0.5*world.TileSize - (40)
	data.GameView.CamPos.Y += (math.Min(float64(data.CurrentTruck.Height), 3) - 1) * 0.5 * world.TileSize
	s.UpdateViews()
}

func (s *packingState) Update(win *pixelgl.Window) {
	if data.FadeTween != nil {
		c, done := data.FadeTween.Update(timing.DT)
		viewport.MainCamera.Mask.R = uint8(c)
		viewport.MainCamera.Mask.G = uint8(c)
		viewport.MainCamera.Mask.B = uint8(c)
		if done {
			data.FadeTween = nil
		}
	}

	debug.AddText("Packing State")
	data.TimerCount.Obj.Hidden = !data.IsTimer
	data.PercCount.Obj.Hidden = data.IsTimer
	data.RightPercentFull.Obj.Hidden = !data.IsTimer
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
		data.LeavePacking = true
	}

	systems.DragSystem()
	systems.FunctionSystem()
	// custom systems
	systems.LeavePackingSystem()
	systems.QueueSystem()
	systems.ScoreSystem()
	// object systems
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.GameView.Update()
	data.ScoreView.Update()

	systems.TrunkClean()

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))
}

func (s *packingState) Draw(win *pixelgl.Window) {
	data.GameView.Canvas.Clear(colornames.Green)
	systems.DrawSystem(win, 0)
	for i := 1; i <= data.CurrentTruck.Height; i++ {
		systems.DrawSystem(win, i)
	}
	systems.DrawSystem(win, 15)
	systems.DrawSystem(win, 20)
	img.Batchers[constants.TestBatch].Draw(data.GameView.Canvas)
	img.Clear()

	data.GameView.Draw(win)

	data.ScoreView.Canvas.Clear(color.RGBA{})
	systems.DrawSystem(win, 29)
	img.Batchers[constants.TestBatch].Draw(data.ScoreView.Canvas)
	systems.DrawSystem(win, 30)
	data.ScoreView.Draw(win)
	img.Clear()

	systems.TemporarySystem()
	if options.Updated {
		s.UpdateViews()
	}
}

func (s *packingState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}

func (s *packingState) UpdateViews() {
	data.GameView.PortPos = viewport.MainCamera.PostCamPos
	data.GameView.PortSize.X = viewport.MainCamera.Rect.W() / data.GameView.Rect.W()
	data.GameView.PortSize.Y = data.GameView.PortSize.X

	svw := viewport.MainCamera.Rect.W() * 0.3
	svh := viewport.MainCamera.Rect.H() * 0.4
	data.ScoreView.SetRect(pixel.R(0, 0, svw, svh))
	data.ScoreView.SetZoom(viewport.MainCamera.Rect.W() / 1600)
	data.ScoreView.PortPos = viewport.MainCamera.PostCamPos
	data.ScoreView.PortPos.Y += (viewport.MainCamera.Rect.H()-data.ScoreView.Rect.H())*0.5 - 20.
	data.ScoreView.PortPos.X -= viewport.MainCamera.Rect.W() * 0.28
}
