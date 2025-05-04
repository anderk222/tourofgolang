package moretypes

import "fmt"

func MutatingMaps() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("\033[34m[MutatingMaps]\033[0m", "The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("\033[34m[MutatingMaps]\033[0m", "The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("\033[34m[MutatingMaps]\033[0m", "The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("\033[34m[MutatingMaps]\033[0m", "The value:", v, "Present?", ok)
}
