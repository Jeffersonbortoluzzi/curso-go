
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// converter struct para json
type cachorro struct {
	Nome  string `json:"nome`
	Raca  string `json:"raca`
	idade uint   `json:"idade`
}

func main() {
	c := cachorro{"Rex", "Golden", 2}
	fmt.Println(c)

	cachorroEmJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJson)
	fmt.Println(bytes.NewBuffer(cachorroEmJson))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorro2EmJson, erro := json.Marshal(c2)
	fmt.Println(bytes.NewBuffer(cachorro2EmJson))
}
