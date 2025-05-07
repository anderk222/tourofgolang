package methods

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods() {

	v := Vertex{12, -34}

	fmt.Println("\033[33m[Methods]\033[0m", v.Abs())

}
