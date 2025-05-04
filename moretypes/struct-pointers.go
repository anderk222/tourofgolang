package moretypes

import "fmt"

// Struct fields can be accessed through a struct pointer.

// To access the field X of a struct when we have the struct pointer p we could write (*p).X. However,
// that notation is cumbersome,
// so the language permits us instead to write just p.X, without the explicit dereference.

func StructPointers() {
	v := Vertex{12, 12}

	p := &v

	p.X = 1e9

	fmt.Println("\x1b[34m[StructPointers]\x1b[0m", p.X)

}
