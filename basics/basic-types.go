package basics

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint       = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	Rune   rune       = 34 // rune == int32
	Byte   byte       = 65 // byte == uint8
)

func BasicTypes() {
	fmt.Printf("\033[34m[BasicTypes]\033[0m Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("\033[34m[BasicTypes]\033[0m Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("\033[34m[BasicTypes]\033[0m Type: %T Value: %v\n", z, z)
}
