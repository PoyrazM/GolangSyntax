package interfaces

import "log"

func validate(i interface{}) {
	number, ok := i.(int)
	log.Println(number, ok)
}

func ValidateTest() {
	var number interface{} = 5
	validate(number)

	var text interface{} = "PoyrazM"
	validate(text)
}
