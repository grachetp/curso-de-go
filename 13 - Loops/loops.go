package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	//Usar for em forma de while
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i", i)
	// 	time.Sleep(time.Millisecond * 50)
	// }

	//Usar for em forma de init
	for j := 0; j <= 10; j += 2 {
		fmt.Println("Incrementando j", j)
		time.Sleep(time.Millisecond * 100)
	}

	//For com cláusula range
	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	for _, valor := range nomes {
		fmt.Println(valor)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobremome": "Grachet",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//Não é possível utilizar o range em Structs

	//Loop Infinito
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }
}
