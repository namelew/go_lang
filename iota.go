package main

import "fmt"

const(
	x = iota + 1 // iota is an untype integer i++
	y
	z
)

func main(){
	fmt.Println(x, y, z)
}