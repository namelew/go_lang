package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

type profissional struct{
	pessoa
	profissao string
	salario float64
}

func main() {
	pessoa1 := pessoa{"João",16}
	trabalhador := profissional{pessoa{"Carlos", 18}, "Artista", 1500}
	fmt.Println(pessoa1.nome, trabalhador.nome)

}