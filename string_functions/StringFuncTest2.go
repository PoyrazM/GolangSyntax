package string_functions

import (
	"fmt"
	s "strings"
)

func StringFuncTest2() {
	name := "Mertcan"
	fmt.Println(s.HasPrefix(name, "Mert"))
	fmt.Println(s.HasSuffix(name, "can"))

	words := []string{"m", "e", "r", "t", "c", "a", "n"}
	result := s.Join(words, "-")
	fmt.Println(result)

	fmt.Println(s.Replace(result, "-", "", -1))
	fmt.Println(s.Split(result, "-"))
	fmt.Println(s.Repeat(result, 5))
}
