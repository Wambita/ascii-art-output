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
	align, output, input, bannerfile := utils.CheckFlag()

	if validInput, offendingCharacter := utils.IsValidInput(input); !validInput {
		log.Fatalf("Error: input contains unallowed character: %q\n", offendingCharacter)
	}
	if bannerfile == "" {
		bannerfile = "standard"
	} else if !utils.ValidBanner(bannerfile) {
		utils.PrintErrorAndExit()
	}
	asciiMap := utils.CreateMap(bannerfile)
	if asciiMap == nil {
		os.Exit(1)
	}

	data := strings.ReplaceAll(input, "\\n", "\n")
	words := strings.Split(data, "\n")
	finalresult := ""

	for _, word := range words {
		var result string
		switch align {
		case "justify":
			result = printArt.Justify(word, asciiMap)
		case "center", "right":
			result = printArt.Align(word, align, asciiMap)
		default:
			result = printArt.Normal(word, asciiMap)
		}
		finalresult += result
	}
	fmt.Print(finalresult)

	fmt.Println(output)
	
	if output !="" {
		utils.WriteToFile(output, finalresult)
	} else {
		fmt.Print(finalresult)
	}
}
