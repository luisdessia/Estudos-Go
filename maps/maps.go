/*
	Maps: 
			1: Heterogêneos

			2: Pode minsturar tipos
			
			3: Estrutura chave - valor

			4: [key] = value

			5: chave tem um tipo e valor pode ter outro

			6: map[k] -> k = chave, v =valor



			map [string]int
			{"Luis": 21, "Kronos": 4}

			map [string] string
			{"Luis" : "Dessia", "Kronos" : "Dessia"}

*/

package main

import "fmt"


func main(){

	idade := map[string]int{}
	idade["Luis"] = 21
	idade["Kronos"] = 4
	fmt.Println(idade)
	fmt.Println(idade["Luis"])
	fmt.Println(idade["Kronos"])

	//*********///

	anoNasc := map [string] int{

		"Luis": 2003,
		"Kronos": 2020,
	}

	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Luis"])
	fmt.Println(anoNasc["Kronos"])
	anoNasc["Cassia"] = 1980
	fmt.Println(anoNasc)

}