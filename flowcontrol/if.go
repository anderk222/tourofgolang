package flowcontrol

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func If() {

	fmt.Println("\x1b[34m[If]\x1b[0m ", sqrt(2), sqrt(-4))

}
