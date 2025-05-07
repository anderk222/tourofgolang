package moretypes

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Range() {

	for i, v := range pow {
		fmt.Printf("\x1b[33m[Range]\x1b[0m  2**%d = %d\n", i, v)
	}

}
