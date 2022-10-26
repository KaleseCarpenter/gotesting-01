package main

import (
	"fmt"
)

// Will take in one number and return a result
func Calculate(x int) (result int) {
	result = x + 2
	return result
}

func Multiply(z int) (result int) {
	result = z * 3
	return result
}

func main() {
	fmt.Println("Writing a Golang Test.", "Yeee!")
	//result := Calculate(2)
	//fmt.Println(result)
	result := Multiply(3)
	fmt.Println(result)

	// new addition and multiplication code uzing interface
	m := multiplication{66, 13}
	a := addition{32, 8}
	printResult(m)
	printResult(a)
}

/*
	Any struct defining a receiver function with exact same signature as that of operate,

will be an honorary member of type operation
*/
type operation interface {
	operate() int
}

// create a struct to use as an interface with the different math functions
type multiplication struct {
	k int
	c int
}

type addition struct {
	x int
	y int
}

func (m multiplication) operate() int {
	return m.k * m.c
}

func (a addition) operate() int {
	return a.x + a.y
}

func printResult(o operation) {
	fmt.Println(o.operate())
}

// func printResult(m multiplication) {
// 	fmt.Println(m.operate())
// }

// func printResult(a addition) {
// 	fmt.Println(a.operate())
// }
