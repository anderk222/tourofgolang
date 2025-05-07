package moretypes

import "fmt"

func Slices() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int

	s = append(s, 12)
	s = append(s, 14, 16)

	var ps []int = primes[1:3]

	fmt.Println("\x1b[33m[Slices]\x1b[0m", s)
	fmt.Println("\x1b[33m[Slices]\x1b[0m", ps)
}
