package main

import "fmt"

func main(){
	y := 24
	x := y << 2; // deslocamento para a esquerda
	z := y >> 4; // deslocamento para a direita
	fmt.Printf("%b\n", x)
	fmt.Printf("%b", z)

}