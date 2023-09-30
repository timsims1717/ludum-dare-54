package states

import (
	"github.com/faiface/pixel/pixelgl"
	"ludum-dare-54/pkg/state"
)

type packingState struct {
	*state.AbstractState
}

func (s *packingState) Unload() {

}

func (s *packingState) Load() {

}

func (s *packingState) Update(win *pixelgl.Window) {

}

func (s *packingState) Draw(win *pixelgl.Window) {

}

func (s *packingState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
