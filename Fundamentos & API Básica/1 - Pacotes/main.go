package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo algo qualquer na main")
	auxiliar.Escrever()
	auxiliar.Escrever2()

	erro := checkmail.ValidateFormat("pedro.grachetgmail.com")
	fmt.Println(erro)
}
