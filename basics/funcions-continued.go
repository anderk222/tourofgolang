package basics

import "fmt"

func FuntionsContinued() {

	fmt.Println("\x1b[33m[FunctionsContinued]\x1b[0m", addSameTypesParameters(3, 5))

}

func addSameTypesParameters(i, j int) int {
	return i + j
}
