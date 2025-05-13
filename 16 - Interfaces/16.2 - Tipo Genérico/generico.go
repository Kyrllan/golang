package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
}

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
}
