package moretypes

import "fmt"

/*
var a [10]int

these slice expressions are equivalent:

a[0:10]
a[:10]
a[0:]
a[:]

*/

func SliceBounds() {

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println("\x1b[33m[SliceBounds]\x1b[0m", s)

	s = s[:2]
	fmt.Println("\x1b[33m[SliceBounds]\x1b[0m", s)

	s = s[1:]
	fmt.Println("\x1b[33m[SliceBounds]\x1b[0m", s)

}
