package game

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

type Game struct {
	sceneManager *SceneManager
	input        Input
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
	if g.sceneManager == nil {
		g.sceneManager = &SceneManager{}
		g.sceneManager.GoTo(&TitleScene{})
	}

	g.input.Update()
	if err := g.sceneManager.Update(&g.input); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
}
