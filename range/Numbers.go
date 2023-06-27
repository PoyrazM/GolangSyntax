package _range

import (
	"fmt"
)

func SumSingleNumbers() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	total := 0

	for _, number := range numbers {

		isSingle := number%2 != 0

		if isSingle {
			total += number
		}
	}

	fmt.Println("Total :", total)
}
