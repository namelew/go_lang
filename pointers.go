package main

import "fmt"

func main() {
	x := 10
	y := &x // mesma coisa de c

	fmt.Println("Endereço de Memória:", y)
	fmt.Println("Valor armazenado:", *y)
	fmt.Printf("%T %T\n", x, y)
}