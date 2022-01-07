package main

// Given a list of strings, return the count of the number of strings where the string
// length is 2 or more and the first and last chars of the string are the same.

import (
	"fmt"
	"os"
)

func main() {
	// get args as list
	args := os.Args
	// instatiate a count variable as an int
	count := 0

	// loop through args testing for length of 2 and first and last values are ==
	for _, arg := range args {
		if len(arg) >= 2 && arg[0] == arg[len(arg)-1] {
			count++
		}
	}

	// print count
	fmt.Println(count)
}
