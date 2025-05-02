package flowcontrol

import (
	"fmt"
)

func sqrtt(x float64) float64 {

	z := 1.0
	var currentZ float64 = 0.0

	for i := 0; i <= 10; i++ {

		currentZ -= (z*z - x) / (2 * z)

		if currentZ == z {
			break
		}

		z = currentZ
	}

	return z
}

func ExerciseLoopsAndFuntions(x float64) {

	fmt.Println("\x1b[34m[ExerciseLoopsAndFuntions]\x1b[0m x=", x, "raiz cuadrada=", sqrtt(x))

}
