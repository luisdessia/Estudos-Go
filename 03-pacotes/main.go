package main

import (
	//go mod init
	"calculator/03-pacotes/math"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main (){

	//pacote externo
	e := echo.New()
	fmt.Println(e)

	sum := math.Sum(1, 2)
	fmt.Println(sum)

	sub := math.Sub(1, 2)
	fmt.Println(sub)
}



/* Principais commandos GoMod (mais usados)

	go mod init - iniciar o package local
	go mod tidy - packages externos
	go mod download - usar quando o tidy "bugar"
	go mod verify - verificar dependências
	go mo vendor - cria cópia de depêndencias

*/

/*
	******Funções Públicas******

	Primeira Letra Maiúscula




	******Funções Privada******

	Primeira Letra Minúscula

*/