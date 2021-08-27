//package main

import "fmt"

func main() {
	slice := make([]int, 5, 10) // cria um slice de tamanho 5 e capacidade 10
	slice[0], slice[1], slice[2], slice[3], _ = 1,2,3,4,5;
	slice[4] = 10;
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	slice = append(slice, 10)
	fmt.Printf("%v %v %v", slice, len(slice), cap(slice)); // ver o que acontece quando len > cap
}