package example_test

import (
	"fmt"
	"io"

	// Needed for initialize side effect
	_ "image/png"

	"github.com/ritabc/twg/example"
)

var _ = ""

func Example_crop() {
	var r io.Reader
	img, err := example.Decode(r)
	if err != nil {
		panic(err)
	}
	err = example.Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}
	var w io.Writer
	err = example.Encode(img, w)
	if err != nil {
		panic(err)
	}
	fmt.Println("See out.jpg for the cropped image.")

	// OUtput:
	// See out.jpg for the cropped image.
}
