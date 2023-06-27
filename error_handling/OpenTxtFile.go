package error_handling

import (
	"log"
	"os"
)

func OpenTxtFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {

		// Type assertion control
		if pErr, ok := err.(*os.PathError); ok {
			log.Fatalln("File Was NOT Found !! ->", pErr.Path)
		} else {
			log.Fatalln("Has An Error")
			return
		}

	} else {
		log.Println("File Was Found ->", file.Name())
	}
}
