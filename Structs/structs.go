//package main

import "fmt"

type cliente struct {
	nome string
	sobrenome string
	fuma bool
}

func main(){
	cliente1 := cliente{
		nome: "Pedro",
		sobrenome: "Silva",
		fuma: false,
	}
	cliente2 := cliente{"Joana", "Ferreira", true}
	fmt.Println(cliente1, cliente2)

}