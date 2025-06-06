package moretypes

import "fmt"

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func StructLiterals() {

	fmt.Println("\x1b[33m[StructLiterals]\x1b[0m", v1, p, v2, v3)
}
