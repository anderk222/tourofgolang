package basics

import (
	"fmt"
	"math"
)

func TypeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println("\x1b[33m[TypeConversion]\x1b[0m", x, y, z)
}
