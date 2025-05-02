package flowcontrol

import (
	"fmt"
	"math"
)

// Variables declared inside an if short statement are also available
// inside any of the else blocks.

//(Both calls to pow return their results before the call to fmt.Println in basics begins.)

func pow_with_else(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func IfAndElse() {

	fmt.Println(
		"\x1b[34m[IfAndElse]\x1b[0m",
		pow_with_else(3, 2, 10),
		pow_with_else(3, 3, 20),
	)

}
