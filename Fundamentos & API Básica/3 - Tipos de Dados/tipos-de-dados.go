package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64
	var num1 int = 1000000000
	fmt.Println(num1)

	//uint unsigned
	var num2 uint32 = 100
	fmt.Println(num2)

	//alias
	//int32 == rune
	var num3 rune = 12345
	fmt.Println(num3)

	var num4 byte = 123
	fmt.Println(num4)

	//float32, float64
	var num5 float32 = 123.123
	fmt.Println(num5)

	var num6 float64 = 123.123123123
	fmt.Println(num6)

	numero7 := 12313.123123
	fmt.Println(numero7)

	//string

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//char == int

	char := 'B'
	fmt.Println(char)

	var texto int
	fmt.Println(texto)

	//boolean
	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool
	fmt.Println(booleano2)

	//erro
	var erro error
	fmt.Println(erro)

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
