package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	s_height = 480
	s_width  = 680
)

type Vector struct {
	X, Y float64
}
type Ball struct {
	x        float64
	y        float64
	Velocity Vector
	image    *ebiten.Image
}

type Player struct {
	x      float64
	y      float64
	score  int
	paddle *ebiten.Image
}

func (p *Player) MoveDown() {
	p.y = p.y + 5
	if p.y >= s_height-40 {
		p.y = s_height - 40
	}
}

func (p *Player) MoveUp() {
	p.y = p.y - 5
	if p.y <= 0 {
		p.y = 0
	}
}

func (b *Ball) Update(game *Pong) {
	b.x += b.Velocity.X
	b.y += b.Velocity.Y
	if b.y <= 0 {
		b.y = 0
		b.Velocity.Y *= -1
	}
	if b.y >= s_height-16 {
		b.y = s_height - 16
		b.Velocity.Y *= -1
	}

	if b.x <= 25 && (game.player1.y < b.y && b.y < game.player1.y+40+16) {
		b.Velocity.X *= -1.2
		b.x = 25
	}

	if b.x >= s_width-45 && (game.player2.y < b.y && b.y < game.player2.y+40+16) {
		b.Velocity.X *= -1.2
		b.x = s_width - 45
	}
	if b.x <= 0 {
		b.Velocity.X = 0
		b.Velocity.Y = 0
		b.x = s_width/2 - 8
		b.y = s_height/2 - 8
		game.player1.score += 1
	}
	if b.x >= s_width-16 {
		b.Velocity.X = 0
		b.Velocity.Y = 0
		b.x = s_width/2 - 8
		b.y = s_height/2 - 8
		game.player2.score += 1
	}
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
		x:        s_width/2 - 8,
		y:        s_height/2 - 8,
		Velocity: Vector{X: 2, Y: 3},
		image:    ballImage,
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

func (p *Pong) keyPressed() {
	var keys []ebiten.Key
	for _, key := range inpututil.AppendPressedKeys(keys) {
		switch key {
		case ebiten.KeyS:
			p.player1.MoveDown()
		case ebiten.KeyW:
			p.player1.MoveUp()
		case ebiten.KeyArrowDown:
			p.player2.MoveDown()
		case ebiten.KeyArrowUp:
			p.player2.MoveUp()
		case ebiten.KeySpace:
			if p.ball.Velocity.X == 0 && p.ball.Velocity.Y == 0 {
				p.ball.Velocity.X = 2
				p.ball.Velocity.Y = 3
			}
		}
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	// Draw ball
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(b.image, options)
}

func (p *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(p.x), float64(p.y))
	screen.DrawImage(p.paddle, options)
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (p *Pong) Update() error {
	// Write your game's logical update.
	p.keyPressed()
	p.ball.Update(p)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (p *Pong) Draw(screen *ebiten.Image) {
	p.ball.Draw(screen)
	p.player1.Draw(screen)
	p.player2.Draw(screen)
	scoreString := fmt.Sprintf("Player 1: %x  -  Player 2: %x", p.player1.score, p.player2.score)
	ebitenutil.DebugPrintAt(screen, scoreString, 20, 20)
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
