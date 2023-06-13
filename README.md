# GO

Curso de GO - Udemy
[Aprenda Golang do Zero! Desenvolva uma APLICAÇÃO COMPLETA! | Udemy](https://www.udemy.com/course/aprenda-golang-do-zero-desenvolva-uma-aplicacao-completa/learn/lecture/22206096?start=0#overview)

## Fundamentos

### Hello World

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
}
```

### Packages

```go
go mod init modulo //Cria um módulo - Centralizador de dependências do GO
go build //Builda o binário da aplicação
go run ./main.go //Roda o main da aplicação.
go install //Builda o binário da aplicação, mas move para a raiz do projeto
go get {pacote} //Instala pacote na aplicação
go mod tidy //Desinstala pacotes que não estão sendo usados na aplicação

```

### Variáveis, Constantes e Tipos de Dados

```go
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
```

### Funções

```go
//Função Simples
func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Função com dois retornos
func calc(n1, n2 int8) (int8, string) {
	return n1 + n2, "batatao"
}

//Chamando funções
func main() {
	//Chamando Funções
	var result = sum(10, 20)
	resul1 := sum(10, 30)
	fmt.Println(result)
	fmt.Println(resul1)	

	//Funções Anonimas por tipo
	var f = func(txt string) {
			fmt.Println(txt)
		}
	
		f("batata")

	//Funções com retorno de dois tipos
	r1, r2 := calc(1, 2)
	fmt.Println(r1, r2)

	//Não declarar as duas variáveis de retorno
	result5, _ := calc(10, 30)
	fmt.Println(result5)
}
```
### Operadores

```go
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
numero++ //numero = numero + 1
numero-- //numero = numero - 1
numero += 10 //numero = numero + 15
numero -= 10 //numero = numero - 10
numero *= 3 //numero = numero * 3
numero /= 2 //numero = numero / 10
```

### Structs

```go
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
```