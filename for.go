package main

import "fmt"

// no go, não existe while ou dowhile, apenas for

func main(){
	//for horas := 0; horas <= 12; horas++ {
	//	for minutos := 0; minutos <= 60; minutos++ {
	//		for segundos := 0; segundos <= 60; segundos++{
	//			fmt.Printf("%v:%v:%v\n", horas, minutos, segundos);
	//		}
	//	}
	//}
	//booleano := true;
	//for booleano == true {
	//	fmt.Println("Eh assim que se faz um while em go");
	//	if(booleano == true){
	//		fmt.Println("E assim, o break");
	//		break;
	//	}
	//}
	//for i := 0; i <= 20; i++{
	//	if(i % 2 != 0){
	//		continue; // continue: "pula uma interação"
	//	}
	//	fmt.Println(i);
	//}
	for i := 33; i <= 122; i++{
		fmt.Printf("%c\n", i);
	}
}