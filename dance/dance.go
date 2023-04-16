package dance

import (
	"github.com/qeesung/asciiplayer/pkg/player"
)

	
// GenDisplayDefault: generates the display from the string of the gif location
func GenDisplayDefault(filename string) {
	// create a new gif terminal player
	playa := player.NewGifTerminalPlayer()

	// ensures that 
	def_opts := &player.DefaultPlayOptions

	// plays on a loop
	// https://github.com/qeesung/asciiplayer/blob/master/pkg/player/gif_player.go#L26-L46
	playa.Play(filename, def_opts)
}
