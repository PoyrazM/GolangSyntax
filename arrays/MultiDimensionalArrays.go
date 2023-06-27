package arrays

import (
	"fmt"
)

func MultiDimensionalArrays() {
	var numbers [2][3]int

	numbers[0][0] = 1
	numbers[0][1] = 3
	numbers[0][2] = 5

	numbers[1][0] = 2
	numbers[1][1] = 4
	numbers[1][2] = 6

	//log.Println(numbers[1][1])

	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i]); j++ {
			fmt.Print(numbers[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
