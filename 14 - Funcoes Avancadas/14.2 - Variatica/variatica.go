package main

func soma(numeros ...int) (total int) {
	for _, numero := range numeros {
		total += numero
	}
	return
}

func main() {
	total := soma(1, 2, 3, 7, 8, 9, 10, 25, 37, 48, 60)
	print(total)
}
