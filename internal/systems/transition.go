package systems

import (
	"ludum-dare-54/internal/constants"
	"ludum-dare-54/internal/data"
	gween "ludum-dare-54/pkg/gween64"
	"ludum-dare-54/pkg/gween64/ease"
	"ludum-dare-54/pkg/state"
	"ludum-dare-54/pkg/timing"
)

func LeaveTransitionSystem() {
	if data.LeaveTransition {
		if data.FadeTween == nil {
			data.FadeTween = gween.New(255., 0, 1, ease.Linear)
		}
		if data.TransitionTimer == nil {
			data.TransitionTimer = timing.New(1)
		}
		if data.TransitionTimer.UpdateDone() {
			state.SwitchState(constants.PackingStateKey)
		}
	}
}
