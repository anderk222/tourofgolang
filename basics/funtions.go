package basics

import "fmt"

func Funtions() {

	fmt.Println("\x1b[34m[Funtions]\x1b[0m", add(13, 43))

}

func add(i int, j int) int {
	return i + j
}
