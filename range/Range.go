package _range

import "log"

func Range() {
	cities := []string{"Paris", "London", "Amsterdam"}

	for _, city := range cities {
		log.Println(city)
	}
}
