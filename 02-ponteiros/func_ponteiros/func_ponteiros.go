package main

import (
	"fmt"
)

/*
Ponteiros: Um ponteiro é uma variável que ao invés de armazenar um valor,
armazena um endereço de memória
*/

func zeroValue(i int) {
	i = 0
	fmt.Println("End de memória do i dentro da func:", &i)
}

func zeroPointer(i *int) {
	*i = 0
}

func main() {

	//Memória -> essa memória tem um endereço -> esse endereço guarda um valor

	i := 1

	fmt.Println("Valor inicial:", i)
	fmt.Println("valor do end de memória:", &i)

	zeroValue(i)
	fmt.Println("zeroval:", i)


	zeroPointer(&i)
	fmt.Println("zeroptr", i)
	fmt.Println("pointer:", &i)

}
