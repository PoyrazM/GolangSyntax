package variables

import "fmt"

func Vars() {
	var text string = "This is a text"
	fmt.Println(text)

	var number int = 10
	fmt.Println(number)
	fmt.Println((10 * 20 / 3) * number)

	var isBoolean bool = true
	fmt.Println(isBoolean)

	var float float32 = 1.25
	fmt.Println(float)

	text2 := "I am a other text"
	fmt.Println(text2)

	number2 := 2.5
	fmt.Println(number2 * 10)
	fmt.Printf("Data type : %T", number2)
}
