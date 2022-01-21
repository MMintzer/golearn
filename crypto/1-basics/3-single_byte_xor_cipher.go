package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

/*
Single-byte XOR cipher
The hex encoded string:

1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric.
Evaluate each output and choose the one with the best score.
*/

func decode_hex(s string) []byte {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return decoded
}

// check for common english character
func good_char(s string) bool {
	common_chars := []string{"a", "e", "i", "o", "u", "r", "s", "t", "l", "n"}
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

	for i := 0; i <= 100; i++ {
		score := score_text(fixed_xor(s, i))
		if score > high_score {
			high_score = score
			best_result = fixed_xor(s, i)
		}
	}

	return best_result
}

func main() {
	fmt.Println(single_byte_XOR_cipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
}
