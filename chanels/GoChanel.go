package chanels

import (
	"fmt"
	"time"
)

func EvenNumbers(evenNumberCn chan int) {
	total := 0
	for i := 0; i <= 10; i += 2 {
		total += i
		fmt.Println("Even Numbers calculator is running")
		time.Sleep(time.Second)
	}

	evenNumberCn <- total
}

func SingleNumbers(singleNumberCn chan int) {
	total := 0
	for i := 1; i <= 10; i += 2 {
		total += i
		fmt.Println("Single Numbers calculator is running")
		time.Sleep(time.Second)
	}

	singleNumberCn <- total
}
