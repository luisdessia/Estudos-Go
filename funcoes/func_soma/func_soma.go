package main

import "fmt"

func main(){

	fmt.Println(soma(42, 13))

	//em go dá para atribuir uma função a uma variável
	Valor_somado :=  soma(10, 12)

	//uma maneira de printar
	fmt.Printf("%v\n", Valor_somado)

	//outra maneira de printar
	fmt.Println(Valor_somado)
}


func soma (x int, y int) int {
	
	return x + y
}