package flowcontrol

import (
	"fmt"
	"runtime"
)

func Switch() {

	fmt.Println("\x1b[33m[Switch]\x1b[0m]")

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
