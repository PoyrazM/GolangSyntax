package error_handling

import (
	"fmt"
	"log"
)

type borderException struct {
	parameter int
	message   string
}

func (b *borderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func guessException(guessedNumber int) (string, error) {
	if guessedNumber < 1 || guessedNumber > 100 {
		return "", &borderException{parameter: guessedNumber, message: "please enter a number in the range of 1-100"}
	}
	return "Congratulations", nil
}

func Guess(guessedNumber int) {
	message, err := guessException(guessedNumber)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println(message)
	}
}
