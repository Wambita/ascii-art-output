package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	printArt "ascii-art-output/print"
	"ascii-art-output/utils"
)

/*
* main: composes the different functionality into a workable solution
* + fetch the command line arguments,
* + validate the arguments and get filepath for the banner file,
* + create map from banner file,
* + output content accordingly and display results to user
 */
func main() {
	args := os.Args[1:]
	bannerfile, flag, input := utils.ValidateArgs(args)
	if validInput, offendingCharacter := utils.IsValidInput(input); !validInput {
		log.Fatalf("Error: input contains unallowed character: %q\n", offendingCharacter)
	}

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

	_, hasOutputFile := utils.CheckFlag(args[0])
	if hasOutputFile {
		utils.WriteToFile(flag, finalresult)
	} else {
		fmt.Println(finalresult)
	}
}
