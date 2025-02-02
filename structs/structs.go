/*
	************** Structs *************

	Forma de criar sua própria estrutura de dados
	Personalizar de acordo com a sua necessidade
	Podemos udar vários tipos diferentes
*/

package main

import "fmt"

//type <nome da estrutura< struc { <campos> }
type Pessoa struct{

	Nome string
	Idade int
}

func main(){

	fmt.Println(Pessoa{"Steph", 28})
	fmt.Println(Pessoa{Nome: "Kronos", Idade: 4})
	fmt.Println(Pessoa{Nome: "Luis"})


	p1 := Pessoa {Nome: "Bob", Idade: 2}

	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	p1.Idade = 3

	fmt.Println(p1.Idade)

	p2 := Pessoa {Nome: "Patrick", Idade: 2}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)

	fmt.Println(pessoas)

	// struct + map

	//alunos := map[string][]Pessoa{}
	//alunos["Programação"] = pessoas
	//fmt.Println(alunos)


	var alunos = map[string] []Pessoa{

		"Programção": {{Nome: "Steph", Idade: 28  }, {Nome: "Bento", Idade: 5}},
		"Engenharia": {{Nome: "Steph2", Idade: 28}, {Nome: "Bento 2", Idade : 5}},
	}

		fmt.Println(alunos)


	//Vamos utilizar outra struct junto, ela será decalrada após a main apenas para estruturar mehlor o estudo, declarações sempre acima da main


	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)

}

type Profissao struct {

	Pessoa
	Tipo string
}

