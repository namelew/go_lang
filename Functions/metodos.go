//package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

func (p pessoa) oi(){
	fmt.Printf("Oi meu nome é %v e tenho %v anos\n", p.nome, p.idade)
}

func main() {
	p := pessoa{"Diogo", 19}
	p.oi()
	//oi() //não funciona
	//oi(p) // também não funciona
}