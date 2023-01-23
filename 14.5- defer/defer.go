package main

import "fmt"

func função1() {
	fmt.Println("Executando a função 1")
}

func função2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) string {
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para ver se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 7 {
		return "aprovado"
	}

	return "reprovado"
}

func main() {
	// defer função1() // adiar a execução até o ultimo momento possivel (até terminar)
	// função2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
