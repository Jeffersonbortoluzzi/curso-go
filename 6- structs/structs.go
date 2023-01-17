package main

import "fmt"

type usuario struct {
	nome     string
	idade    int
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Jeff"
	u.idade = 23
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua Teste", 001}

	usuario2 := usuario{"Silvia", 22, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 23} //para passar somente 1 parâmetro
	fmt.Println(usuario3)
}
