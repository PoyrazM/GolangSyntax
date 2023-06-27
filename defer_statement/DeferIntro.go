package defer_statement

import "log"

func AfterClose() {
	log.Println("Closed.")
}

func Close() {
	defer AfterClose()
	log.Println("Connection lost ...")
}
