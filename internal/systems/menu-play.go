package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/typeface"
)

// Play Sub Menus
var (
	SmartCarMenu  *data.MenuItem
	MiniVanMenu   *data.MenuItem
	CargoVanMenu  *data.MenuItem
	SemiTruckMenu *data.MenuItem
	WagonMenu     *data.MenuItem
	BackFromCar   *data.MenuItem
	PickedCar     *data.Truck

	EasyMenu    *data.MenuItem
	MediumMenu  *data.MenuItem
	HardMenu    *data.MenuItem
	BackFromDif *data.MenuItem
	PickedDiff  *constants.Difficulty
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
				if hvc.Hover {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowDiffMenu()
						PickedCar = car
						click.Consume()
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
				if hvc.Hover {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowDiffMenu()
						PickedCar = car
						click.Consume()
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
				if hvc.Hover {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowDiffMenu()
						PickedCar = car
						click.Consume()
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
				if hvc.Hover {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowDiffMenu()
						PickedCar = car
						click.Consume()
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
				if hvc.Hover {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowDiffMenu()
						PickedCar = car
						click.Consume()
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
				if hvc.Hover {
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
		txt.SetPos(pixel.V(0, 200.))
		txt.SetColor(constants.BaseUIText)
		txt.SetText(difficulty.Label)
		txt.Obj.SetRect(pixel.R(0, 0, 250, 100))
		txt.Obj.Rect = txt.Obj.Rect.Moved(pixel.V(0, -30))
		EasyMenu = &data.MenuItem{
			Text: txt,
			Func: nil,
		}
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, txt.Obj).
			AddComponent(myecs.Drawable, txt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					txt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						HideAllMenus()
						PickedDiff = difficulty
						click.Consume()
					}
				} else {
					txt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, EasyMenu)

		infoTxt := typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.12, 0., 0.)
		infoTxt.Obj.Layer = 50
		infoTxt.SetPos(pixel.V(0, 155.))
		infoTxt.SetColor(constants.BaseUIText)
		infoTxt.SetText(fmt.Sprintf(`Allowed Missed Deliveries: %d
Allowed Abandoned Wares: %d
Minimum Wares: %d`, difficulty.NumberofMissedDeliveries-1, difficulty.NumberofAbandonedWares-1, 2))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, infoTxt.Obj).
			AddComponent(myecs.Drawable, infoTxt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewFrameFunc(func() bool {
				infoTxt.Obj.Hidden = txt.Obj.Hidden
				return false
			}))
	}
}

func ShowDiffMenu() {
	HideAllMenus()
	EasyMenu.Text.Obj.Hidden = false
	MediumMenu.Text.Obj.Hidden = false
	HardMenu.Text.Obj.Hidden = false
	BackFromDif.Text.Obj.Hidden = false
}
