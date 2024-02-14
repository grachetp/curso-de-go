package main

import "fmt"

func Soma(numeros ...int) int {
	//fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func Escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	//resultado := Soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	resultado := Soma()
	fmt.Println(resultado)

	Escrever("Contando até 6", 1, 2, 3, 4, 5, 6)
}
