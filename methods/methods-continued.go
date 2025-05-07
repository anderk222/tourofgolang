package methods

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {

	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MethodsContinued() {

	f := MyFloat(-math.Sqrt2)

	var f2 MyFloat = 0.0

	fmt.Println("\033[33m[MethodsContinued]\033[0m", f.Abs(), f2.Abs())

}
