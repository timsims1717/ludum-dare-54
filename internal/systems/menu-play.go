package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/sfx"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
	"ludum-dare-54/pkg/typeface"
)

// Play Sub Menus
var (
	SmartCarMenu   *typeface.Text
	MiniVanMenu    *typeface.Text
	CargoVanMenu   *typeface.Text
	SemiTruckMenu  *typeface.Text
	WagonMenu      *typeface.Text
	BackFromCar    *typeface.Text
	PickedTruck    *data.Truck
	PickedTruckKey constants.TruckTypes

	EasyMenu    *typeface.Text
	MediumMenu  *typeface.Text
	HardMenu    *typeface.Text
	BackFromDif *typeface.Text
	PickedDiff  constants.DifficultyType
)

func InitCarMenu() {
	if SmartCarMenu == nil {
		car := data.AvailableTrucks[constants.SmartCar]
		SmartCarMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		SmartCarMenu.Obj.Layer = 50
		SmartCarMenu.SetPos(pixel.V(-140, 200.))
		SmartCarMenu.SetColor(constants.BaseUIText)
		SmartCarMenu.SetText(car.TruckLabel)
		SmartCarMenu.Obj.SetRect(pixel.R(0, 0, 250, 100))
		SmartCarMenu.Obj.Rect = SmartCarMenu.Obj.Rect.Moved(pixel.V(0, -30))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, SmartCarMenu.Obj).
			AddComponent(myecs.Drawable, SmartCarMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					SmartCarMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						PickedTruck = car
						PickedTruckKey = constants.SmartCar
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					SmartCarMenu.SetColor(constants.BaseUIText)
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
				infoTxt.Obj.Hidden = SmartCarMenu.Obj.Hidden
				return false
			}))
	}
	if MiniVanMenu == nil {
		car := data.AvailableTrucks[constants.Minivan]
		MiniVanMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		MiniVanMenu.Obj.Layer = 50
		MiniVanMenu.SetPos(pixel.V(140, 200.))
		MiniVanMenu.SetColor(constants.BaseUIText)
		MiniVanMenu.SetText(car.TruckLabel)
		MiniVanMenu.Obj.SetRect(pixel.R(0, 0, 250, 100))
		MiniVanMenu.Obj.Rect = MiniVanMenu.Obj.Rect.Moved(pixel.V(0, -30))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, MiniVanMenu.Obj).
			AddComponent(myecs.Drawable, MiniVanMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					MiniVanMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						PickedTruck = car
						PickedTruckKey = constants.Minivan
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					MiniVanMenu.SetColor(constants.BaseUIText)
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
				infoTxt.Obj.Hidden = MiniVanMenu.Obj.Hidden
				return false
			}))
	}
	if CargoVanMenu == nil {
		car := data.AvailableTrucks[constants.CargoVan]
		CargoVanMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		CargoVanMenu.Obj.Layer = 50
		CargoVanMenu.SetPos(pixel.V(-140, 80.))
		CargoVanMenu.SetColor(constants.BaseUIText)
		CargoVanMenu.SetText(car.TruckLabel)
		CargoVanMenu.Obj.SetRect(pixel.R(0, 0, 250, 100))
		CargoVanMenu.Obj.Rect = CargoVanMenu.Obj.Rect.Moved(pixel.V(0, -30))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, CargoVanMenu.Obj).
			AddComponent(myecs.Drawable, CargoVanMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					CargoVanMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						PickedTruck = car
						PickedTruckKey = constants.CargoVan
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					CargoVanMenu.SetColor(constants.BaseUIText)
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
				infoTxt.Obj.Hidden = CargoVanMenu.Obj.Hidden
				return false
			}))
	}
	if SemiTruckMenu == nil {
		car := data.AvailableTrucks[constants.SemiTruck]
		SemiTruckMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		SemiTruckMenu.Obj.Layer = 50
		SemiTruckMenu.SetPos(pixel.V(140, 80.))
		SemiTruckMenu.SetColor(constants.BaseUIText)
		SemiTruckMenu.SetText(car.TruckLabel)
		SemiTruckMenu.Obj.SetRect(pixel.R(0, 0, 250, 100))
		SemiTruckMenu.Obj.Rect = SemiTruckMenu.Obj.Rect.Moved(pixel.V(0, -30))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, SemiTruckMenu.Obj).
			AddComponent(myecs.Drawable, SemiTruckMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					SemiTruckMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						PickedTruck = car
						PickedTruckKey = constants.SemiTruck
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					SemiTruckMenu.SetColor(constants.BaseUIText)
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
				infoTxt.Obj.Hidden = SemiTruckMenu.Obj.Hidden
				return false
			}))
	}
	if WagonMenu == nil {
		car := data.AvailableTrucks[constants.Wagon]
		WagonMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		WagonMenu.Obj.Layer = 50
		WagonMenu.SetPos(pixel.V(-140, -40))
		WagonMenu.SetColor(constants.BaseUIText)
		WagonMenu.SetText(car.TruckLabel)
		WagonMenu.Obj.SetRect(pixel.R(0, 0, 250, 100))
		WagonMenu.Obj.Rect = WagonMenu.Obj.Rect.Moved(pixel.V(0, -30))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, WagonMenu.Obj).
			AddComponent(myecs.Drawable, WagonMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					WagonMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						PickedTruck = car
						PickedTruckKey = constants.Wagon
						click.Consume()
						ShowDiffMenu()
					}
				} else {
					WagonMenu.SetColor(constants.BaseUIText)
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
				infoTxt.Obj.Hidden = WagonMenu.Obj.Hidden
				return false
			}))
	}
	if BackFromCar == nil {
		BackFromCar = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		BackFromCar.Obj.Layer = 50
		BackFromCar.SetPos(pixel.V(140, -40))
		BackFromCar.SetColor(constants.BaseUIText)
		BackFromCar.SetText("Back")
		BackFromCar.Obj.SetRect(pixel.R(0, 0, 250, 100))
		BackFromCar.Obj.Rect = BackFromCar.Obj.Rect.Moved(pixel.V(0, -30))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, BackFromCar.Obj).
			AddComponent(myecs.Drawable, BackFromCar).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					BackFromCar.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						ShowMainMenu()
						click.Consume()
					}
				} else {
					BackFromCar.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, BackFromCar)
	}
}

func ShowCarMenu() {
	HideAllMenus()
	SmartCarMenu.Obj.Hidden = false
	MiniVanMenu.Obj.Hidden = false
	CargoVanMenu.Obj.Hidden = false
	SemiTruckMenu.Obj.Hidden = false
	WagonMenu.Obj.Hidden = false
	BackFromCar.Obj.Hidden = false
}

func InitDifficultyMenu() {
	if EasyMenu == nil {
		difficulty := constants.DifficultyLevels[constants.Easy]
		EasyMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		EasyMenu.Obj.Layer = 50
		EasyMenu.SetPos(pixel.V(-200, 200.))
		EasyMenu.SetColor(constants.BaseUIText)
		EasyMenu.SetText(difficulty.Label)
		EasyMenu.Obj.SetRect(pixel.R(0, 0, 350, 150))
		EasyMenu.Obj.Rect = EasyMenu.Obj.Rect.Moved(pixel.V(0, -55))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, EasyMenu.Obj).
			AddComponent(myecs.Drawable, EasyMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					EasyMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						PickedDiff = constants.Easy
						click.Consume()
						StartGame()
					}
				} else {
					EasyMenu.SetColor(constants.BaseUIText)
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
				infoTxt.Obj.Hidden = EasyMenu.Obj.Hidden
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
		MediumMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		MediumMenu.Obj.Layer = 50
		MediumMenu.SetPos(pixel.V(200, 200.))
		MediumMenu.SetColor(constants.BaseUIText)
		MediumMenu.SetText(difficulty.Label)
		MediumMenu.Obj.SetRect(pixel.R(0, 0, 350, 150))
		MediumMenu.Obj.Rect = MediumMenu.Obj.Rect.Moved(pixel.V(0, -55))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, MediumMenu.Obj).
			AddComponent(myecs.Drawable, MediumMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					MediumMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						PickedDiff = constants.Medium
						click.Consume()
						StartGame()
					}
				} else {
					MediumMenu.SetColor(constants.BaseUIText)
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
				infoTxt.Obj.Hidden = MediumMenu.Obj.Hidden
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
		HardMenu = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		HardMenu.Obj.Layer = 50
		HardMenu.SetPos(pixel.V(-200, 30.))
		HardMenu.SetColor(constants.BaseUIText)
		HardMenu.SetText(difficulty.Label)
		HardMenu.Obj.SetRect(pixel.R(0, 0, 350, 150))
		HardMenu.Obj.Rect = HardMenu.Obj.Rect.Moved(pixel.V(0, -55))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, HardMenu.Obj).
			AddComponent(myecs.Drawable, HardMenu).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					HardMenu.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						PickedDiff = constants.Hard
						click.Consume()
						StartGame()
					}
				} else {
					HardMenu.SetColor(constants.BaseUIText)
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
				infoTxt.Obj.Hidden = HardMenu.Obj.Hidden
				return false
			}))
		HardInfo = func() {
			infoTxt.SetText(fmt.Sprintf(`Allowed Missed Deliveries: %d
Allowed Abandoned Wares: %d
Minimum Wares: %d`, difficulty.NumberofMissedDeliveries-1, difficulty.NumberofAbandonedWares-1, data.TargetWares(PickedTruck, difficulty)))
		}
	}
	if BackFromDif == nil {
		BackFromDif = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.21, 0., 0.)
		BackFromDif.Obj.Layer = 50
		BackFromDif.SetPos(pixel.V(200, 30.))
		BackFromDif.SetColor(constants.BaseUIText)
		BackFromDif.SetText("Back")
		BackFromDif.Obj.SetRect(pixel.R(0, 0, 350, 150))
		BackFromDif.Obj.Rect = BackFromDif.Obj.Rect.Moved(pixel.V(0, -55))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, BackFromDif.Obj).
			AddComponent(myecs.Drawable, BackFromDif).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover && !data.Starting {
					BackFromDif.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						sfx.SoundPlayer.PlaySound("buttonpress", 0.)
						ShowCarMenu()
						click.Consume()
					}
				} else {
					BackFromDif.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, BackFromDif)
	}
}

func ShowDiffMenu() {
	HideAllMenus()
	EasyMenu.Obj.Hidden = false
	MediumMenu.Obj.Hidden = false
	HardMenu.Obj.Hidden = false
	BackFromDif.Obj.Hidden = false
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
	data.Abandon = false
	data.CurrentTruck = nil
	data.FirstLoad = true
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
