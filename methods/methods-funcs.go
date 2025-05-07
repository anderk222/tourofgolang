package methods

import (
	"fmt"
	"math"
)

/*
Methods are functions
Remember: a method is just a function with a receiver argument.

Here's Abs written as a regular function with no change in functionality.
*/

func Abs(v Vertex) float64 {

	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

func MethodsFuncs() {

	v := Vertex{12, -34}
	fmt.Println("\033[33m[MethodsFuncs]\033[0m", Abs(v))
}
