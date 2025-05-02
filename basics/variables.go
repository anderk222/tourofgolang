package basics

import "fmt"

var c, python, java bool

func Variables() {

	var i int

	colorCode := "32" // Verde

	fmt.Println("\x1b[34m[Variables]\x1b[0m")

	// Armas el string combinando el color dinámicamente
	colorStart := fmt.Sprintf("\035[%sm", colorCode)
	colorEnd := "\033[0m"

	fmt.Println(colorStart, i, c, python, java, colorEnd)

}
