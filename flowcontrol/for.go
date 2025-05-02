package flowcontrol

import "fmt"

func For() {

	sum := 0
	// we don't need parentheses:)) java sucks
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("\x1b[34m[For]\x1b[0m", sum)
}
