package main

import "fmt"

type hot int
var y hot
var x int

func main()  {
	y = 10
	x = 2
	fmt.Printf("x: %v %T\n", x, x)
	x = int(y)
	fmt.Printf("x: %v %T", x, x)	
}