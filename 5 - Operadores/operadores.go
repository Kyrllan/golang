package main

import "fmt"

func main() {
	// Aritimeticos
	soma := 1 + 2
	subtracao := 10 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 2
	modulo := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, modulo)

	var numero1 int32 = 10
	var numero2 int32 = 20
	var soma2 int32 = numero1 + numero2 //Go não permite soma de tipos diferentes
	fmt.Println(soma2)

	// Atribuição
	var variavel1 string = "Variavel 1" //igual
	variavel2 := "Variavel 2"           //igual
	fmt.Println(variavel1, variavel2)

	// Relacionais - retorna true ou false
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println("----------------------")
	// Lógicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Unários
	numero3 := 10
	fmt.Println(numero3)
	numero3++ //numero = numero + 1
	fmt.Println(numero3)
	numero3 -= 4
	fmt.Println(numero3)
	numero3 *= 2
	fmt.Println(numero3)

	// Ternário
	//texto := condicao ? "Verdadeiro" : "Falso" - Não existe ternário em Go
	var texto string
	if numero3 > 10 {
		texto = "Verdadeiro"
	} else {
		texto = "Falso"
	}
	fmt.Println(texto)

}
