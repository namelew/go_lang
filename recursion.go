package main

import "fmt"

func main() {
	fmt.Println(fatorial(5))
}

func fatorial(x int) int{
	if(x == 0){
		return 1
	}
	return x*fatorial(x-1)
}