package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Given a list of strings, return a list with the strings in sorted order, except group all the strings that
// begin with 'x' first. e.g. ['mix', 'xyz', 'apple', 'xanadu', 'aardvark']
// yields ['xanadu', 'xyz', 'aardvark', 'apple', 'mix']

func main() {
	// make 2 lists one with x words, one without x words
	// sort both lists
	// combine both lists and print
	args := os.Args[1:]

	no_x_slice := make([]string, 0)
	x_slice := make([]string, 0)

	for _, word := range args {
		if string(word[0]) == "x" {
			x_slice = append(x_slice, word)
		} else {
			no_x_slice = append(no_x_slice, word)
		}
	}

	sort.Strings(no_x_slice)
	sort.Strings(x_slice)
	result := append(x_slice, no_x_slice...)
	formatted := strings.Join(result, `, `)
	fmt.Print(`[`, `'`, formatted, `'`, `]`)
}
