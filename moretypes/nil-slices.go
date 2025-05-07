package moretypes

import "fmt"

func NilSlices() {

	var slice []int

	fmt.Println("\033[33m[NilSlices]\033[0m", slice, len(slice), cap(slice))

	if slice == nil {

		fmt.Println("\033[33m[NilSlices]\033[0m", "nil!")

	}

}
