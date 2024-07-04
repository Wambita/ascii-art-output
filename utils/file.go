package utils

import (
	"log"
	"os"
)

/*
*	WriteToFile: responsible for writing the string data it is given into the file with the given name
*	Will create the file if it doesnt exist.
 */
func WriteToFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}
