package main

import "fmt"

func main() {

	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

	fmt.Println(array1)
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)
	slice2 := slice[1:3]

	fmt.Println(slice2)

	slice2[1] = 100
	fmt.Println(slice2)
	fmt.Println(slice)

	//Arrays Internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	//Arrays Internos
	slice4 := make([]float32, 5, 6) // Parametros: Tipo, Tamanho, Capacidade
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 10)
	fmt.Println(slice4)
	slice4 = append(slice4, 20)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
