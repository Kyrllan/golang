package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" // tipada
	variavel2 := "Variável 2"           // tipo implicito
	// varias variaveis declaradas em uma linha
	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)
	// varias variaveis declaradas em uma linha implicito
	variavel5, variavel6 := "Variável 5", "Variável 6"
	// redeclarando variaveis
	variavel5, variavel6 = variavel6, variavel5

	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3, variavel4)
	fmt.Println(variavel5, variavel6)
}
