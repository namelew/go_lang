//package main

import "fmt"

func main() {
	slice := []int{1,2,3,4,5}
	for indice, valor := range slice{
		fmt.Printf("slice[%v] = %v\n", indice, valor);
	}
	fmt.Println();
	slice  = append(slice, 100)

	for indice, valor := range slice{
		fmt.Printf("slice[%v] = %v\n", indice, valor);
	}
}