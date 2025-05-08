package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	variavel1++

	fmt.Println(variavel1, variavel2)

	// Ponteiro é uma referência de memória
	var variavel3 int = 100
	var ponteiro *int = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //Dereferenciando o ponteiro
	variavel3 += 50
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro) //Dereferenciando o ponteiro
}
