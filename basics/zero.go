package basics

import "fmt"

// Variables declared without an explicit initial value are given their zero value.

// The zero value is:

// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.
func Zero() {

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("\x1b[34m[Zero]\x1b[0m %v %v %v %q\n", i, f, b, s)

}
