package main

import (
	"fmt"
	"sync"
)

/*
func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}*/

var wg = sync.WaitGroup{}

func main() {
	var msg = "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)
	msg = "Goodbye"
	wg.Wait()
}
