package main

import "fmt"

func main() {
	amigos := map[string]int{
		"pedro": 12345,
		"joão": 123221,
		"gabriel": 54321,
	}
	fmt.Println(amigos);
	amigos["ramom"] = 32145
	fmt.Println(amigos);
	delete(amigos, "ramom");
	if sera, ok := amigos["ramom"]; !ok {
		fmt.Println("nao tem!");
	}else{
		fmt.Println(sera);
	}
	// maps não possuem ordem
}