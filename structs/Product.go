package structs

import "fmt"

func Product() {
	fmt.Println(product{
		name:  "Computer",
		price: 5000,
		brand: "Apple",
	})

	fmt.Println(product{
		"Computer",
		4000,
		"Lenovo",
		20,
	})
}

type product struct {
	name         string
	price        float64
	brand        string
	discountRate int
}
