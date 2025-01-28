package main

import "fmt"

func main(){

	fmt.Println(subtracao(10, 5))
	
	//em go dá para atribuir uma função a uma variável
	Valor_subtraido := subtracao(20, 12)

	fmt.Println(Valor_subtraido)

}


func subtracao (x , y int) int {

	return x - y
}