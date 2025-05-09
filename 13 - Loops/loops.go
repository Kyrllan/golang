package main

func main() {
	/* 	i := 0

	   	for i < 10 {
	   		time.Sleep(time.Second)
	   		fmt.Println(i)
	   		i++
	   	}

	   	for j := 0; j < 5; j++ {
	   		time.Sleep(time.Second)
	   		fmt.Println(j)
	   	} */

	nomes := []string{"João", "Davi", "Kyrllan"}

	for i, nome := range nomes {
		println(i, nome)
	}

	for i, letra := range "PALAVRA" {
		println(i, string(letra))
	}

	usuario := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		println(chave, valor)
	}

}
