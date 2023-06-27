package pointers

import "fmt"

func ListPointer(numbers []int) {
	numbers[0] = 200
	fmt.Println("Number is in the Function : ", numbers[0])
}
