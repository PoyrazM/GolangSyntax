package variables

import "fmt"

func LogicalVars() {

	var status bool

	var name1 string = "Mertcan"
	var name2 string = "Kağan"

	status = name1 == name2
	fmt.Println("Is Equals :", status)

	status = name1 != name2
	fmt.Println("Is Not Equals :", status)
}
