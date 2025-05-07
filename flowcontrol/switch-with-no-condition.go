package flowcontrol

import (
	"fmt"
	"time"
)

func SwitchWithNoCondition() {
	t := time.Now()

	fmt.Println("\033[33m[SwitchWithNoCondition]\033[0m")
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
