package compare

import (
	"reflect"
	"testing"
)

func TestSquare(t *testing.T) {
	arg := 4
	want := 16
	got := Square(arg)
	if got != want {
		t.Errorf("Square(%d) = %d; want %d", arg, got, want)
	}
}
func TestDog(t *testing.T) {
	morty := Dog{
		Name: "Morty",
		Age:  3,
	}
	morty2 := Dog{
		Name: "Morty",
		Age:  3,
	}
	if morty != morty2 {
		t.Errorf("morty != morty2")
	}
	odie := Dog{
		Name: "Odie",
		Age:  12,
	}
	if morty == odie {
		t.Errorf("morty == odie")
	}
}

func TestPtf(t *testing.T) {
	morty := &Dog{
		Name: "Morty",
		Age:  3,
	}
	morty2 := &Dog{
		Name: "Morty",
		Age:  3,
	}
	if morty == morty2 {
		t.Errorf("morty == morty2")
	}
}

func TestDogWithFn(t *testing.T) {
	// morty := DogWithFn{
	// 	Name: "Morty",
	// 	Age:  3,
	// }
	// morty2 := DogWithFn{
	// 	Name: "Morty",
	// 	Age:  3,
	// }
	// if morty != morty // cannot compare two structs with funcs
	// reflect.DeepEqual(morty, morty2) // this is true

	mortyP := &DogWithFn{
		Name: "Morty",
		Age:  3,
	}
	morty2P := &DogWithFn{
		Name: "Morty",
		Age:  3,
	}

	if !reflect.DeepEqual(mortyP, morty2P) {
		t.Errorf("MortyP != Morty2P") // Passes
	}
}
