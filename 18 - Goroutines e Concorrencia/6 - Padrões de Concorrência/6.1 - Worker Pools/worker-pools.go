package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultados := <-resultados
		fmt.Println(resultados)
	}

}

// Definindo o que o canal faz.
// tarefas só recebe dados
// resultados só envia dados
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonnacci(numero)
	}
}

func fibonnacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonnacci(posicao-2) + fibonnacci(posicao-1)
}
