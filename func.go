//package main

import "fmt"

func main() {
	total, q := soma_multipla(0, 1, 3, 4, 15);
	normal := soma(10, 41); 
	basica();
	fmt.Println(normal);
	fmt.Println(total, q);

}

func basica(){
	fmt.Println("Sou uma função");
}

func soma(x,y int) int{
	return x+y;
}

func soma_multipla(x ...int) (int, int){
	soma := 0;
	for _,v := range x{
		soma += v;
	}
	return soma, len(x);
}