package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/typeface"
)

// Play Sub Menus
var (
	SmartCarMenu   *data.MenuItem
	MiniVanMenu    *data.MenuItem
	CargoVanMenu   *data.MenuItem
	SemiTruckMenu  *data.MenuItem
	WagonMenu      *data.MenuItem
	BackFromCar    *data.MenuItem
	PickedTruck    *data.Truck
	PickedTruckKey constants.TruckTypes

	EasyMenu    *data.MenuItem
	MediumMenu  *data.MenuItem
	HardMenu    *data.MenuItem
	BackFromDif *data.MenuItem
	PickedDiff  constants.DifficultyType
)

func InitCarMenu() {
	if SmartCarMenu == nil {
		car := data.AvailableTrucks[constants.SmartCar]
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(-140, 200.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(car.TruckLabel)
		txt.Obj.SetRect(pixel.R(0, 0, 250, 100))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -30))
		SmartCarMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						PickedTruck = car
						PickedTruckKey = constants.SmartCar
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, SmartCarMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(-140, 155.))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText(fmt.Sprintf("Trunk Size: %dx%dx%d", car.Width, car.Depth, car.Height))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
	}
	if MiniVanMenu == nil {
		car := data.AvailableTrucks[constants.Minivan]
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(140, 200.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(car.TruckLabel)
		txt.Obj.SetRect(pixel.R(0, 0, 250, 100))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -30))
		MiniVanMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						PickedTruck = car
						PickedTruckKey = constants.Minivan
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MiniVanMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(140, 155))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText(fmt.Sprintf("Trunk Size: %dx%dx%d", car.Width, car.Depth, car.Height))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
	}
	if CargoVanMenu == nil {
		car := data.AvailableTrucks[constants.CargoVan]
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(-140, 80.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(car.TruckLabel)
		txt.Obj.SetRect(pixel.R(0, 0, 250, 100))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -30))
		CargoVanMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						PickedTruck = car
						PickedTruckKey = constants.CargoVan
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, CargoVanMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(-140, 35.))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText(fmt.Sprintf("Trunk Size: %dx%dx%d", car.Width, car.Depth, car.Height))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
	}
	if SemiTruckMenu == nil {
		car := data.AvailableTrucks[constants.SemiTruck]
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(140, 80.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(car.TruckLabel)
		txt.Obj.SetRect(pixel.R(0, 0, 250, 100))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -30))
		SemiTruckMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						PickedTruck = car
						PickedTruckKey = constants.SemiTruck
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, SemiTruckMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(140, 35.))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText(fmt.Sprintf("Trunk Size: %dx%dx%d", car.Width, car.Depth, car.Height))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
	}
	if WagonMenu == nil {
		car := data.AvailableTrucks[constants.Wagon]
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(-140, -40))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(car.TruckLabel)
		txt.Obj.SetRect(pixel.R(0, 0, 250, 100))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -30))
		WagonMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						PickedTruck = car
						PickedTruckKey = constants.Wagon
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, WagonMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(-140, -75))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText(fmt.Sprintf("Trunk Size: %dx%dx%d", car.Width, car.Depth, car.Height))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
	}
	if BackFromCar == nil {
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(140, -40))
		txt.SetColor(constants.BaseUIText)
		txt.SetText("Back")
		txt.Obj.SetRect(pixel.R(0, 0, 250, 100))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -30))
		BackFromCar = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowMainMenu()
						click.Consume()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, BackFromCar)
	}
}

func ShowCarMenu() {
	HideAllMenus()
	SmartCarMenu.Text.Obj.Hidden = false
	MiniVanMenu.Text.Obj.Hidden = false
	CargoVanMenu.Text.Obj.Hidden = false
	SemiTruckMenu.Text.Obj.Hidden = false
	WagonMenu.Text.Obj.Hidden = false
	BackFromCar.Text.Obj.Hidden = false
}

func InitDifficultyMenu() {
	if EasyMenu == nil {
		difficulty := constants.DifficultyLevels[constants.Easy]
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(-200, 200.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(difficulty.Label)
		txt.Obj.SetRect(pixel.R(0, 0, 350, 150))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -55))
		EasyMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						PickedDiff = constants.Easy
						click.Consume()
						StartGame()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, EasyMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(-200, 130.))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText("oops")
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
		EasyInfo = func() {
			infoTxt.SetText(fmt.Sprintf(`Allowed Missed Deliveries: %d
Allowed Abandoned Wares: %d
Minimum Wares: %d`, difficulty.NumberofMissedDeliveries-1, difficulty.NumberofAbandonedWares-1, data.TargetWares(PickedTruck, difficulty)))
		}
	}
	if MediumMenu == nil {
		difficulty := constants.DifficultyLevels[constants.Medium]
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(200, 200.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(difficulty.Label)
		txt.Obj.SetRect(pixel.R(0, 0, 350, 150))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -55))
		MediumMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						PickedDiff = constants.Medium
						click.Consume()
						StartGame()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MediumMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(200, 130.))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText("oops")
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
		MediumInfo = func() {
			infoTxt.SetText(fmt.Sprintf(`Allowed Missed Deliveries: %d
Allowed Abandoned Wares: %d
Minimum Wares: %d`, difficulty.NumberofMissedDeliveries-1, difficulty.NumberofAbandonedWares-1, data.TargetWares(PickedTruck, difficulty)))
		}
	}
	if HardMenu == nil {
		difficulty := constants.DifficultyLevels[constants.Hard]
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(-200, 30.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(difficulty.Label)
		txt.Obj.SetRect(pixel.R(0, 0, 350, 150))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -55))
		HardMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						PickedDiff = constants.Hard
						click.Consume()
						StartGame()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, HardMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(-200, -40))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText("oops")
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
		HardInfo = func() {
			infoTxt.SetText(fmt.Sprintf(`Allowed Missed Deliveries: %d
Allowed Abandoned Wares: %d
Minimum Wares: %d`, difficulty.NumberofMissedDeliveries-1, difficulty.NumberofAbandonedWares-1, data.TargetWares(PickedTruck, difficulty)))
		}
	}
	if BackFromDif == nil {
		txt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		txt.Obj.Layer = 50
		txt.SetPos(pixel.V(200, 30.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText("Back")
		txt.Obj.SetRect(pixel.R(0, 0, 350, 150))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -55))
		BackFromDif = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowCarMenu()
						click.Consume()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, BackFromDif)
	}
}

func ShowDiffMenu() {
	HideAllMenus()
	EasyMenu.Text.Obj.Hidden = false
	MediumMenu.Text.Obj.Hidden = false
	HardMenu.Text.Obj.Hidden = false
	BackFromDif.Text.Obj.Hidden = false
	EasyInfo()
	MediumInfo()
	HardInfo()
}

var (
	EasyInfo   func()
	MediumInfo func()
	HardInfo   func()
)

func StartGame() {
	data.CurrentTruck = nil
	PickedTruck = nil
	data.PickedTruckKey = PickedTruckKey
	data.PickedDiffKey = PickedDiff
	data.Starting = true
	if data.FadeTween == nil {
		data.FadeTween = gween.New(255., 0, 1, ease.Linear)
	}
	if data.TransitionTimer == nil {
		data.TransitionTimer = timing.New(1)
	}
}

func StartSystem() {
	if data.Starting {
		if data.TransitionTimer.UpdateDone() {
			HideAllMenus()
			state.SwitchState(constants.PackingStateKey)
			data.FirstLoad = true
		}
	}
}
