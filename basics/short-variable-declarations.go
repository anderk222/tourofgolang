package basics

import "fmt"

// No puedes usar := a  nivel de paquete

func ShortVariableDeclarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println("\x1b[34m[ShortVariableDeclarations]\x1b[0m", i, j, k, c, python, java)
}
