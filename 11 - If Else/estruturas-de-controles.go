package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	//Criando variável no escopo do IF
	//if init
	//se variávelNova, recebe numero; outroNumero maior que 0, faço isso
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else if numero < -10 {
		fmt.Println("Número é menor que menos dez")
	} else {
		fmt.Println("Número é menor que zero")
	}

	//fmt.Println(outroNumero) variável de escopo local do IF
}
