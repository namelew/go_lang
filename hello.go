//package main // comentando pra evitar erros no próximo programa

import "fmt"

func main() {
	_, erros := fmt.Println("Hello word", "outra frase", 100)
	fmt.Println(erros)
}