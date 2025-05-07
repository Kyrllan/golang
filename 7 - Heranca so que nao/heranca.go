package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
}

type estudante struct {
	//pessoa pessoa //declara sem o tipo, senao vai criar um campo chamado pessoa com todos os campos dentro dele
	pessoa
	curso     string
	faculdade string
}

func main() {
	pessoa1 := pessoa{"João", "Silva", 20}
	fmt.Println(pessoa1)
	estudante1 := estudante{pessoa{"João", "Silva", 20}, "Engenharia", "UTFPR"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)
}
