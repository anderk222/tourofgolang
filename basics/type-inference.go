package basics

import "fmt"

//i := 42           // int
//f := 3.142        // float64
//g := 0.867 + 0.5i // complex128

func TypeInference() {

	v := 0.867 + 0.5i // change me!

	fmt.Printf("\x1b[33m[TypeInference]\x1b[0m v is of type %T\n", v)

}
