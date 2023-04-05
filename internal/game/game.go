package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	// Add your game-specific structs, variables, and assets here.
}

type regularTermination struct{}

func (regularTermination) Error() string {
	return "regular termination"
}

func (g *Game) Update() error {
	// Handle input.
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return regularTermination{}
	}

	// Update game state.
	// Update your game-specific logic here.

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw game assets.
	// Draw your game-specific assets here.
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) IsRegularTermination(err error) bool {
	_, ok := err.(regularTermination)
	return ok
}
