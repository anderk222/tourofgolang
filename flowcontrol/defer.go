package flowcontrol

import "fmt"

func Defer() {

	fmt.Println("\x1b[34m[Defer]\x1b[0m")

	defer fmt.Println("\x1b[34m[Defer]\x1b[0m", "world")

	fmt.Print("hello ")
}
