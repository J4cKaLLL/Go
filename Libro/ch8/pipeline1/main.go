package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {

		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
		fmt.Println("cierra counter")
	}()
	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
		fmt.Println("cierra sqarterr")
	}()
	// Printer
	for x := range squares {
		fmt.Println(x)
	}
}
