package game

import (
	ldtk_utils "github.com/Soockee/clockwork-kingdom/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/rs/zerolog/log"
	"github.com/solarlune/ldtkgo"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

type Game struct {
	LDTKProject    *ldtkgo.Project
	EbitenRenderer *ldtk_utils.EbitenRenderer
	CurrentLevel   int
	ActiveLayers   []bool
	BGImage        *ebiten.Image
}

func NewGame() *Game {
	g := &Game{
		ActiveLayers: []bool{true, true},
	}

	var err error

	// g.LDTKProject, err = ldtkgo.Open("assets/clockwork-kingdom.ldtk")

	if err != nil {
		panic(err)
	}

	// g.EbitenRenderer = ldtk_utils.NewEbitenRenderer(ldtk_utils.NewDiskLoader("assets/"))
	//
	// g.RenderLevel()
	a := 123123
	log.Debug().Msgf("%d", a)
	return g
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
	// screen.Fill(g.LDTKProject.Levels[g.CurrentLevel].BGColor) // We want to use the BG Color when possible

	// if g.BGImage != nil {
	// 	opt := &ebiten.DrawImageOptions{}
	// 	bgImage := g.LDTKProject.Levels[g.CurrentLevel].BGImage
	// 	opt.GeoM.Translate(-bgImage.CropRect[0], -bgImage.CropRect[1])
	// 	opt.GeoM.Scale(bgImage.ScaleX, bgImage.ScaleY)
	// 	screen.DrawImage(g.BGImage, opt)
	// }
	//
	// for i, layer := range g.EbitenRenderer.RenderedLayers {
	// 	if g.ActiveLayers[i] {
	// 		screen.DrawImage(layer.Image, &ebiten.DrawImageOptions{})
	// 	}
	// }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) RenderLevel() {

	if g.CurrentLevel >= len(g.LDTKProject.Levels) {
		g.CurrentLevel = 0
	}

	if g.CurrentLevel < 0 {
		g.CurrentLevel = len(g.LDTKProject.Levels) - 1
	}

	level := g.LDTKProject.Levels[g.CurrentLevel]

	if level.BGImage != nil {
		g.BGImage, _, _ = ebitenutil.NewImageFromFile(level.BGImage.Path)
	} else {
		g.BGImage = nil
	}

	g.EbitenRenderer.Render(level)
}

func (g *Game) IsRegularTermination(err error) bool {
	_, ok := err.(regularTermination)
	return ok
}

// IDK, needed?
type regularTermination struct{}

func (regularTermination) Error() string {
	return "regular termination"
}
