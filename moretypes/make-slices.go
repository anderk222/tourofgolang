package moretypes

import "fmt"

/*
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:
*/

func MakeSlices() {

	a := make([]int, 5)
	printSlicee("a", a)

	b := make([]int, 0, 5)
	printSlicee("b", b)

	c := b[:2]
	printSlicee("c", c)

	d := c[2:5]
	printSlicee("d", d)

}
func printSlicee(s string, x []int) {
	fmt.Printf("\x1b[34m[MakeSlices]\x1b[0m %s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
