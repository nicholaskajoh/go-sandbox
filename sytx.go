package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo(bar float64) {
	fmt.Println("The square root of", bar, "is", math.Sqrt(bar))
}

func main() {
	foo(16)

	fmt.Println("Random number:", rand.Intn(500))
}
