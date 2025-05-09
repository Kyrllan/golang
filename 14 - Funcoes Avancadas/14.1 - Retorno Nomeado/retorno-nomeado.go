package main

func calculosNomeados(n1, n2 int) (soma, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, subtracao := calculosNomeados(10, 20)
	println(soma, subtracao)
}
