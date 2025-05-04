package moretypes

import "fmt"

// Go has pointers. A pointer holds the memory address of a value.
// The type *T is a pointer to a T value. Its zero value is nil.

// var p *int

func Pointers() {

	i, j := 42, 2701

	p := &i                                      // point to i
	fmt.Println("\x1b[34m[Pointers]\x1b[0m", *p) // read i through the pointer
	*p = 21                                      // set i through the pointer
	fmt.Println("\x1b[34m[Pointers]\x1b[0m", i)  // see the new value of i

	p = &j                                      // point to j
	*p = *p / 37                                // divide j through the pointer
	fmt.Println("\x1b[34m[Pointers]\x1b[0m", j) // see the new value of j
}
