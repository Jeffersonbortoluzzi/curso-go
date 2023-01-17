
package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("Herança")

	p1 := pessoa{"Jeff", "Bortoluzzi", 23, 196}
	fmt.Println(p1)
	
	estudante1 := estudante{p1, "Engenharia Elétrica", "Anhanguera"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)
	fmt.Println(estudante1.altura)

}
