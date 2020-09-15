package draw_test

import (
	"image/draw"
	"testing"

	twgdraw "github.com/ritabc/twg/draw"
)

func TestFibGradient(t *testing.T) {
	var im draw.Image
	twgdraw.FibGradient(im)
}

func TestFibFunc(t *testing.T) {
	got := twgdraw.Fib(2)
	if got != 1 {
		t.Errorf("Fib(2) = %d, want 1", got)
	}
}
