package moretypes

/*
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers.
When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func ExcersiceSlice() {

	fmt.Println("\x1b[33m[ExcersiceSlice]\x1b[0m", Pic(10, 10))

	pic.Show(Pic)

}

func Pic(dx, dy int) [][]uint8 {
	var image [][]uint8 = make([][]uint8, dy)

	for i := 0; dy > i; i++ {

		image[i] = make([]uint8, dx)

		for j := 0; j < dx; j++ {
			image[i][j] = uint8(j ^ i) // (x+y)/2, x*y, and x^y
		}
	}

	return image
}
