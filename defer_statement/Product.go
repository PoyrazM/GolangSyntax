package defer_statement

import "log"

type product struct {
	productName string
	price       float64
	currency    string
}

func (p product) save() {
	log.Println("Saved Product :", p.productName)
	defer logMessage()
}

func logMessage() {
	log.Println("Logged.")
}

func SaveProductTest() {
	p := product{productName: "MacBook Pro 32 GB", price: 80000, currency: "TL"}
	defer p.save()
	log.Println("Product successfully created.")
}
