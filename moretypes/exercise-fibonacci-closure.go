package moretypes

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {

	current, last, first := 0, 0, true

	return func() int {

		if !first && current == 0 && last == 0 {
			last = 1
		} else {
			first = false
		}

		current, last = (last + current), current

		return current

	}
}

func ExersiceFibonacciClousure() {

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
