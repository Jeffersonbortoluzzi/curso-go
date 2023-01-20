package main

import "fmt"

func main() {
	variavel1 := 10
	variavel2 := variavel1

	variavel2++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int // (*) guarda o endereço de memória de um valor inteiro

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro)
	fmt.Println(variavel3, ponteiro)
}
