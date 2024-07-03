package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	printArt "ascii-art-justify/print"
	"ascii-art-justify/utils"
)

// fetch the command line arguments,
// validate the arguments and get filepath for the banner file,
// create map from banner file,
// align content accordingly and display results to user
func main() {
	args := os.Args[1:]
	bannerfile, flag, input := utils.ValidateArgs(args)
	fmt.Println(flag)
	if validInput, offendingCharacter := utils.IsValidInput(input); !validInput {
		log.Fatalf("Error: input contains unallowed character: %q\n", offendingCharacter)
	}
	// filename, _ := utils.CheckFlag(input)

	asciiMap := utils.CreateMap(bannerfile)
	if asciiMap == nil {
		os.Exit(1)
	}
	data := strings.ReplaceAll(input, "\\n", "\n")
	words := strings.Split(data, "\n")
	finalresult := ""
	for _, word := range words {
		result := printArt.Normal(word, asciiMap)
		finalresult += result
	}
	utils.WriteToFile(flag, finalresult)
}
