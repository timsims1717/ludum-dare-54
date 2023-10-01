package states

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/internal/systems"
	"ludum-dare-54/pkg/debug"
	"ludum-dare-54/pkg/img"
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

}

func (s *packingState) Load() {
	systems.CreateTruck()
	if data.GameView == nil {
		data.GameView = viewport.New(nil)
		data.GameView.SetRect(pixel.R(0, 0, 640, 360))
	}
	data.GameView.CamPos = pixel.ZV
	data.GameView.CamPos.X += (float64(data.CurrentTruck.Width) - 1) * 0.5 * world.TileSize
	data.GameView.CamPos.Y += (math.Min(float64(data.CurrentTruck.Height), 3) - 1) * 0.5 * world.TileSize
	data.ScoreView = viewport.New(nil)
	data.ScoreView.CamPos = pixel.V(0, data.ScoreView.Rect.H()*0.5)

	data.NewScore()
	data.SetDifficulty(constants.Easy)
	data.BottomDrop = pixel.R(-200, -130, 340, -40)
	data.LeftDrop = pixel.R(-200, -130, -40, 190)
	s.UpdateViews()
	systems.ScoreboardInit()
}

func (s *packingState) Update(win *pixelgl.Window) {
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

	systems.DragSystem()
	systems.FunctionSystem()
	// custom systems
	systems.QueueSystem()
	// object systems
	systems.InterpolationSystem()
	systems.ParentSystem()
	systems.ObjectSystem()

	data.PercCount.SetText(fmt.Sprintf("%d%% Full", data.CurrentTruck.PercentFilled))
	data.GameView.Update()
	data.ScoreView.Update()

	systems.TrunkClean()
	data.CheckForFailure()
	if !data.IsTimer && data.CurrentTruck.PercentFilled >= data.CurrentDifficulty.InitialTrunkTargetFill {
		systems.StartTimer()
	}
	if data.IsTimer {
		systems.UpdateTimer()
		data.RightPercentFull.SetText(fmt.Sprintf("%d%% Full", data.CurrentTruck.PercentFilled))
		if data.CurrentTruck.PercentFilled >= data.CurrentDifficulty.InitialTrunkTargetFill {
			data.RightPercentFull.SetColor(pixel.ToRGBA(colornames.Green))
			//TODO: Unlock Button
		} else {
			data.RightPercentFull.SetColor(pixel.ToRGBA(colornames.Red))
			//TODO: Lock Button
		}
	}
	data.RightLoadedWares.SetText(fmt.Sprintf("Loaded Wares: %d", len(data.CurrentTruck.Wares)))
	data.RightLoadHeight.SetText(fmt.Sprintf("Load Height: %d / %d", 0, data.CurrentTruck.Depth))
	//data.LeftTitle.SetText(fmt.Sprintf("DELIVERIES\n%d Complete\n%d Missed\n$%d.00", data.CurrentScore.SuccessfulDeliveries,
	//data.CurrentScore.MissedDeliveries, data.CurrentScore.Cash))
	data.LeftCompletes.SetText(fmt.Sprintf("%d Complete", data.CurrentScore.SuccessfulDeliveries))
	data.LeftAbandoned.SetText(fmt.Sprintf("%d Missed", data.CurrentScore.MissedDeliveries))
	data.LeftAbandoned.SetText(fmt.Sprintf("%d Abandonded", data.CurrentScore.AbandonedWares))
	data.LeftCash.SetText(fmt.Sprintf("$%d", data.CurrentScore.Cash))

	myecs.UpdateManager()
	debug.AddText(fmt.Sprintf("Entity Count: %d", myecs.FullCount))

	if data.CurrentTruck.PercentFilled >= data.CurrentDifficulty.InitialTrunkTargetFill && data.DepartureTimer.Done() {
		//AutoLeave
	}
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

	data.GameView.Canvas.Draw(win, data.GameView.Mat)

	data.ScoreView.Canvas.Clear(colornames.White)
	systems.DrawSystem(win, 29)
	img.Batchers[constants.TestBatch].Draw(data.ScoreView.Canvas)
	systems.DrawSystem(win, 30)
	data.ScoreView.Canvas.Draw(win, data.ScoreView.Mat)

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
	data.GameView.PortSize.Y = viewport.MainCamera.Rect.H() / data.GameView.Rect.H()

	svw := math.Max(viewport.MainCamera.Rect.W()*0.2, 330)
	svh := math.Max(viewport.MainCamera.Rect.H()*0.4, 330)
	data.ScoreView.SetRect(pixel.R(0, 0, svw, svh))
	data.ScoreView.SetZoom(viewport.MainCamera.Rect.W() / 1600)
	data.ScoreView.PortPos = viewport.MainCamera.PostCamPos
	data.ScoreView.PortPos.Y += (viewport.MainCamera.Rect.H()-data.ScoreView.Rect.H())*0.5 - 20.
	data.ScoreView.PortPos.X -= viewport.MainCamera.Rect.W() * 0.28
}
