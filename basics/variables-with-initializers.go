package basics

import "fmt"

var i, j int = 1, 2

func VariablesWithInitializers() {

	var c, python, java = true, false, "no!"
	fmt.Println("\x1b[34m[VariablesWithInitializers]\x1b[0m", i, j, c, python, java)

}
