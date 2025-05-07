package basics

import "fmt"

func NamedResults() {

	fmt.Print("\x1b[33m[NamedResults]\x1b[0m ")

	fmt.Print(split(12))

}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
