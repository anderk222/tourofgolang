package flowcontrol

import "fmt"

func ForIsGosWhile() {

	sum := 1
	for sum < 1000 {
		sum += sum
	}

	fmt.Println("\x1b[34m[ForIsGoWhile]\x1b[0m", sum)
}
