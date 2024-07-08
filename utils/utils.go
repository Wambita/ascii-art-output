package utils

import (
	"flag"
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
	output := flag.String("output", "", "output file name")
	align := flag.String("align", "", "text alignment (left, right, center, justify)")

	// silence flag errors
	tmp := os.Stderr
	os.Stderr = nil

	flag.Usage = func() {
		// PrintErrorAndExit()
		fmt.Fprintf(os.Stdout , "\nUsage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard\n\n")
		os.Exit(0)
	}

	flag.Parse()
	args := flag.Args()

	// Check for correct number of arguments
	if len(args) < 1 || len(args) > 2 {
		PrintErrorAndExit()
	}

	text := args[0]
	bannerfile := ""
	if len(args) == 2 {
		bannerfile = args[1]
	}

	// Validate align if provided
	if *align != "" && *align != "left" && *align != "right" && *align != "center" && *align != "justify" {
		PrintErrorAndExit()
	}

	// go run . --align=left hello
	// go run . --output=any.txt hello
	// => check if flag is '-align'
	if len(os.Args) == 3 {
		if (*align != "" && strings.HasPrefix(os.Args[1], "-align")) || (*output != "" && strings.HasPrefix(os.Args[1], "-output")) {
			PrintErrorAndExit()
		}
	}

	// go run . --align=left --output=any.txt hello [banner]
	// go run . --output=any.txt --align=left hello [banner]
	if len(os.Args) > 3 {
		if (*align != "" && strings.HasPrefix(os.Args[1], "-align")) || (*output != "" && strings.HasPrefix(os.Args[2], "-output")) {
			PrintErrorAndExit()
		}
		if (*output != "" && strings.HasPrefix(os.Args[1], "-output")) || (*align != "" && strings.HasPrefix(os.Args[2], "-align")) {
			PrintErrorAndExit()
		}
	}

	// Validate output if provided
	if *output != "" {
		if !strings.HasSuffix(*output, ".txt") || len(*output) < 5 {
			PrintErrorAndExit()
		}
		// Avoid altering the banner files
		if *output == "standard.txt" || *output == "shadow.txt" || *output == "thinkertoy.txt" {
			PrintErrorAndExit()
		}
	}
	os.Stderr = tmp
	return *align, *output, text, bannerfile
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
