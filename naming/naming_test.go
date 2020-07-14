package naming

import "testing"

func TestDog(t *testing.T) {

}

func TestDog_Bark_muzzled(t *testing.T) {

}

func TestDog_Bark_unmuzzled(t *testing.T) {

}

func TestSpeak_spanish(t *testing.T) {}

func TestColor(t *testing.T) {
	arg := "blue"
	want := "#0000FF"
	got := Color("blue")
	if got != want {
		t.Errorf("Color(%q) = %s; want %s", arg, got, want)
	}
}
