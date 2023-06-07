package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "teste1"
		variavel4 string = "teste2"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "var 5", "var 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
