package loops

import (
	"log"
)

func FindFriendsNumber() {
	number1 := 220
	number2 := 284

	total1 := calculateNumber(number1)
	total2 := calculateNumber(number2)

	status := checkIsFriendNumbers(total1, number2) && checkIsFriendNumbers(total2, number1)

	if status {
		log.Println("These numbers are friend numbers")
	} else {
		log.Println("These numbers are NOT friend numbers !")
	}
}

func calculateNumber(number int) int {
	total := 0

	for i := 1; i < number; i++ {
		if number%i == 0 {
			total += i
		}
	}
	return total
}

func checkIsFriendNumbers(number int, total int) bool {
	isFriendNumbers := false
	if total == number {
		isFriendNumbers = true
	}
	return isFriendNumbers
}
