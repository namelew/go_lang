//package main

import "fmt"

type new_type int
var x new_type;
func main(){
	x = 10;
	fmt.Printf("x: %v %T", x, x)
}