package main

import (
	"fmt"
	//"time"
)

//loops
//Laços de Repetição
//Repetir Tarefas

func main(){

    sum := 0

    //i ++ -> soma 1
    //i -- -> subtrai 1
    for i := 0; i < 10; i++ {
        
        sum += i
        fmt.Printf("sum = %d, indice = %d\n", sum, i)
        //é a mesma coisa que: sum = sum +1
        //sum -= -> sum = sum -1
    }

    fmt.Println(sum)


    //****loop infinito****

   /* for {
        fmt.Println("Golang do zero")
    }*/
    

    //****loop infinito**** controlado

   /* for {
        fmt.Println("Golang do zero")
        time.Sleep(2 * time.Second)
    }*/


    //****for com range****

    frutas := [] string{"Laranja", "Maça", "Banana", "Uva", "Kiwi"}

        //_, exclui a necessidade do uso do i (indice)
    for _, fruta := range frutas  {
        
        fmt.Println(fruta)
    }
}