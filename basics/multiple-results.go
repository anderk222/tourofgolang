package basics

import "fmt"

func MultipleResults() {

	// a, b, c := swap("MultipleResults", "Hello", "World")

	fmt.Println(swap("\x1b[34m[MultipleResults]\x1b[0m", "Hello", "World"))

}

func swap(x, y, z string) (string, string, string) {
	return x, y, z
}
