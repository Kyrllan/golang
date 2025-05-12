package main

func clousure() func() {
	texto := "Dentro da função Closure"

	funcao := func() {
		println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	println(texto)
	funcaoNova := clousure()
	funcaoNova()
}
