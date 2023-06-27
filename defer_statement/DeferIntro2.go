package defer_statement

import "log"

func numberMessage(number int32) string {
	if number%2 == 0 {
		return "Even Number"
	} else {
		return "Single Number"
	}
}

func deferMessage() {
	log.Println("It is finished.")
}

func TestNumberMessage() {
	defer deferMessage()
	message := numberMessage(1)
	log.Println(message)
}
