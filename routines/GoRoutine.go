package routines

import (
	"log"
	"time"
)

func EvenNumbers() {
	for i := 0; i <= 10; i += 2 {
		log.Println("Even Number : ", i)
		time.Sleep(time.Second)
	}
}

func SingleNumbers() {
	for i := 1; i <= 10; i += 2 {
		log.Println("Single Number : ", i)
		time.Sleep(time.Second)
	}
}
