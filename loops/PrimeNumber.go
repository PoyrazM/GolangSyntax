package loops

import (
	"fmt"
	"log"
)

func FindPrimeNumber() {
	userNumber := 0

	log.Println("Please enter the number :")
	fmt.Scanln(&userNumber)

	isPrime := true

	for i := 2; i < userNumber; i++ {
		if userNumber%i == 0 {
			isPrime = false
		}
	}

	if isPrime {
		log.Println("The Number is Prime number.")
	} else {
		log.Println("The Number is NOT Prime number.")
	}
}
