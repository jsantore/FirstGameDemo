package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/colornames"
	_ "image/png"
	"log"
	"math/rand"
)

const(
	ScreenWidth = 800
	ScreenHeight = 700
)

type Sprite struct{
	pict *ebiten.Image
	xLoc 	int
	yLoc 	int
	dx 		int
	dy 		int
}

type Game struct{
	playerSprite Sprite
	coinSprite Sprite
	drawOps ebiten.DrawImageOptions
	collectedGold bool
}

func gotGold(player, gold Sprite) bool{
	goldWidth, goldHeight := gold.pict.Size()
	playerWidth, playerHeight := player.pict.Size()
	if player.xLoc < gold.xLoc+goldWidth &&
		player.xLoc+playerWidth > gold.xLoc &&
		player.yLoc < gold.yLoc+goldHeight &&
		player.yLoc+playerHeight > gold.yLoc{
		return true
	}
	return false
}

func (game *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft){
		game.playerSprite.dx=-3
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight){
		game.playerSprite.dx = 3
	} else if inpututil.IsKeyJustReleased(ebiten.KeyRight) || inpututil.IsKeyJustReleased(ebiten.KeyLeft){
		game.playerSprite.dx = 0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp){
		game.playerSprite.dy= -3
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown){
		game.playerSprite.dy = 3
	} else if inpututil.IsKeyJustReleased(ebiten.KeyUp) || inpututil.IsKeyJustReleased(ebiten.KeyDown){
		game.playerSprite.dy = 0
	}
	game.playerSprite.yLoc += game.playerSprite.dy
	game.playerSprite.xLoc += game.playerSprite.dx
	if game.collectedGold ==false{
		game.collectedGold = gotGold(game.playerSprite, game.coinSprite)
	}
	return nil
}

func (game Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Mediumaquamarine)
	game.drawOps.GeoM.Reset()
	game.drawOps.GeoM.Translate(float64(game.playerSprite.xLoc), float64(game.playerSprite.yLoc))
	screen.DrawImage(game.playerSprite.pict, &game.drawOps)
	if !game.collectedGold {
		game.drawOps.GeoM.Reset()
		game.drawOps.GeoM.Translate(float64(game.coinSprite.xLoc), float64(game.coinSprite.yLoc))
		screen.DrawImage(game.coinSprite.pict, &game.drawOps)
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main(){
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Comp510 First Graphics")
	gameObject := Game{}
	loadImage(&gameObject)
	gameObject.playerSprite.yLoc = ScreenHeight/2
	width, height:= gameObject.coinSprite.pict.Size()
	//rand.Seed(982982439842398)
	gameObject.coinSprite.xLoc = rand.Intn(ScreenWidth-width)
	gameObject.coinSprite.yLoc = rand.Intn(ScreenHeight-height)
	if err:= ebiten.RunGame(&gameObject); err!= nil{
		log.Fatal("Oh no! something terrible happened", err)
	}
}

func loadImage(game *Game) {
	pict, _, err:= ebitenutil.NewImageFromFile("galleon.png")
	if err != nil{
		log.Fatal("failed to load image", err)
	}
	game.playerSprite.pict = pict
	coins, _, err:= ebitenutil.NewImageFromFile("gold-coins-large.png")
	if err != nil{
		log.Fatal("failed to load image", err)
	}
	game.coinSprite.pict = coins
}
