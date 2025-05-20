package main

import "fmt"

func main() {
	//CONCORRÊNCIA != PARALELISMO
	go escrever("Olá mundo")       //essa função vai ficar rodando infinitamente
	escrever("Programando em Go!") //essa função vai ficar rodando infinitamente
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
	}
}
