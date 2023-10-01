package states

var (
	PackingStateKey    = "packing_state"
	TransitionStateKey = "transition_state"

	PackingState    = &packingState{}
	TransitionState = &transitionState{}
)
