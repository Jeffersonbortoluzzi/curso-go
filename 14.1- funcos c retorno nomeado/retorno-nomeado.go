package main

import "fmt"

func calcularMatematicos(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

// func main() {
// 	fmt.Println(calcularMatematicos(10, 22))
// }

//preferivel escrever assim
func main() {
	soma, subtracao := calcularMatematicos(10, 20)
	fmt.Println(soma, subtracao)

}
