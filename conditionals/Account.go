package conditionals

import (
	"log"
)

func Account() {
	var accountPrice float64 = 1000
	var requestedPrice float64 = 900

	if requestedPrice > accountPrice {
		log.Fatalln("Requested Price is bigger than Account Price !")
	}

	if requestedPrice <= accountPrice {
		log.Println("Your money is getting ready.")
		accountPrice -= requestedPrice
	}

	log.Println("Process finished. Remaining Amount Is :", accountPrice)
}
