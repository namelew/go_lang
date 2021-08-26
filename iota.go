package main

import "fmt"

const(
	x = iota; // iota is an untype integer i++
	y = iota;
	z = iota;
)

func main(){
	fmt.Println(x, y, z)
}