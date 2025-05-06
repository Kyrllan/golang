package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int = 10000000   //int8 int16 int32 int64
	var numero2 uint = 10000000 //uint8 uint16 uint32 uint64 apenas positivos

	//alias
	var numero3 rune = 123456 //rune é um alias para int32
	var numero4 byte = 123    //byte é um alias para uint8

	var numeroReal1 float32 = 123.45 //float32 float64

	//Fim Números

	var string1 string = "Teste1" //tipo string
	string2 := "Teste2"           //tipo implicito

	char1 := "A"[0] //tipo character retorna o valor da tabela ASCII
	char2 := 'B'    //aspas simples com 1 caracter é tipo character retorna o valor da tabela ASCII

	// Fim Strings

	var texto int     //retrorna 0
	var texto2 string //retrorna ""

	var booleano bool = true //tipo boolean
	var booleano2 bool       //falso
	// Fim Booleanos

	var erro error //tipo error
	erro2 := errors.New("Erro interno")

	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numeroReal1)
	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(char1)
	fmt.Println(char2)
	fmt.Println(texto)
	fmt.Println(texto2)
	fmt.Println(booleano)
	fmt.Println(booleano2)
	fmt.Println(erro)
	fmt.Println(erro2)
}
