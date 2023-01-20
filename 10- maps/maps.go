package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Jeff",
		"sobrenome": "Bortoluzzi",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])
	fmt.Println(usuario["sobrenome"])

	fmt.Println("--------------------")

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Silvia",
			"ultimo":   "Satiko",
		},
		"curso": {
			"nome":   "Enfermagem",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["nome do pai"] = map[string]string{
		"nome": "Silvio",
	}
	fmt.Println(usuario2)
}
