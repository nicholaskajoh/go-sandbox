package main

import "fmt"

func add(x float64, y float64) float64 {
	return x + y
}

func str(a, b string) (string, string) {
	return a, b
}

func main() {
	var num1 float64 = 5.6
	var num2 float64 = 9.5
	num3 := 9.0
	fmt.Println(add(num1, num2))
	fmt.Println(add(num1, num3))

	w1, w2 := "Hello", "World"
	fmt.Println(str(w1, w2))
}
