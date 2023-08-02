package main

import (
	"fmt"
	"time"
)

func main() {
	//Make cria um canal e só trafega dados desse tipo(recebe e envia esse tipo)
	canal := make(chan string)

	go escrever("Olá mundo!", canal) //a clásula go vai executar e não vai esperar ser concluída para continuar

	for {
		mensagem, aberto := <-canal //Enquanto o canal não receber um valor, ele vai ficar "travando" o sistema.
		fmt.Println(mensagem)
		if !aberto {
			break
		}
	}

	//Segunda forma de fazer
	canal = make(chan string)
	go escrever("Olá mundo!", canal)

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Millisecond * 200)
	}
	close(canal)
}
