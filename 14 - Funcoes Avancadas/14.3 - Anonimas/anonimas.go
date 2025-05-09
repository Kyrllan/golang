package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função anônima")
	}()

	func(texto string) {
		fmt.Println("Outra função anônima")
	}("Passando parâmetros")
}
