package flowcontrol

import "fmt"

func Run() {
	fmt.Println("\033[35m[flowcontrol]\033[0m")

	For()
	ForContinued()
	ForIsGosWhile()
	Forever()
	If()
	IfWithShortStatement()
	IfAndElse()
	ExerciseLoopsAndFuntions(9)
	Switch()
	SwitchWithNoCondition()
	Defer()

	DeferMulti()

}
