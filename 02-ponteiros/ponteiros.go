package main

import "fmt"

/*
Ponteiros: Um ponteiro é uma variável que ao invés de armazenar um valor,
armazena um endereço de memória
*/


func main (){

	//Memória -> essa memória tem um endereço -> esse endereço guarda um valor

	i := 1

	fmt.Println("Valor inicial", i)
	fmt.Println("valor do end de memória:", &i)//&var -> aponta para o end de memória

	a :=&i

	fmt.Println("Variavel a:", a)
	fmt.Println("Variavel *a", *a)//quando uso *, estamos falando do endereço e não do valor
	fmt.Println("Variavel &a", &a)

}