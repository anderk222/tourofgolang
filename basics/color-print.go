package basics

import "fmt"

func ColorPrint() {

	fmt.Println("\033[34m[ColorPrint]\033[0m", "\x1b[34mThis is a color print example.\x1b[0m")

	fmt.Println("\033[34m[ColorPrint]\033[0m", "\033[36mAnother one\033[0m")

}
