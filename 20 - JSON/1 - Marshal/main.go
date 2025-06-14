package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorro := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(cachorro)

	cachorroEmJson, erro := json.Marshal(cachorro)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(bytes.NewBuffer(cachorroEmJson))
}
