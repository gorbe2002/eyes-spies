package logoDisplay

import (
	"fmt"
	"github.com/qeesung/image2ascii/convert"
	_ "image/jpeg"
	_ "image/png"
)

func ShowLogo(filename string) {
	// Create Convert options
	conv_opts := convert.DefaultOptions

	converter := convert.NewImageConverter()
	fmt.Println(converter.ImageFile2ASCIIString(filename, &conv_opts))
}
