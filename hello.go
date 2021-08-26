package main

import "fmt"

func main() {
	_, erros := fmt.Println("Hello word", "outra frase", 100)
	fmt.Println(erros)
}