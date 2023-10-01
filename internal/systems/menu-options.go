package systems

import (
	"fmt"
	"github.com/faiface/pixel"
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/options"
	"ludum-dare-54/pkg/typeface"
)

var (
	MusicVolume *typeface.Text
	MVol        = 75
	AudioVolume *typeface.Text
	AVol        = 75
	Vsync       *typeface.Text
	Fullscreen  *typeface.Text
	BackFromOpt *typeface.Text
)

func InitOptionsMenu() {
	if MusicVolume == nil {
		MusicVolume = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		MusicVolume.Obj.Layer = 50
		MusicVolume.SetPos(pixel.V(0, 215))
		MusicVolume.SetColor(constants.BaseUIText)
		MusicVolume.SetText("Music Volume: 75")
		MusicVolume.Obj.SetRect(pixel.R(0, 0, 680, 100))
		MusicVolume.Obj.Rect = MusicVolume.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, MusicVolume.Obj).
			AddComponent(myecs.Drawable, MusicVolume).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					MusicVolume.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					rClick := hvc.Input.Get("rightClick")
					if click.JustReleased() {
						MVol += 25
						if MVol > 100 {
							MVol = 0
						}
						MusicVolume.SetText(fmt.Sprintf("Music Volume: %d", MVol))
						click.Consume()
					} else if rClick.JustReleased() {
						MVol -= 25
						if MVol < 0 {
							MVol = 100
						}
						MusicVolume.SetText(fmt.Sprintf("Music Volume: %d", MVol))
						rClick.Consume()
					}
				} else {
					MusicVolume.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, MusicVolume)
	}
	if AudioVolume == nil {
		AudioVolume = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		AudioVolume.Obj.Layer = 50
		AudioVolume.SetPos(pixel.V(0, 100))
		AudioVolume.SetColor(constants.BaseUIText)
		AudioVolume.SetText("Audio Volume: 75")
		AudioVolume.Obj.SetRect(pixel.R(0, 0, 680, 100))
		AudioVolume.Obj.Rect = AudioVolume.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, AudioVolume.Obj).
			AddComponent(myecs.Drawable, AudioVolume).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					AudioVolume.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					rClick := hvc.Input.Get("rightClick")
					if click.JustReleased() {
						AVol += 25
						if AVol > 100 {
							AVol = 0
						}
						AudioVolume.SetText(fmt.Sprintf("Audio Volume: %d", AVol))
						click.Consume()
					} else if rClick.JustReleased() {
						AVol -= 25
						if AVol < 0 {
							AVol = 100
						}
						AudioVolume.SetText(fmt.Sprintf("Audio Volume: %d", AVol))
						rClick.Consume()
					}
				} else {
					AudioVolume.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, AudioVolume)
	}
	if Fullscreen == nil {
		Fullscreen = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		Fullscreen.Obj.Layer = 50
		Fullscreen.SetPos(pixel.V(0, -15))
		Fullscreen.SetColor(constants.BaseUIText)
		Fullscreen.SetText("Fullscreen: Off")
		Fullscreen.Obj.SetRect(pixel.R(0, 0, 570, 100))
		Fullscreen.Obj.Rect = Fullscreen.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, Fullscreen.Obj).
			AddComponent(myecs.Drawable, Fullscreen).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					Fullscreen.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						options.FullScreen = !options.FullScreen
						if options.FullScreen {
							Fullscreen.SetText("Fullscreen: On")
						} else {
							Fullscreen.SetText("Fullscreen: Off")
						}
					}
				} else {
					Fullscreen.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, Fullscreen)
	}
	if Vsync == nil {
		Vsync = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		Vsync.Obj.Layer = 50
		Vsync.SetPos(pixel.V(0, -130))
		Vsync.SetColor(constants.BaseUIText)
		Vsync.SetText("VSync: On")
		Vsync.Obj.SetRect(pixel.R(0, 0, 370, 100))
		Vsync.Obj.Rect = Vsync.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, Vsync.Obj).
			AddComponent(myecs.Drawable, Vsync).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					Vsync.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						options.VSync = !options.VSync
						if options.VSync {
							Vsync.SetText("VSync: On")
						} else {
							Vsync.SetText("VSync: Off")
						}
					}
				} else {
					Vsync.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, Vsync)
	}
	if BackFromOpt == nil {
		BackFromOpt = typeface.New("main", typeface.NewAlign(typeface.Center, typeface.Center), 1.2, 0.4, 0, 0.)
		BackFromOpt.Obj.Layer = 50
		BackFromOpt.SetPos(pixel.V(0, -245))
		BackFromOpt.SetColor(constants.BaseUIText)
		BackFromOpt.SetText("Back")
		BackFromOpt.Obj.SetRect(pixel.R(0, 0, 300, 100))
		BackFromOpt.Obj.Rect = BackFromOpt.Obj.Rect.Moved(pixel.V(0, -26))
		myecs.Manager.NewEntity().
			AddComponent(myecs.Object, BackFromOpt.Obj).
			AddComponent(myecs.Drawable, BackFromOpt).
			AddComponent(myecs.DrawTarget, data.MenuView).
			AddComponent(myecs.Update, data.NewHoverClickFn(data.GameInput, data.MenuView, func(hvc *data.HoverClick) {
				if hvc.Hover {
					BackFromOpt.SetColor(constants.HoverUIText)
					click := hvc.Input.Get("click")
					if click.JustReleased() {
						ShowMainMenu()
					}
				} else {
					BackFromOpt.SetColor(constants.BaseUIText)
				}
			}))
		MenuItems = append(MenuItems, BackFromOpt)
	}
}

func ShowOptionsMenu() {
	HideAllMenus()
	MusicVolume.Obj.Hidden = false
	AudioVolume.Obj.Hidden = false
	Fullscreen.Obj.Hidden = false
	Vsync.Obj.Hidden = false
	BackFromOpt.Obj.Hidden = false
}
