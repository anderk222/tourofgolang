package moretypes

import "fmt"

type Cordinate = struct {
	Lat, Long float64
}

func Maps() {

	m := make(map[string]Cordinate)

	m["Bell Labs"] = Cordinate{40.68433, -74.3996}

	fmt.Println("\x1b[33m[Maps]\x1b[0m", m["Bell Labs"])

}
