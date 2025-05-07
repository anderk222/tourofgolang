package basics

import "fmt"

// Constants cannot be declared using the := syntax.

const Pi float32 = 3.14

const OriginCountry string = "Ecuador"

func Constants() {

	const World = "世界"
	fmt.Println("\033[33m[Constants]\033[0m", "Hello", World)
	fmt.Println("\033[33m[Constants]\033[0m", "Happy", Pi, "Day")

	const Truth = true
	fmt.Println("\033[33m[Constants]\033[0m", "Go rules?", Truth)

	fmt.Println("\033[33m[Constants]\033[0m", "Where are u from?")
	fmt.Println("\033[33m[Constants]\033[0m", "I am from", OriginCountry)

}
