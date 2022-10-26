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
	fmt.Println("Writing a Golang Test.", "Yeee!")
	//result := Calculate(2)
	//fmt.Println(result)
	result := Multiply(3)
	fmt.Println(result)
}

func Multiply(k int) (result int) {
	result = k * 3
	return result
}
