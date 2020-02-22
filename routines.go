package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	fmt.Println(s)
}

func main() {
	wg.Add(1)
	go say("Hello")
	wg.Add(1)
	go say("Hi")
	wg.Wait()

	wg.Add(1)
	go say("Hola")
	wg.Wait()
}
