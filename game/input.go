package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Input manages the input state including gamepads and keyboards.
type Input struct {
	gamepadIDs []ebiten.GamepadID
}

func (i *Input) Update() {

}

// 特定のキーが押されているかをチェックする
func pressedKey(str ebiten.Key) bool {
	inputArray := inpututil.PressedKeys()
	for _, v := range inputArray {
		if v == str {
			return true
		}
	}
	return false
}

func (i *Input) IsShootingPressed() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}
