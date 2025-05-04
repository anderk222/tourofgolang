package moretypes

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func ExcersiceMaps() {

	fmt.Println("\x1b[34m[ExcersiceMaps]\x1b[0m")

	wc.Test(WordCount)

}

func WordCount(s string) (m map[string]int) {

	words := strings.Split(s, " ")
	m = make(map[string]int)

	for _, word := range words {

		m[word]++

		// _, ok := m[word]

		// if ok {
		// m[word]++
		// } else {
		// m[word] = 1
		// }
	}

	return m
}
