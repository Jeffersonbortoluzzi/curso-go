package main

import "fmt"

func main() {
	fmt.Println("Executando a função main")

}

func init() { // executa primeiro que a função main
	fmt.Println("Executando a função init")
}
