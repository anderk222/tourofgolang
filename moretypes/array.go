package moretypes

import "fmt"

func Array() {

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println("\x1b[34m[Array]\x1b[0m", a[0], a[1])
	fmt.Println("\x1b[34m[Array]\x1b[0m", a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("\x1b[34m[Array]\x1b[0m", primes)

	vertexs := [2]Vertex{
		{X: 1, Y: 2},
		{X: 3, Y: 4},
	}

	fmt.Println("\x1b[34m[Array]\x1b[0m", vertexs)
}
