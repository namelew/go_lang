//package main

import "fmt"

func main() {
	// slices são "infinitos" e arrays possuem tamanho "n" e não pode ser alterado
	array := [5]int{1,2,3,4,5}
	fmt.Println(array)
	slice := []int{1,2,3,4,5}
	slice = append(slice, 12)
	fmt.Println(slice)
}