package main

import "fmt"


//IF = SE
//ELSE = SE NÃO

func main(){

	numero := 3

	//if (condição) {( <ação) }
	if numero == 1 {//true

		fmt.Println("Valor é igual a 1")		
	}else{
		fmt.Println("O valor não é igual a 1")
	}

	if numero ==1 {//true
		fmt.Println("Valor é igual a 1")		
	}else if numero == 2 {
		fmt.Println("Valor é igual a 2")		
	}else{
		fmt.Println("Valor é diferente de 1 e 2")
	}


	numero = 8
	if numero % 2 == 0 {
		fmt.Printf("%d é par", numero)				
	}else{
		fmt.Printf("%d é impar", numero)
	}
	
	//operações

	/*
	fmt.Println(1+1)
	fmt.Println(2-1)
	fmt.Println(2 * 2)
	fmt.Println(10 / 2)
	fmt.Println(10 % 2)
	*/



}