package error_handling

import (
	"errors"
	"log"
)

func checkGuessNumber(guessNumber int) (string, error) {
	if guessNumber < 1 || guessNumber > 100 {
		return "", errors.New("please enter a number in the range of 1-100")
	}
	return "Congratulations", nil
}

func GuessNumberExceptionTest() {
	message, err := checkGuessNumber(10)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println(message)
	}
}
