package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	variavel3, variavel4 := "Variavel 3", "Variavel 4"
	fmt.Println(variavel3, variavel4)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel4, variavel2 = variavel3, variavel1
	fmt.Println(variavel3, variavel1)
}
