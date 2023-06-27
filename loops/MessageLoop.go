package loops

import "log"

func MessageLoop() {
	message := "This is a message !"

	for i := 0; i < 10; i++ {
		log.Println("Index of :", i, "->", message)
	}
}

func MessageLoop2() {
	message := "This is a message 2 !"

	i := 0
	for i < 10 {
		i++
		if i == 5 {
			continue
		}
		log.Println("Index of :", i, "->", message)
	}
}
