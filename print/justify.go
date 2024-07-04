package printArt

import (
	"strings"

	"ascii-art-output/utils"
)

func Justify(input string, asciiMap map[rune][]string) string {
	result := ""
	splitWord := strings.Split(input, " ")
	// get the number of gaps
	numGaps := len(splitWord) - 1
	// get the combined width of the words as ascii
	combinedWidth := 0
	for _, word := range splitWord {
		for _, char := range word {
			combinedWidth += len(asciiMap[char][0])
		}
	}
	// get terminal width
	terminalWidth := utils.GetTerminalWidth()
	if combinedWidth > terminalWidth || numGaps == 0 {
		result = Normal(input, asciiMap)
		return result
	}
	// calculate the total number of spaces
	totalSpaces := terminalWidth - combinedWidth
	// divide the spaces equally depending on the number of words -1 for cat -e
	spacesBetween := totalSpaces / numGaps
	extraSpaces := totalSpaces % numGaps
	// print
	for i := 0; i < 8; i++ {
		lineOutput := ""
		for j, word := range splitWord {
			for _, char := range word {
				lineOutput += asciiMap[char][i]
			}
			if j < numGaps {
				lineOutput += strings.Repeat(" ", spacesBetween)
				if j < extraSpaces {
					lineOutput += " "
				}
			}
		}

		result += lineOutput + "\n"
	}
	return result
}
