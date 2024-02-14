package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// Ponteiros é uma referência de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro, *ponteiro) //desreferenciação

}

// var ponteiro *tipo -> define uma variável do tipo ponteiro
// ponteiro = &variavel -> ponteiro recebe o endereço de memória da variável
// fmt.Println(ponteiro, *ponteiro) -> escreve o ponteiro(Endereço de memória) e escreve o valor que está no ponteiro(valor do endereço de memória)
