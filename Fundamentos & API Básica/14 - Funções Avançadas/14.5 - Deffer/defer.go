package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func AlunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, o resutaldo será retornado.")
	fmt.Println("Entrando na Função para verificar se o aluno está aprovado")
	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {
	// defer funcao1()
	// //defer = adiar. Objetivo: adiar até o último momento possível.
	// Útiu para utilizar com banco de dados e etc
	// funcao2()

	fmt.Println(AlunoEstaAprovado(7, 8))
}
