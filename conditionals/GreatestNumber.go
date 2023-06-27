package conditionals

import "log"

func FindGreaterNumber() {

	number1, number2, number3 := 10, 15, 110

	greatestNumber := number1

	isNumber2Greatest := number2 > greatestNumber
	if isNumber2Greatest {
		greatestNumber = number2
	}

	isNumber3Greatest := number3 > greatestNumber
	if isNumber3Greatest {
		greatestNumber = number3
	}

	log.Println("Greatest Number Is :", greatestNumber)
}
