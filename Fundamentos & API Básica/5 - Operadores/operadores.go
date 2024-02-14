package main

import "fmt"

func main() {
	//Aritméticos
	// + - / * %
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 1 * 2
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var num1 int16 = 10
	var num2 int16 = 25
	soma2 := num1 + num2
	fmt.Println(soma2)

	//Atribuição
	var variavel1 string = "STRING"
	variavel2 := "STRING2"
	fmt.Println(variavel1, "---", variavel2)

	//Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//Lógicos
	fmt.Println(true && true)  //And
	fmt.Println(true || false) //Or
	fmt.Println(!true)         //Not

	//Unários
	numero := 10
	numero++     //numero = numero + 1
	numero--     //numero = numero - 1
	numero += 10 //numero = numero + 15
	numero -= 10 //numero = numero - 10
	numero *= 3  //numero = numero * 3
	numero /= 2  //numero = numero / 10

	//Ternário

}
