package structs

import "log"

type customer struct {
	firstName   string
	lastName    string
	age         int
	phoneNumber string
}

func (c customer) save() {
	log.Println("Saved Customer : ", c)
}

func (c customer) update() {
	log.Println("Updated Customer : ", c)
}

func AddCustomer() {
	c := customer{firstName: "Mertcan", lastName: "Poyraz", age: 24, phoneNumber: "512312312"}
	c.save()
	c = customer{firstName: "Mertcan", lastName: "Poyraz", age: 24, phoneNumber: "123124124"}
	c.update()
}
