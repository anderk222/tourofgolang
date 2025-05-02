package flowcontrol

import "fmt"

func Forever() {

	fmt.Print("\x1b[34m[Forever]\x1b[0m ")

	for {
		fmt.Println("if I wouldn't call 'break', This will run forever")
		break
	}
}
