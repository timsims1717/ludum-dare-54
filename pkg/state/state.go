package state

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image/color"
)

type State interface {
	Unload(*pixelgl.Window)
	Load(*pixelgl.Window)
	Update(*pixelgl.Window)
	Draw(*pixelgl.Window)
	SetAbstract(*AbstractState)
}

type AbstractState struct {
	State
	LoadPrc  float64
	ShowLoad bool
	loaded   bool
}

func New(state State) *AbstractState {
	aState := &AbstractState{
		State: state,
	}
	state.SetAbstract(aState)
	return aState
}

var (
	popState   = false
	pushState  = false
	nextState  = "unknown"
	stateStack []string
	stackPtr   = -1
	states     = map[string]*AbstractState{}
	defState   string

	loading = false
	done    = make(chan struct{})

	clearColor color.Color
)

func init() {
	clearColor = colornames.Black
}

func Register(key string, state *AbstractState) {
	if _, ok := states[key]; ok {
		fmt.Printf("Warning: state.Register - state '%s' already registered\n", key)
	} else if state == nil {
		panic(fmt.Sprintf("state.Register - state %s is nil\n", key))
	} else {
		if defState == "" {
			defState = key
			nextState = key
		}
		states[key] = state
	}
}

func SetClearColor(col color.Color) {
	clearColor = col
}

func Update(win *pixelgl.Window) {
	updateState(win)
	if loading {
		select {
		case <-done:
			loading = false
		default:
			if LoadingScreen != nil {
				LoadingScreen.Update(win)
			}
		}
	} else {
		if len(stateStack) > 0 && stackPtr > -1 {
			if cState, ok := states[stateStack[stackPtr]]; ok {
				cState.Update(win)
			} else {
				panic(fmt.Sprintf("state.Update - state %s doesn't exist\n", stateStack[stackPtr]))
			}
		}
	}
}

func Draw(win *pixelgl.Window) {
	win.Clear(clearColor)
	if loading {
		if lState, ok2 := states[stateStack[stackPtr]]; ok2 && lState.ShowLoad && LoadingScreen != nil {
			LoadingScreen.Draw(win)
			return
		}
	} else {
		for _, state := range stateStack {
			cState, ok1 := states[state]
			if !ok1 {
				panic(fmt.Sprintf("state.Draw - state %s doesn't exist\n", state))
			} else {
				cState.Draw(win)
			}
		}
	}
}

func updateState(win *pixelgl.Window) {
	if !loading {
		if popState {
			// states need to be popped
			if len(stateStack) == 0 || stackPtr == -1 {
				panic("state.Pop - tried to pop with an empty state stack")
			} else {
				stackPtr--
			}
			for si := len(stateStack) - 1; si > stackPtr; si-- {
				if cState, ok := states[stateStack[si]]; ok {
					// unload
					cState.Unload(win)
				}
			}
			if stackPtr == -1 {
				stateStack = []string{}
			} else {
				stateStack = stateStack[:stackPtr+1]
			}
			popState = false
		}
		if pushState {
			if cState, ok := states[nextState]; ok {
				stateStack = append(stateStack, nextState)
				stackPtr++
				go func() {
					// initialize
					cState.Load(win)
					done <- struct{}{}
				}()
				loading = true
			}
			pushState = false
		}
	}
	if len(stateStack) == 0 || stackPtr == -1 {
		PushState(defState)
	}
}

func SwitchState(s string) {
	PopState()
	PushState(s)
}

func PushState(s string) {
	if !pushState {
		pushState = true
		nextState = s
	} else {
		panic(fmt.Sprintf("state.PushState - tried to push state %s when a push is already happening\n", s))
	}
}

func PopState() {
	if !popState {
		popState = true
	} else {
		panic(fmt.Sprintf("state.PopState - tried to pop a state when a pop is already happening\n"))
	}
}
