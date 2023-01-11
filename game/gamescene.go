package shooting

import (
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth   = 800 // プレイ範囲は480
	ScreenHeight  = 640
	playAreaLeft  = 160
	playAreaRight = 640
	enemyNum      = 3
)

type GameScene struct {
	input Input
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *GameScene) Update() error {
	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {

}
