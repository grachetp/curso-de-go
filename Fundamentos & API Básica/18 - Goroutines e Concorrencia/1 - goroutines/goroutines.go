package main

import (
	"fmt"
	"time"
)

func main() {
	//Concorrência é diferente de paralelismo
	go escrever("Olá, mundo!") // goroutine - Executa essa função, mas não espera ela terminar para continuar executando o programa
	escrever("Programando em Golang!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
