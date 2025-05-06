package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo o arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("kyrllan@hotmail.com")
	fmt.Println(erro)
}
