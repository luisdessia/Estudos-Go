//Array

	/*
		1- Tamanho fixo, de zoro ou mais elementos do mesmo tipo;
		
		2- Acessamos os valores com índice: a[0], a[1]...;
		
		3 - Função embutida len() retorna tamanho do array;

		4 - Por conta do tamanho fixo, não é tão usado. Usado somente em casos específicos.

	*/


//Slice

	/*
		1- Tipo o array, mas sem tamanho fixo;

		2- Acessamos os valores com índice a[0], a[1]

		3- Função embutida len() retorna o tamanho do slice;

		4 - Função append() usada para adicionar valor.

	*/


	package main

	import "fmt"

	func main(){


	//*********Array**********

		var array [2]string

		array[0] = "Hello"
		array[1] = "World"

		fmt.Println(array[0])
		fmt.Println(array[1])
		fmt.Println(array[0], array[1])
		fmt.Println(array)

		numPrimos := [6] int {2, 3, 5, 7, 11, 13}

		fmt.Println(numPrimos)

		//posição 0 até posição 4
		fmt.Println(numPrimos[0:4])

		//Antes da posição 2
		fmt.Println(numPrimos[:2])

		//Após posição 2
		fmt.Println(numPrimos[2:])


	//********************************************//

	//*******Slice*******

		//var slice [] string

		slice := make([] string, 5)

		slice[0] = "Hello"
		slice[1] = "World"
		fmt.Println(slice[0], slice[1])
		fmt.Println(slice)

		//função len determina o tamanho do slice
		fmt.Println(len(slice))

		numPares := []int{2, 4, 6, 8, 10, 12}
		fmt.Println((numPares))

		//função append adiciona um novo elemento
		//pode usar quantos elemento quiser num mesmo comando (append(numPares, 14, 16, 18, 20 ...))
		numPares = append(numPares, 14)
		fmt.Println(numPares)
	}