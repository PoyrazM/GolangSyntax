package interfaces

import "fmt"

type mortgage struct {
	creditPaymentTotal float64
	homeAddress        string
	rate               float64
}

type car struct {
	creditPaymentTotal float64
	carDetails         string
	rate               float64
}

type creditCalculator interface {
	calculate() float64
}

func (m mortgage) calculate() float64 {
	return m.creditPaymentTotal * (m.rate / 1200)
}

func (c car) calculate() float64 {
	return c.creditPaymentTotal * (c.rate / 1200)
}

func calculateMonthlyPayment(credits []creditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal += credits[i].calculate()
	}
	return paymentTotal
}

func CalculateUserCredit() {
	credit1 := mortgage{rate: 10, creditPaymentTotal: 100000, homeAddress: "Istanbul"}
	credit2 := mortgage{rate: 12, creditPaymentTotal: 50000, homeAddress: "Ankara"}
	credit3 := car{rate: 15, creditPaymentTotal: 60000, carDetails: "Renault Clio"}

	credits := []creditCalculator{credit1, credit2, credit3}
	total := calculateMonthlyPayment(credits)

	fmt.Println("Total Payment :", total)
}
