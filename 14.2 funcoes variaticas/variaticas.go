package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(indice string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(indice, numero)
	}
}

func main() {
	totalDaSoma := soma(31, 25, 53, 42, 15, 64, 7, 85, 79, 170)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 10, 2, 3, 4, 5, 6, 7, 8)
}
