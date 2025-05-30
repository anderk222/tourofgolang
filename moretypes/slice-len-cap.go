package moretypes

import "fmt"

func SliceLenCap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("\x1b[33m[SliceLenCap]\x1b[0m len=%d cap=%d %v\n", len(s), cap(s), s)
}
