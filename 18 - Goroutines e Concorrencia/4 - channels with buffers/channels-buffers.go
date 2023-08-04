package main

import "fmt"

//Por padrão usamos canais em funções separadas
//Uma função que envia o valor e outra que recebe

func main() {
	canal := make(chan string, 2) //Canal com capacidade de 2. Só vai bloquear a capacidade máxima do canal.

	canal <- "Olá mundo!"
	canal <- "Programando em Golang!"

	mensagem := <-canal
	fmt.Println(mensagem)
}
