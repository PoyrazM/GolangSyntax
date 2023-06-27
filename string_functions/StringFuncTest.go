package string_functions

import (
	"fmt"
	"log"
	s "strings"
)

func StringFuncTest() {
	name := "Mertcan"
	fmt.Println(s.Count(name, "M"))
	fmt.Println(s.Contains(name, "M"))
	fmt.Println(s.EqualFold(name, "Mertcan"))

	result := s.Index(name, "n")
	if result != -1 {
		log.Println("This text has a char")
	} else {
		log.Fatalln("This text has NOT a char !")
	}

	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))

}
