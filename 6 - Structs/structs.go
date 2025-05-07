package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	var usuario1 usuario
	usuario1.nome = "João"
	usuario1.idade = 20
	fmt.Println(usuario1)
	usuario2 := usuario{"João", 20, endereco{"Rua dos bobos", 0}}
	fmt.Println(usuario2)
}
