package myecs

import (
	"github.com/bytearena/ecs"
	"ludum-dare-54/pkg/object"
)

var (
	FullCount   = 0
	IDCount     = 0
	LoadedCount = 0
)

type ClearFlag bool

var (
	Manager = ecs.NewManager()

	Object = Manager.NewComponent()
	Parent = Manager.NewComponent()
	Temp   = Manager.NewComponent()
	Update = Manager.NewComponent()

	Interpolation = Manager.NewComponent()

	Drawable   = Manager.NewComponent()
	Animated   = Manager.NewComponent()
	DrawTarget = Manager.NewComponent()

	// Tags
	IsObject         = ecs.BuildTag(Object)
	IsTemp           = ecs.BuildTag(Temp, Object)
	HasParent        = ecs.BuildTag(Object, Parent)
	IsDrawable       = ecs.BuildTag(Object, Drawable)
	HasAnimation     = ecs.BuildTag(Animated, Object)
	HasUpdate        = ecs.BuildTag(Update)
	HasInterpolation = ecs.BuildTag(Object, Interpolation)
)

func UpdateManager() {
	LoadedCount = 0
	IDCount = 0
	FullCount = 0
	for _, result := range Manager.Query(IsObject) {
		if t, ok := result.Components[Object].(*object.Object); ok {
			FullCount++
			if t.ID != "" {
				IDCount++
				if t.Loaded {
					LoadedCount++
				}
			}
		}
	}
}
