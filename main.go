package main

import (
	"fmt"
)

// Will take in one number and return a result
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go Testing Tutorial")
	result := Calculate(2)
	fmt.Println(result)
}
