package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
Detect single-character XOR
One of the 60-character strings in this file has been encrypted by single-character XOR.

Find it.

(Your code from #3 should help.)

*/
// error checkor function from https://gobyexample.com/reading-files
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// decodes hex string returns byte slice
func decode_hex(s string) []byte {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return decoded
}

// check for common english character (adding the " " character was necessary to get this to work  Makes sense)
func good_char(s string) bool {
	common_chars := []string{"a", "e", "i", "o", "u", "r", "s", "t", "l", "n", " "}
	for i := 0; i < len(common_chars); i++ {
		if strings.ToLower(s) == common_chars[i] {
			return true
		}
	}
	return false
}

// take score of a text using common english characters
func score_text(s string) int {
	score := 0
	for i := 0; i < len(s); i++ {
		if good_char(string(s[i])) {
			score++
		}
	}
	return score
}

func fixed_xor(a string, b int) string {
	decoded := decode_hex(a)

	result := make([]byte, 0, len(a))
	for _, v := range decoded {
		result = append(result, (v ^ byte(b)))

	}
	return string(result)
}

// brute force XOR and keep track of best guess for decrypted message
func single_byte_XOR_cipher(s string) string {
	high_score := 0
	best_result := ""

	for i := 0; i <= 255; i++ {
		score := score_text(fixed_xor(s, i))
		if score > high_score {
			high_score = score
			best_result = fixed_xor(s, i)
		}
	}

	return best_result
}

// returns the lines of a file as a slice of strings
func lines_in_file(filename string) []string {
	data, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(data)
	result := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result
}

func main() {
	// 1. open file and read in (hopefully as a slice of strings)
	lines := lines_in_file("/Volumes/Mac/golearn/crypto/1-basics/4.txt")

	// 2. Run my single byte XOR decoder against each slice picking the best choice and
	// appending to my potential answer lists
	answers := []string{}
	for _, v := range lines {
		answers = append(answers, single_byte_XOR_cipher(v))
	}

	// 3. run my word scorer against each line picking the best
	high_score := 0
	best_result := ""
	for _, v := range answers {
		if score_text(v) > high_score {
			high_score = score_text(v)
			best_result = v
			fmt.Println(high_score, best_result)
		}
	}

	fmt.Println(best_result)
}
