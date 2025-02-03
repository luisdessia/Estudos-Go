package main

import "fmt"


func main(){

	posicao := 3

	switch posicao{
	case 1:
		fmt.Println("Primeiro Lugar")
	case 2:
		fmt.Println("Segundo Lugar")
	case 3:
		fmt.Println("Terceiro Lugar")
	}

	nome := "bob"
	switch nome{
	case "steph":
		fmt.Println("É uma pessoa")
	case "bento":
		fmt.Println("É um cachorro")
	case "bob":
		fmt.Println("É um personagem")	
	}

}