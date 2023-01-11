package shooting

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	title_op := &ebiten.DrawImageOptions{}
	screen.DrawImage(imageTitle, title_op)
}
