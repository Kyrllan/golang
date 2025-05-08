package main

import "fmt"

func main() {

	usuario := map[string]string{ //Dentro colchetes o tipo da chave e depois o tipo do valor
		"nome":      "João",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
