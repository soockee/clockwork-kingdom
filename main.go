package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/Soockee/clockwork-kingdom/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	ebiten.SetWindowSize(game.ScreenWidth*2, game.ScreenHeight*2)
	ebiten.SetWindowTitle("Clockwork Kingdom")

	game := game.NewGame()
	if err := ebiten.RunGame(game); err != nil && !game.IsRegularTermination(err) {
		log.Err(err)
	} else if game.IsRegularTermination(err) {
		log.Info().Msg("Game gracefully shut down")
	}
}
