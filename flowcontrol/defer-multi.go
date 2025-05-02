package flowcontrol

import "fmt"

// El orden ejecucion de los defers, si tenemos varios, se ejecutan primero los ultimos
// instanciados en nuestro programa

func DeferMulti() {

	fmt.Println("\x1b[34m[DeferMulti]\x1b[0m", "counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println("\x1b[34m[DeferMulti]\x1b[0m", i)
	}

	fmt.Println("\x1b[34m[DeferMulti]\x1b[0m", "done")
}
