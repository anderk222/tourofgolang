package moretypes

import "fmt"

func MapLiteralsContinued() {

	var m = map[string]Cordinate{
		"Bell Labs": {234.3434, -123.43434},
		"Google":    {2343.3434, -1232.343434},
	}

	fmt.Println("\x1b[33m[MapLiteralsContinued]\x1b[0m", m)

}
