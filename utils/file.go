package utils

import (
	"log"
	"os"
)

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
