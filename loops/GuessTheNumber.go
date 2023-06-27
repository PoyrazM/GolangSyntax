package loops

import (
	"fmt"
	"log"
)

func GuessTheNumber() {
	number := 100
	guessedNumber := 0

	log.Println("Please guess a number :")
	_, err := fmt.Scanf("%d", &guessedNumber)
	if err != nil {
		log.Fatal("It is not a number !", guessedNumber, ". Please try again.")
	}

	try := 1
	for number != guessedNumber {
		if number > guessedNumber {
			log.Println("It is wrong guess ! Please enter the higher number.")
			fmt.Scanln(&guessedNumber)
		} else if number < guessedNumber {
			log.Println("It is wrong guess ! Please enter the lesser number.")
			fmt.Scanln(&guessedNumber)
		}
		try++
	}

	message := getTryMessage(try)

	log.Println("Congratulations ! Your guessed number is true : ", guessedNumber)
	log.Println("You knew it on your trial :", try, message)
}

func getTryMessage(try int) string {
	message := ""

	if try <= 3 {
		message = "Amazing."
	} else if try > 6 {
		message = "Bad."
	} else {
		message = "Good."
	}
	return message
}
