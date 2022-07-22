package main

import (
	"fmt"
	"net/http"
)

func test(a int, b string) (int, string) {
	return (a + 1), ("Hello " + b)
}

func passBy(val1 int, val2 *int) {
	val1 = 10
	*val2 = 100
}

func sum(nums ...int) int {
	sum := 0
	for n := range nums {
		sum += n
	}
	return sum
}
func sayHello(name string) {
	fmt.Println("Hello: ", name)
}

func factorial(m int) int {
	if m == 1 {
		return 1
	}
	return m * factorial(m-1)
}

func anonymous() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("reached /")
	})
	f := func() {
		fmt.Println("Anonymous Function")
	}
	f()
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b - a
	}
}

func main() {
	// fmt.Println(test(3, "World"))
	// a, b := 1, 2
	// passBy(a, &b)
	// fmt.Println(a, b)
	// fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	//hello := sayHello
	//hello("jame")
	//fmt.Println(factorial(4))
	//anonymous()
	// counter1 := counter()
	// fmt.Println(counter1())
	// fmt.Println(counter1())
	// fmt.Println(counter1())
	// counter2 := counter()
	// fmt.Println(counter2())
	// fmt.Println(counter2())
	// fmt.Println(counter2())

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
