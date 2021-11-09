package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
}

// Game implements ebiten.Game interface.
type Pong struct {
	ball *Ball
}

func (b *Ball) Draw(screen *ebiten.Image) {
	// Draw ball
	image := ebiten.NewImage(16, 16)
	ebitenutil.DrawRect(image, 0, 0, 16, 16, color.White)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(160, 120)
	screen.DrawImage(image, options)
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
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (p *Pong) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	pong := &Pong{
		ball: &Ball{},
	}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Ghetto Pong")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(pong); err != nil {
		log.Fatal(err)
	}
}
