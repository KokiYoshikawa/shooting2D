package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TitleScene struct {
	count int
}

var imageTitle *ebiten.Image

func init() {
	var err error
	imageTitle, _, err = ebitenutil.NewImageFromFile("public/title.png")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *TitleScene) Update(state *GameState) error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	title_op := &ebiten.DrawImageOptions{}
	screen.DrawImage(imageTitle, title_op)
}
