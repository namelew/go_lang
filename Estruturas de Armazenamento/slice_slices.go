//package main

import "fmt"

func main() {
	// slices são "infinitos" e arrays possuem tamanho "n" e não pode ser alterado
	slice := []int{1,2,3,4,5};
	slice = append(slice, 12);
	slice_slice := slice[2:3]; // mesma ideia do python
	fmt.Println(slice_slice);
	slice = append(slice[:2], slice[3:] ...); // deletando um resgistro
	fmt.Println(slice);
}