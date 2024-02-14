package main

import "fmt"

func Fibonnacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return Fibonnacci(posicao-2) + Fibonnacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")
	//Sequência de Fibonnacci
	// 1 1 2 3 5 8 13 21
	fmt.Println(Fibonnacci(uint(30)))
}
