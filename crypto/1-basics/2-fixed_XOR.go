package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

//Fixed XOR
//Write a function that takes two equal-length buffers and produces their XOR combination.

// If your function works properly, then when you feed it the string:

// 1c0111001f010100061a024b53535009181c
// ... after hex decoding, and when XOR'd against:

// 686974207468652062756c6c277320657965
// ... should produce:

// 746865206b696420646f6e277420706c6179

func fixed_xor(a string, b string) string {
	// hex decode both input strings into slice of bytes
	decoded1, err := hex.DecodeString(a)
	if err != nil {
		log.Fatal(err)
	}
	decoded2, err := hex.DecodeString(b)
	if err != nil {
		log.Fatal(err)
	}

	// make result slice to catch results of the XORs
	result := make([]byte, 0, len(a))
	// loop over the decoded slices performing XOR and  appending to result
	for k, v := range decoded1 {
		result = append(result, (v ^ decoded2[k]))

	}

	// encode the byte slice to hex and return
	return hex.EncodeToString(result)
}

func main() {
	fmt.Println(fixed_xor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
}
