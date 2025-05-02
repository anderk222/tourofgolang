package basics

import "fmt"

func Run() {

	fmt.Println("\033[35m[basics]\033[0m")

	Imports()

	ExportdNames()

	Funtions()

	FuntionsContinued()

	MultipleResults()

	NamedResults()

	Variables()

	VariablesWithInitializers()

	BasicTypes()
	Zero()
	TypeConversion()
	TypeInference()

	Constants()

	NumericConstants()
	// var count int = 10

	// fmt.Print("\rCount ", 0)
	// for i := 0; i < count; i++ {

	// 	time.Sleep(1 * time.Second)
	// 	fmt.Print("\rCount \x1b[33m", i, "\x1b[0m")

	// }
	// if runtime.GOOS == "windows" {
	// 	cmd := exec.Command("cmd", "/c", "cls")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// } else {
	// 	cmd := exec.Command("clear")
	// 	cmd.Stdout = os.Stdout
	// 	cmd.Run()
	// }

	// fmt.Printf("\nCount %d", count)

	ColorPrint()

}
