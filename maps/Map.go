package maps

import (
	"fmt"
	"log"
)

func Dictionary() {
	dictionary := make(map[string]string)
	dictionary["table"] = "masa"
	dictionary["pencil"] = "kalem"
	dictionary["book"] = "kitap"

	fmt.Println(dictionary)
	fmt.Println(dictionary["book"])
	delete(dictionary, "book")
	fmt.Println(dictionary["book"])
	fmt.Println(len(dictionary))
	fmt.Println(dictionary)

	value, isHas := dictionary["pencil"]
	log.Println(isHas)
	if isHas {
		log.Println(value)
	}

	dictionary2 := map[string]string{
		"glass":      "bardak",
		"microphone": "mikrofon",
		"headphones": "kulaklık",
	}

	log.Println(dictionary2)
}
