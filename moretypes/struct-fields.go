package moretypes

import "fmt"

func StructFields() {

	v := Vertex{19, 12}

	v.X = 49

	fmt.Println("\x1b[33m[StructFields]\x1b[0m", v.X)

}
