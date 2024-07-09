package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
* CheckFlag: check if the correct flag has been passed
 */
func CheckFlag() (string, string, string, string) {
	// Parse flags manually
	output, align := "", ""
	args := os.Args[1:]
	var hasOutput, hasAlign bool
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if strings.HasPrefix(arg, "--") {
			parts := strings.SplitN(arg[2:], "=", 2)
			if len(parts) != 2 {
				PrintErrorAndExit()
			}
			switch parts[0] {
			case "output":
				if hasOutput {
					PrintErrorAndExit()
				} else {
					output = parts[1]
					hasOutput = true
				}
			case "align":
				if hasAlign {
					PrintErrorAndExit()
				} else {
					align = parts[1]
					hasAlign = true
				}
			default:
				PrintErrorAndExit()
			}
		} else {
			// Non-flag arguments
			args = args[i:]
			break
		}
	}

	// Validate arguments
	if len(args) < 1 || len(args) > 2 {
		PrintErrorAndExit()
	}
	text := args[0]
	bannerfile := ""
	if len(args) == 2 {
		bannerfile = args[1]
	}

	// Validate align if provided
	if align != "" && align != "left" && align != "right" && align != "center" && align != "justify" {
		PrintErrorAndExit()
	}

	// Validate output if provided
	if output != "" {
		if !strings.HasSuffix(output, ".txt") || len(output) < 5 {
			PrintErrorAndExit()
		}
		// Avoid altering the banner files
		if (output == "standard.txt" || output == "shadow.txt" || output == "thinkertoy.txt") ||
			(strings.HasSuffix(output, "/standard.txt") || strings.HasSuffix(output, "/shadow.txt") || strings.HasSuffix(output, "/thinkertoy.txt")) {
			PrintErrorAndExit()
		}
	}

	return align, output, text, bannerfile
}

/*
* PrintErrorAndExit: print and exit program due to usage error
* Prints the error message as is required by the client.
* -[ DO not change the error message ]-
 */
func PrintErrorAndExit() {
	fmt.Println(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)
	os.Exit(0)
}

// function to get the terminal width instead of using one fixed width
func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	sizeString := string(output)
	sizestring := strings.Split(sizeString, " ")

	size, err := strconv.Atoi(strings.Trim(sizestring[1], "\n"))
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return size
}

// function to get spaces to add for alignment depending on the alignment flag
func GetSpaces(flag string, asciiString string) int {
	terminalWidth := GetTerminalWidth()

	spaces := 0
	switch flag {
	case "right":
		spaces = terminalWidth - len(asciiString)
	case "left":
		spaces = 0
	case "center":
		spaces = (terminalWidth - len(asciiString)) / 2
	}
	return spaces
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
