package states

import (
	"github.com/faiface/pixel/pixelgl"
	"ludum-dare-54/pkg/state"
)

type gameState struct {
	*state.AbstractState
}

func (s *gameState) Unload() {

}

func (s *gameState) Load() {

}

func (s *gameState) Update(win *pixelgl.Window) {

}

func (s *gameState) Draw(win *pixelgl.Window) {

}

func (s *gameState) SetAbstract(aState *state.AbstractState) {
	s.AbstractState = aState
}
