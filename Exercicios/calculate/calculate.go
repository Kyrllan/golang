package main

import "fmt"

func calculate(a, b int, op string) int {
	result := 0
	if op != "+" && op != "-" && op != "*" && op != "/" {
		return 0
	}
	if op == "+" {
		result = a + b
	}
	if op == "-" {
		result = a - b
	}
	if op == "*" {
		result = a * b
	}
	if op == "/" {
		result = a / b
	}
	return result

}

func main() {
	var a int
	var b int
	var op string
	fmt.Printf("Digite o primeiro número: ")
	fmt.Scanln(&a)
	fmt.Printf("Digite o segundo número: ")
	fmt.Scanln(&b)
	fmt.Printf("Digite a operação (+, -, *, /): ")
	fmt.Scanln(&op)
	result := calculate(a, b, op)
	fmt.Printf("Resultado: %d\n", result)
}
