package main

import "fmt"

func main() {

	// float32   IEEE-754
	var foo32 float32 = 2147483647
	fmt.Println(foo32)

	// float64   IEEE-754
	var foo64 float64 = 9223372036854775807
	fmt.Println(foo64)
}
