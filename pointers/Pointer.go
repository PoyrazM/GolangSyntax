package pointers

import (
	"fmt"
	"log"
)

func Pointer(number *int) {
	*number++
	fmt.Println("Number is in the Function :", *number)
}

func Test() {
	number := 10
	Pointer(&number)
	log.Println(number)
}
