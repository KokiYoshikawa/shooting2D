package main

import (
	"log"

	"github.com/KokiYoshikawa/gogame/shooting"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(shooting.ScreenWidth*2, shooting.ScreenHeight*2)
	ebiten.SetWindowTitle("Blocks (Ebitengine Demo)")
	if err := ebiten.RunGame(&shooting.Game{}); err != nil {
		log.Fatal(err)
	}
}
