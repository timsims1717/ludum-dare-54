package systems

import (
	"ludum-dare-54/internal/data"
	"ludum-dare-54/internal/myecs"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/object"
)

func QueueSystem() {
	for i := 0; i < 8; i++ {
		if data.ItemQueue[i] == nil {
			found := false
			for j := i + 1; j < 8; j++ {
				if data.ItemQueue[j] != nil {
					item := data.ItemQueue[j]
					ip := object.NewInterpolation(object.InterpolateY).
						AddGween(item.Object.Pos.Y, rightQueueY(i), 0.4, ease.OutCubic)
					item.Entity.AddComponent(myecs.Interpolation, ip)
					data.ItemQueue[i], data.ItemQueue[j] = data.ItemQueue[j], nil
					found = true
				}
			}
			if !found {
				// create new item at [i]
			}
		}
	}
}

var (
	bottomSlot = -40.
	slotSize   = 64.
)

func rightQueueY(i int) float64 {
	return -bottomSlot + (float64(i) * slotSize)
}
