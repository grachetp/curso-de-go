package main

import (
	"fmt"
	"time"
)

func main() {
	//Make cria um canal e só trafega dados desse tipo(recebe e envia esse tipo)
	canal := make(chan string)

	go escrever("Olá mundo!", canal) //a clásula go vai executar e não vai esperar ser concluída para continuar

	fmt.Println("Depois da função escrever() começar a ser executada")

	mensagem := <-canal //Enquanto o canal não receber um valor, ele vai ficar "travando" o sistema.
	fmt.Println(mensagem)
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Millisecond * 200)
	}
}
