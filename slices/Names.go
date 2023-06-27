package slices

import "fmt"

func Names() {
	names := make([]string, 3)

	fmt.Println(names)
	names[0] = "Mertcan"
	names[1] = "Mertcan2"
	names[2] = "Mertcan3"
	names = append(names, "Mertcan4")
	fmt.Println(names)
	fmt.Println(len(names))
}
