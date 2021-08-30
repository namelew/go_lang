//package main

import "fmt"

func main() {
	fmt.Println(retorna_func()(4))
}

func retorna_func() func (int) int{
	return func(x int) int{
		return x * 2
	}
}