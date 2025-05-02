package flowcontrol

import "fmt"

func ForContinued() {

	sum := 1
	// its seems like a while i java mmmmm
	// for ; sum < 1000 like de example, but go compiler keep moving ";" :(

	for sum < 1000 {
		sum += sum
	}
	fmt.Println("\x1b[34m[ForContinued]\x1b[0m", sum)
}
