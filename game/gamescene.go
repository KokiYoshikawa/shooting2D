package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameScene struct {
	input    Input
	gameover bool
}

type Chara struct {
	life     int
	center   Center
	size     Size
	image    *ebiten.Image
	moveX    int
	moveY    int
	moveMaxX int
	moveMaxY int
	bullets  [3]bullets
}

type bullets struct {
	life   bool
	center Center
	size   Size
	image  *ebiten.Image
	moveX  int
	moveY  int
}
type Center struct {
	x int
	y int
}

type Size struct {
	width  int
	height int
}

var (
	row, col        int
	play_area_width int
	my_chara        Chara
	enemys          [enemyNum]Chara
	left_image      *ebiten.Image
	right_image     *ebiten.Image
	gameover_image  *ebiten.Image
)

func NewGameScene() *GameScene {
	return &GameScene{}
}

func init() {
	var err error
	play_area_width = ScreenWidth - playAreaLeft

	left_image, _, err = ebitenutil.NewImageFromFile("public/left_area.png")
	if err != nil {
		log.Fatal(err)
	}

	right_image, _, err = ebitenutil.NewImageFromFile("public/right_area.png")
	if err != nil {
		log.Fatal(err)
	}

	gameover_image, _, err = ebitenutil.NewImageFromFile("public/gameover.png")
	if err != nil {
		log.Fatal(err)
	}

	//画像を読み込む
	my_chara.image, _, err = ebitenutil.NewImageFromFile("public/test.png")
	if err != nil {
		log.Fatal(err)
	}
	// 画像のサイズを取得
	my_chara.size.width, my_chara.size.height = my_chara.image.Size()
	col, row = 0, 0

	//自機
	initMyCharaAgrrangement()

	// 敵
	initEnemyArrangement()
}

func (g *GameScene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (s *GameScene) Update(state *GameState) error {
	if s.gameover {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			initMyCharaAgrrangement()
			initEnemyArrangement()
			state.SceneManager.GoTo(&TitleScene{})
		}
		return nil
	}

	if pressedKey(ebiten.KeyArrowRight) {
		// 右移動
		row += 5
	}
	if pressedKey(ebiten.KeyArrowLeft) {
		// 左移動
		row -= 5
	}
	if pressedKey(ebiten.KeyArrowDown) {
		// 下移動
		col += 5
	}
	if pressedKey(ebiten.KeyArrowUp) {
		// 上移動
		col -= 5
	}

	// 自機が撃つ
	if state.Input.IsShootingPressed() {
		for i := 0; i < 3; i++ {
			if !my_chara.bullets[i].life {
				my_chara.bullets[i].life = true
				my_chara.bullets[i].center.x = my_chara.center.x
				my_chara.bullets[i].center.y = my_chara.center.y
				break
			}
		}
	}

	// 自弾の位置更新
	for i := 0; i < 3; i++ {
		if my_chara.bullets[i].life {
			my_chara.bullets[i].center.x += my_chara.bullets[i].moveX
			my_chara.bullets[i].center.y += my_chara.bullets[i].moveY
		}

		// 画面外に出た弾を消去
		if my_chara.bullets[i].center.y < 0 {
			my_chara.bullets[i].life = false
		}

		// 自弾当たり判定
		for y := 0; y < enemyNum; y++ {
			if enemys[y].life > 0 {
				if collision(enemys[y].center, my_chara.bullets[i].center, enemys[y].size, my_chara.bullets[i].size) {
					my_chara.bullets[i].life = false
					my_chara.bullets[i].center.x = my_chara.center.x
					my_chara.bullets[i].center.y = my_chara.center.y
					enemys[y].center.x = 480
					enemys[y].center.y = 640
					enemys[y].life -= 1

					break
				}
			}
		}
	}

	// 敵の位置更新
	for i := 0; i < enemyNum; i++ {
		if enemys[i].life > 0 {
			enemys[i].moveMaxX += enemys[i].moveX
			if enemys[i].moveMaxX >= 20 {
				enemys[i].moveX *= -1
				enemys[i].moveMaxX = 20
			}
			if enemys[i].moveMaxX <= -20 {
				enemys[i].moveX *= -1
				enemys[i].moveMaxX = -20
			}

			enemys[i].center.x += enemys[i].moveX
			enemys[i].center.y += enemys[i].moveY
		}
	}

	s.gameover = checkEnemyNum()

	return nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	//画面左
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(0.0), float64(0.0))
	screen.DrawImage(left_image, op)

	//画面右
	op.GeoM.Translate(float64(playAreaRight), float64(0.0))
	screen.DrawImage(right_image, op)

	// 位置
	my_chara.center.x = ScreenWidth/2 + row
	my_chara.center.y = ScreenHeight - my_chara.size.height/2 + col

	// 自機が右端に当たった時
	if my_chara.center.x > playAreaRight-my_chara.size.width/2 {
		my_chara.center.x = playAreaRight - my_chara.size.width/2
		row -= 5
	}
	// 自機が左端に当たった時
	if my_chara.center.x < playAreaLeft+my_chara.size.width/2 {
		my_chara.center.x = playAreaLeft + my_chara.size.width/2
		row += 5
	}
	// 自機が下端に当たった時
	if my_chara.center.y > ScreenHeight-my_chara.size.height/2 {
		my_chara.center.y = ScreenHeight - my_chara.size.height/2
		col -= 5
	}
	// 自機が上端に当たった時
	if my_chara.center.y < my_chara.size.height/2 {
		my_chara.center.y = my_chara.size.height / 2
		col += 5
	}

	for i := 0; i < 3; i++ {
		if my_chara.bullets[i].life {
			mb_op := &ebiten.DrawImageOptions{}
			mb_op.GeoM.Translate(-float64(my_chara.bullets[i].size.width)/2, -float64(my_chara.bullets[i].size.height)/2)
			mb_op.GeoM.Translate(float64(my_chara.bullets[i].center.x), float64(my_chara.bullets[i].center.y))
			screen.DrawImage(my_chara.bullets[i].image, mb_op)
		}
	}

	for i := 0; i < enemyNum; i++ {
		if enemys[i].life != 0 {
			e_op := &ebiten.DrawImageOptions{}
			e_op.GeoM.Translate(-float64(enemys[i].size.width)/2, -float64(enemys[i].size.height)/2)
			e_op.GeoM.Translate(float64(enemys[i].center.x), float64(enemys[i].center.y))
			screen.DrawImage(enemys[i].image, e_op)
		}
	}

	// 画像のオプションの準備
	op = &ebiten.DrawImageOptions{}

	// 画像の中心をスクリーンの左上に移動させる
	// 原点が画面の左上だからである
	op.GeoM.Translate(-float64(my_chara.size.width)/2, -float64(my_chara.size.height)/2)

	// 画像をスクリーンの中心に持ってくる
	op.GeoM.Translate(float64(my_chara.center.x), float64(my_chara.center.y))

	// 画像を描画する
	screen.DrawImage(my_chara.image, op)

	if g.gameover {
		liWidth, _ := left_image.Size()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(liWidth), float64(0.0))
		screen.DrawImage(gameover_image, op)
	}
}

func checkEnemyNum() bool {
	for i := 0; i < enemyNum; i++ {
		if enemys[i].life != 0 {
			return false
		}
	}
	return true
}

func initMyCharaAgrrangement() {
	var err error
	// 自弾の読み込み
	for i := 0; i < 3; i++ {
		my_chara.bullets[i].image, _, err = ebitenutil.NewImageFromFile("public/bullet.png")
		if err != nil {
			log.Fatal(err)
		}
		my_chara.bullets[i].life = false
		my_chara.bullets[i].size.width, my_chara.bullets[i].size.height = my_chara.bullets[i].image.Size()
		my_chara.bullets[i].moveX, my_chara.bullets[i].moveY = 0, -3
	}
}

func initEnemyArrangement() {
	var err error
	for i := 0; i < enemyNum; i++ {
		enemys[i].image, _, err = ebitenutil.NewImageFromFile("public/test.png")
		if err != nil {
			log.Fatal(err)
		}
		enemys[i].life = 1
		enemys[i].center.x = (ScreenWidth/2 - 100) + (i * 100)
		enemys[i].center.y = 100
		enemys[i].size.width, enemys[i].size.height = enemys[i].image.Size()
		enemys[i].moveX, enemys[i].moveY = 1, 0
		enemys[i].moveMaxX, enemys[i].moveMaxY = 0, 0

		for n := 0; n < 3; n++ {
			enemys[i].bullets[n].image, _, err = ebitenutil.NewImageFromFile("public/bullet.png")
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// 当たり判定
func collision(center_a Center, center_b Center, size_a Size, size_b Size) bool {
	distance_x := center_a.x - center_b.x
	distance_y := center_a.y - center_b.y

	if distance_x < 0 {
		distance_x *= -1
	}

	if distance_y < 0 {
		distance_y *= -1
	}

	sum_width := (size_a.width + size_b.width) / 2
	sum_height := (size_a.height + size_b.height) / 2

	if distance_x < sum_width && distance_y < sum_height {
		return true
	}

	return false
}
