package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivos Struct")

	//Primeira Forma
	var user1 usuario

	fmt.Println(user1)
	fmt.Println(user1.nome, user1.idade)

	user1.idade = 22
	user1.nome = "Pedro Grachet"

	fmt.Println(user1)
	fmt.Println(user1.nome, user1.idade)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	//Segunda Forma
	user2 := usuario{"Pedro Grachet", 22, enderecoExemplo}
	fmt.Println(user2)
	fmt.Println(user2.nome, user2.idade)

	//Terceira forma
	user3 := usuario{idade: 22}
	fmt.Println(user3)
	fmt.Println(user3.nome, user3.idade)
}
