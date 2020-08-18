package compare

func Square(i int) int {
	return i * i
}

// Dog is used as a demo type
type Dog struct {
	Name string
	Age  int
}

type DogWithFn struct {
	Name string
	Age  int
	Fn   func()
}
