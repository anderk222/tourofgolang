package basics

import (
	"fmt"
	"math"
)

func ExportdNames() {

	// fmt.Println(math.pi); // Esto daría error porque Go es case-sensitive.

	fmt.Println("\x1b[33m[ExportdNames]\x1b[0m", "Pi", math.Pi) // Esto funciona porque `math.Pi` está correctamente capitalizado.

}
