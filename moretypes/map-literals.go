package moretypes

import "fmt"

func MapLiterals() {

	var m = map[string]Cordinate{
		"Bell Labs": Cordinate{34.343, 3434.343},
		"Google":    Cordinate{33343.54, 34343.34434},
	}

	fmt.Println("\033[33m[MapLiterals]\033[0m", m)

}
