package main

import (
	"fmt" 
)

func main() {

	fmt.Println("strings " + "work " + "as " + "expected")
	fmt.Println("so do ints", 1 + 1 , 5 + 5)
	fmt.Println("there are also floats", 9001 / 13.37 )

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}