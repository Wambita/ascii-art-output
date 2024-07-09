package utils

import (
	"fmt"
	"os"
	"testing"
)

/*
* Test WriteToFile in file.go
* Writes a bit of data to it and check its stats
* if the size changes, truncate the file and exit
 */
func TestWriteFile(t *testing.T) {
	filename := "writefiletester.txt"

	WriteToFile(filename, "A bit of data")

	stats, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Please create the file with name %q in the current directory", filename)
		return
	}

	if stats.Size() == 0 {
		t.Errorf("Expected to have data in file after invoking the WriteToFile function, but it did not happen")
	}

	os.Truncate(filename, 0)
}
