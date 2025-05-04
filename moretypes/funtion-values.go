package moretypes

import (
	"fmt"
	"math"
)

func FunctionValues() {

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println("\x1b[34m[FunctionValues]\x1b[0m", hypot(5, 12))

	fmt.Println("\x1b[34m[FunctionValues]\x1b[0m", compute(hypot))
	fmt.Println("\x1b[34m[FunctionValues]\x1b[0m", compute(math.Pow))

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
