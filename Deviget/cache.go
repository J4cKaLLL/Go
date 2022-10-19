// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	a := [2]int{2, 3}
	b := [2]int{4, 6}
	c := [2]int{9, 12}
	fmt.Printf("Esta es la distancia promedio: %d", validate(a, b, c))
}

func validate(a [2]int, b [2]int, c [2]int) (result int) {

	x1 := b[0] - a[0]
	y1 := b[1] - a[1]
	z1 := math.sqrt(x1*x1 + y1*y1)

	x2 := b[0] - c[0]
	y2 := b[1] - c[1]
	z2 := math.sqrt(x2*x2 + y2*y2)

	x3 := c[0] - a[0]
	y3 := c[1] - a[1]
	z3 := math.sqrt(x3*x3 + y3*y3)

	return (z1 + z2 + z3) / 3

}
