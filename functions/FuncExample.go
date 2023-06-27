package functions

import "log"

func Sum(number1 int, number2 int) int {
	return number1 + number2
}

func GreetName(name string) {
	log.Println("Hello ,", name)
}
