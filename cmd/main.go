package cmd

import (
	"image"
	"image/draw"
	"image/png"
	"os"

	twgdraw "github.com/ritabc/twg/draw"
)

func main() {
	const w, h = 1000, 10000
	var im draw.Image
	im = image.NewRGBA(image.Rectangle{Max: image.Point{X: w, Y: h}})
	im = twgdraw.FibGradient(im)
	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, im)
	if err != nil {
		panic(err)
	}
}
