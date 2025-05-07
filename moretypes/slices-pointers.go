package moretypes

import "fmt"

/*
A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.
*/
func SlicesPointers() {

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println("\x1b[33mSlicesPointers]\x1b[0m", names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println("\x1b[33mSlicesPointers]\x1b[0m", a, b)

	b[0] = "XXX"
	fmt.Println("\x1b[33mSlicesPointers]\x1b[0m", a, b)
	fmt.Println("\x1b[33mSlicesPointers]\x1b[0m", names)
}
