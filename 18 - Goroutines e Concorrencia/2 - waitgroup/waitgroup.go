package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //Grupo de espera - Quantidade de goroutine que serão executadas ao mesmo tempo

	go func() {
		escrever("Programando em Golang!")
		waitGroup.Done() //Tira uma goroutine do contador do grupo de espera
	}()

	go func() {
		escrever("Olá, mundo!")
		waitGroup.Done() //Tira uma goroutine do contador do grupo de espera
	}()

	waitGroup.Wait() //Espera a contagem das goroutines chegar em 0

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
