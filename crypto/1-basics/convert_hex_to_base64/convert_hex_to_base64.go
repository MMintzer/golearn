package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
)

// Convert hex to base64
// 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
// Should produce:
// SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
// So go ahead and make that happen. You'll need to use this code for the rest of the exercises.

// Cryptopals Rule
// Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.

// bless copy and paste from docs, I guess
// https://pkg.go.dev/encoding/hex

// this helped for encoding to base64 https://gobyexample.com/base64-encoding
func main() {
	const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)
	sEnc := b64.StdEncoding.EncodeToString([]byte(decoded))
	fmt.Println(sEnc)

}
