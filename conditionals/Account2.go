package conditionals

import "log"

func AccountTwo() {
	var accountPrice float64 = 1000
	var requestedPrice float64 = 1000

	if requestedPrice > accountPrice {
		log.Fatalln("Requested Price is bigger than Account Price !")
	} else if accountPrice == requestedPrice {
		log.Println("Your money is getting ready.")
		log.Println("Attention ! You have no money left in your bank account !")
		accountPrice = withDrawMoney(accountPrice, requestedPrice)
	} else {
		log.Println("Your money is getting ready.")
		accountPrice = withDrawMoney(accountPrice, requestedPrice)
	}

	log.Println("Process finished. Remaining Amount Is :", accountPrice)
}

func withDrawMoney(accountMoney, requestedPrice float64) float64 {
	var leftMoney float64
	leftMoney = accountMoney - requestedPrice
	return leftMoney
}
