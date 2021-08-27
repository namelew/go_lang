//package main

import "fmt"

var x[10] int

func main() {
	// array in Go != arrays in C
	x[0] = 20
	x[1] = 1
	fmt.Printf("%T\n", x) // int[n] != int[n+1]
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(len(x))
}