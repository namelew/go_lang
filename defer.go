//package main

import "fmt"

func main() {
	defer fmt.Println("Com defer") // defer faz esse código ser sempre a ultima coisa a ser executada
	fmt.Println("Sem defer")
}