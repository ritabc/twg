package cover

import "testing"

func TestTriangle(t *testing.T) {
	tests := []struct {
		base, height, want float64
	}{
		{2, 5, 5},
		{2, 2, 2},
		{11, 4, 22},
	}
	for _, tt := range tests {
		got := Triangle(tt.base, tt.height)
		if got != tt.want {
			t.Errorf("Triangle(%f, %f) = %f; want %f", tt.base, tt.height, got, tt.want)
		}
	}
}

// Running coverage on this will give us false sense of security, we'll not find the bug this way
func TestSquare(t *testing.T) {
	for i := 0.0; i < 100.0; i++ {
		want := Triangle(i, i) * 2
		got := Square(i)
		if got != want {
			t.Errorf("Square(%f) = %f; want %f", i, got, want)
		}
	}
}
