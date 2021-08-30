//package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade int
}

type dentista struct{
	pessoa
	esp string
	salario float64
}

type arquiteto struct{
	pessoa
	tipo_predio string
	salario float64
}

func (x pessoa) oi(){
	fmt.Println("Oi, bom dia. Meu nome é", x.nome)
}

func (x dentista) oi(){
	fmt.Println("Oi, bom dia. Meu nome é", x.nome)
}

func (x arquiteto) oi(){
	fmt.Println("Oi, bom dia. Meu nome é", x.nome)
}

type gente interface {
	oi()
}

func serhumano(g gente){
	g.oi()
	switch g.(type) {
		case dentista:
			fmt.Printf("Especialização: %v\n", g.(dentista).esp)
			fmt.Printf("Salario: %v\n", g.(dentista).salario)
		case arquiteto:
			fmt.Printf("Tipo de Prédio: %v\n", g.(arquiteto).tipo_predio)
			fmt.Printf("Salário: %v\n", g.(arquiteto).salario)
	}
}

func main() {
	d := dentista{pessoa{"Wagner", "Pereira", 28},"aparelhos",1500.50}
	a := arquiteto{pessoa{"Kleiton", "da Silva", 45},"Construção Civil", 8000}
	//d.oi() 
	//a.oi()
	serhumano(d) // interface permiti isso
	serhumano(a)
}