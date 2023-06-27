package arrays

import "log"

func Numbers2() {
	numbers := [5]int{10, 20, 40, 50, 25}

	greatestNumber := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > greatestNumber {
			greatestNumber = numbers[i]
		}
	}
	log.Println("Greatest Number :", greatestNumber)
}
