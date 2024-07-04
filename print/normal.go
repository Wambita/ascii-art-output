package printArt

/*
* Normal: Generates the artistic representation of the plaintext
 */
func Normal(input string, asciiMap map[rune][]string) string {
	result := ""
	if input == "" {
		result += "\n"
	} else {
		for i := 0; i < 8; i++ {
			lineOutput := ""
			for _, char := range input {
				lineOutput += asciiMap[char][i]
			}
			result += lineOutput + "\n"
		}
	}
	return result
}
