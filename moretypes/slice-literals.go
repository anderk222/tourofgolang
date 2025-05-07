package moretypes

import "fmt"

func SliceLiterals() {

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("\033[33mSliceLiterals]\033[0m", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("\033[33mSliceLiterals]\033[0m", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("\033[33mSliceLiterals]\033[0m", s)
}
