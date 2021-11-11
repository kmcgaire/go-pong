package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	s_height = 480
	s_width  = 680
)

type Ball struct {
	x     int
	y     int
	image *ebiten.Image
}

type Player struct {
	x      int
	y      int
	paddle *ebiten.Image
}

// Game implements ebiten.Game interface.
type Pong struct {
	ball    *Ball
	player1 *Player
	player2 *Player
}

func NewPong() *Pong {
	paddle := ebiten.NewImage(15, 40)
	paddle.Fill(color.White)

	ballImage := ebiten.NewImage(16, 16)
	ballImage.Fill(color.White)

	ball := &Ball{
		x:     s_width/2 - 8,
		y:     s_height/2 - 8,
		image: ballImage,
	}

	player1 := &Player{
		x:      10,
		y:      50,
		paddle: paddle,
	}
	player2 := &Player{
		x:      s_width - 30,
		y:      100,
		paddle: paddle,
	}
	pong := &Pong{
		ball:    ball,
		player1: player1,
		player2: player2,
	}
	return pong
}

func (b *Ball) Draw(screen *ebiten.Image) {
	// Draw ball
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.image, options)
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	log.Printf("p.x %v p.y %v", p.x, p.y)
	options.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.paddle, options)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (p *Pong) Update() error {
	// Write your game's logical update.
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (p *Pong) Draw(screen *ebiten.Image) {
	p.ball.Draw(screen)
	p.player1.Draw(screen)
	p.player2.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (p *Pong) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	pong := NewPong()

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(s_width, s_height)
	ebiten.SetWindowTitle("Ghetto Pong")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(pong); err != nil {
		log.Fatal(err)
	}
}
