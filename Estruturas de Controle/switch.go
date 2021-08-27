package main

import "fmt"

func main() {
	x:= 0
	switch x {
		case 1:
			fmt.Println("1");
			fallthrough; // pula pra próxima comparação
		case 2:
			fmt.Println("2")
		case 3: 
			fmt.Println("3")
		default:
			fmt.Println("Nao sei")
	}
}

