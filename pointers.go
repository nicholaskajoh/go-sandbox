package main

import "fmt"

func main() {
	x := 20
	y := &x // mem address of x
	fmt.Println(y, *y)
}
