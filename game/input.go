package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Input manages the input state including gamepads and keyboards.
type Input struct {
	gamepadIDs []ebiten.GamepadID
}

func (i *Input) Update() {

}
