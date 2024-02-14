package main

import "fmt"

func diaDasemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Numero inválido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		fallthrough //Faz teu código cair dentro da próxima condição, mas sem fazer verificação
	case numero == 2:
		return "Segunda-Feira"
		//break não existe no Golang
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Numero inválido"
	}
}

func main() {
	fmt.Println("Switch")
	var diaDaSemana = diaDasemana(5)
	var diaDaSemana2 = diaDaSemana2(5)
	fmt.Println(diaDaSemana, diaDaSemana2)
}
