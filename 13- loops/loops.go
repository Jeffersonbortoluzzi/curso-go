package main

import (
	"fmt"
	// "time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i =", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println("-----------------")

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j =", j)
	// }

	nomes := [...]string{"João", "Davi", "Lucas"}

	//indice do array, valor do array = nomes(joao, davi, lucas)
	for indice, nome := range nomes { 
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	usuario := map[string]string {
		"nome": "Jefferson",
		"sobrenome": "Bortoluzzi",
	}
	
	for chave, valor := range usuario{
		fmt.Println(chave, valor)
	}

}
