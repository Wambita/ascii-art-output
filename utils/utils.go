package utils

import (
	"fmt"
	"os"
	"strings"
)

/*
* ValidateArgs: assigns arguments appropriately depending on length of os.Args
*/
func ValidateArgs(args []string) (string, string, string) {
	var shouldSave bool
	var userInput string
	var flag string
	bannerfile := "standard"

	// usage: go run . --output=right something standard
	if len(args) == 3 {
		if flag, shouldSave = CheckFlag(args[0]); shouldSave {
			userInput = args[1]
			bannerfile = args[2]
		} else {
			PrintErrorAndExit()
		}

		// usage: go run . --output=right something
	} else if len(args) == 2 {
		if flag, shouldSave = CheckFlag(args[0]); shouldSave {
			userInput = args[1]

			// usage: go run . something standard
		} else {
			userInput = args[0]
			if ValidBanner(args[1]) {
				bannerfile = args[1]
			} else {
				PrintErrorAndExit()
			}
		}

		// usage: go run . something
	} else if len(args) == 1 {
		userInput = args[0]
		if strings.HasPrefix(userInput, "--output=") {
			PrintErrorAndExit()
		}
	} else {
		PrintErrorAndExit()
	}

	if len(userInput) == 0 {
		PrintErrorAndExit()
	}
	return bannerfile, flag, userInput
}

/*
* CheckFlag: check if the correct flag has been passed
 */
func CheckFlag(input string) (string, bool) {
	if strings.HasPrefix(input, "--output=") {
		filename := strings.TrimPrefix(input, "--output=")
		if !(strings.HasSuffix(filename, ".txt") && len(filename) >= 5) {
			PrintErrorAndExit()
		} else {
			// avoid altering the banner files
			if filename == "standard.txt" || filename == "shadow.txt" || filename == "thinkertoy.txt" {
				PrintErrorAndExit()
			} else {
				return filename, true
			}
		}
	}
	return "", false
}

/*
* PrintErrorAndExit: print and exit program due to usage error
* Prints the error message as is required by the client.
* -[ DO not change the error message ]-
 */
func PrintErrorAndExit() {
	fmt.Println(`
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
	`)
	os.Exit(0)
}

/*
* ValidBanner: check if the correct banner filename has been passed
*/
func ValidBanner(banner string) bool {
	return banner == "standard" || banner == "shadow" || banner == "thinkertoy"
}

/*
* IsValidInput: check if the input string contains unprintable and unsupported characters that are not within the ascii printable range
 */ 
func IsValidInput(input string) (bool, string) {
	NonPrintableChars := []string{"\\a", "\\b", "\\t", "\\v", "\\f", "\\r", "\a", "\b", "\t", "\v", "\f", "\r"}
	for _, char := range NonPrintableChars {
		if contains := strings.Contains(input, char); contains {
			return false, string(char)
		}
	}
	// other characters
	for _, ch := range input {
		if !((ch >= 32 && ch <= 126) || ch == '\n') {
			return false, string(ch)
		}
	}
	return true, input
}
