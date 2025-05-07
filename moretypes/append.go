package moretypes

import "fmt"

func Append() {

	var s []int
	printSliceee(s)

	// append works on nil slices.
	s = append(s, 0)
	printSliceee(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSliceee(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSliceee(s)

}

func printSliceee(s []int) {
	fmt.Printf("\033[33m[Append]\033[0m len=%d cap=%d %v\n", len(s), cap(s), s)
}
