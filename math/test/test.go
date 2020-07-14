package main

import (
	"fmt"

	"github.com/ritabc/twg/math"
)

func main() {
	sum := math.Sum([]int{10, -2, 3})
	if sum != 11 {
		panic(fmt.Sprintf("FAIL: Wanted 11, but received %d", sum))
	}
	add := math.Add(5, 10)
	if add != 15 {
		panic(fmt.Sprintf("FAIL: Wanted 15, but received %d", add))
	}
	fmt.Println("PASS")
}
