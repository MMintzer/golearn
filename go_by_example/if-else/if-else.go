package main

import "fmt"

func main() {
	matt := "learning go"

	if matt == "learning go" {
		fmt.Println("Matt is " + matt)
	} else {
		fmt.Println("Matt is not " + matt)
	}
}
