package main

import "fmt"

func main() {
	x := 10
	y := func(x int) int {
		//fmt.Println(x,"vezes 80 é", x*80)
		return x*80;
	}
	fmt.Println(x,"vezes 80 é", y(x))
	
}