package main

import "fmt"

// MÉTODOS

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "Andou")
}

func main() {

	carro := Carro{
		Nome: "Kadett GSI",
	}

	carro.andar()

	resultado := func(x ...int) func() int {

		res := 0

		for _, valor := range x {
			res += valor
		}
		return func() int {
			return res * res
		}
	}

	fmt.Println(resultado(11, 12, 13, 10)())

}

// valor e o tipo do valor  // retorno

func soma(a int, b int) (result int) {
	result = a + b
	return
}

// funcoes variadicas

func somaTudo(x ...int) int {
	resultado := 0

	for _, valor := range x {
		resultado += valor
	}
	return resultado
}
