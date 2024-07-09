# test data files

## description

This directory contains files needed to test the creation of the artistic representation map. Each tests for a case in which our program should fail or succeed during the creation process. 
- empty.txt represents the case where there is an empty banner file
- few characters.txt represents the case where there are fewer lines. Note that omitting a single character on a line simply represents a redesign and not a failure case. This has the implication that there might be fewer than the required 95 characters.
- manycharacters.txt represents the case where one creates a banner file that does not meet the maximum line requirements. This has the implication that there might be more than the 95 printable characters in the map.
- standard.txt is the perfect exmaple of a banner file.

It is important to note that we test for the creation functionality, because once the map is created, we utilize go functions to print on screen. This do not need to be tested any further.

**used in** these files are used in utils/create_map_test.go file