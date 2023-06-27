package _range

import (
	"fmt"
)

func ForRangeMap() {
	dictionary := map[string]string{
		"book":  "kitap",
		"table": "masa",
		"glass": "bardak",
	}

	for k, v := range dictionary {
		fmt.Println(k, " = ", v)
	}
}
