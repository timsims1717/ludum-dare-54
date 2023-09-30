package systems

import (
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/object"
)

func DragSystem() {
	for _, result := range myecs.Manager.Query(myecs.IsDrag) {
		obj, okO := result.Components[myecs.Object].(*object.Object)
		dt, okD := result.Components[myecs.Drag].(*data.DragTimer)
		if okO {
			pos := data.GameInput.World
			pos = data.GameView.ProjectWorld(pos)
			pos = data.GameView.ConstrainR(pos, obj.Rect)
			obj.Pos = pos
			click := data.GameInput.Get("click")
			if okD {
				if dt.Quick {
					if click.JustPressed() {
						result.Entity.RemoveComponent(myecs.Drag)
					}
				} else {
					dt.Timer.Update()
					if dt.Timer.Done() {
						if !click.Pressed() {
							result.Entity.RemoveComponent(myecs.Drag)
						}
					} else {
						dt.Quick = !click.Pressed()
					}
				}
			} else if !click.Pressed() {
				result.Entity.RemoveComponent(myecs.Drag)
			}
		}
	}
}
