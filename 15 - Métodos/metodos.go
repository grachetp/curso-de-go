package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//Altera o valor por ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Pedro", 22}

	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
