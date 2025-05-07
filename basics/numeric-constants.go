package basics

import "fmt"

// Numeric constants are high-precision values.

// An untyped constant takes the type needed by its context.

// Try printing needInt(Big) too.

// (An int can store at maximum a 64-bit integer, and sometimes less.)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func NumericConstants() {

	fmt.Println("\x1b[33m[NumericConstants]\x1b[0m", needInt(Small))
	fmt.Println("\x1b[33m[NumericConstants]\x1b[0m", needFloat(Small))
	fmt.Println("\x1b[33m[NumericConstants]\x1b[0m", needFloat(Big))

}
