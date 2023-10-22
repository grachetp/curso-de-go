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

### Herança(Não é herança)

```go
package main

import "fmt"

type pessoa struct {
 nome      string
 sobrenome string
 idade     uint8
 altura    uint8
}

type estudante struct {
 pessoa
 curso     string
 faculdade string
}

func main() {
 fmt.Println("Herança")

 p1 := pessoa{"João", "Pedro", 20, 178}
 fmt.Println(p1)

 e1 := estudante{p1, "Engenharia", "USP"}
 fmt.Println(e1)

 fmt.Println(e1.nome)
}
```

### Ponteiros

```go
func main() {
 fmt.Println("Ponteiros")

 var variavel1 int = 10
 var variavel2 int = variavel1

 fmt.Println(variavel1, variavel2)

 variavel1++

 fmt.Println(variavel1, variavel2)

 // Ponteiros é uma referência de memória
 var variavel3 int
 var ponteiro *int

 variavel3 = 100
 ponteiro = &variavel3

 fmt.Println(variavel3, *ponteiro) //desreferenciação

}
```

### Arrays e Slices

```go
package main

import (
 "fmt"
 "reflect"
)

func main() {
 fmt.Println("Arrays e Slices")
 fmt.Println("Arrays")
 var array1 [5]string //Tamanho fixo
 array1[0] = "Posição 1"
 array1[1] = "Posição 2"
 fmt.Println(array1)

 array2 := [5]string{"1", "2", "3", "4", "5"}
 fmt.Println(array2)

 array3 := [...]int{1, 2, 3, 4, 5, 6, 7} //Tamanho baseado na quantidade de valores passados
 array3[3] = 10
 fmt.Println(array3)

 fmt.Println("Slices")
 slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12, 13, 13}
 fmt.Println(slice)

 fmt.Println(reflect.TypeOf(slice)) //Verificar tipo de uma variável
 fmt.Println(reflect.TypeOf(array3))

 slice = append(slice, 12)
 fmt.Println(slice)

 slice2 := array2[1:3] //inclusivo:exclusivo -> pega 1, 2, ignora o 3.  //Ponteiro
 fmt.Println(slice2)

 array2[1] = "Posição alterada"
 fmt.Println(slice2)
}
```

### Arrays Internos

```go
package main

import (
 "fmt"
 "reflect"
)

func main() {
 fmt.Println("Arrays e Slices")
 fmt.Println("Arrays")
 var array1 [5]string //Tamanho fixo
 array1[0] = "Posição 1"
 array1[1] = "Posição 2"
 fmt.Println(array1)

 array2 := [5]string{"1", "2", "3", "4", "5"}
 fmt.Println(array2)

 array3 := [...]int{1, 2, 3, 4, 5, 6, 7} //Tamanho baseado na quantidade de valores passados
 array3[3] = 10
 fmt.Println(array3)

 fmt.Println("Slices")
 slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 12, 13, 13}
 fmt.Println(slice)

 fmt.Println(reflect.TypeOf(slice)) //Verificar tipo de uma variável
 fmt.Println(reflect.TypeOf(array3))

 slice = append(slice, 12)
 fmt.Println(slice)

 slice2 := array2[1:3] //inclusivo:exclusivo -> pega 1, 2, ignora o 3.  //Ponteiro
 fmt.Println(slice2)

 array2[1] = "Posição alterada"
 fmt.Println(slice2)

 //Arrays Internos
 slice3 := make([]float32, 10, 11)
 fmt.Println(slice3)
 fmt.Println(len(slice3)) //tamanho
 fmt.Println(cap(slice3)) //capacidade

 slice3 = append(slice3, 14)
 fmt.Println(slice3)
 fmt.Println(len(slice3)) //tamanho
 fmt.Println(cap(slice3)) //capacidade

 slice3 = append(slice3, 14)
 fmt.Println(len(slice3)) //tamanho
 fmt.Println(cap(slice3)) //capacidade
 fmt.Println(slice3)

 slice4 := make([]float32, 5)
 fmt.Println(slice4)
 slice4 = append(slice4, 10)
 fmt.Println(len(slice4))
 fmt.Println(cap(slice4))

 //Array lista com tamanho fixo
 //Slice lista com tamanho dinâmico
}
```

### Maps

```go
package main

import "fmt"

func main() {
 fmt.Println("Maps")

 //map[tipodachave]tipodosvalores
 usuario := map[string]string{
  "nome":      "Pedro",
  "sobrenome": "Grachet",
 }

 fmt.Println(usuario)
 fmt.Println(usuario["nome"])
 fmt.Println(usuario["sobrenome"])

 usuario2 := map[string]map[string]string{
  "nome": {
   "primeiro": "João",
   "ultimo":   "Oliveira",
  },
  "curso": {
   "nomeDoCurso": "Engenharia de Software",
  },
 }

 fmt.Println(usuario2)
 fmt.Println(usuario2["nome"]["ultimo"])
 fmt.Println(usuario2["curso"]["nomeDoCurso"])

 delete(usuario2, "curso")
 fmt.Println(usuario2)

 usuario2["signo"] = map[string]string{
  "nome": "gemeos",
 }
 fmt.Println(usuario2)
}
```

### Estruturas de Controle

### If Else

```go
package main

import "fmt"

func main() {
 fmt.Println("Estruturas de controle")

 numero := 10

 if numero > 15 {
  fmt.Println("Maior que 15")
 } else {
  fmt.Println("Menor que 15")
 }

 //Criando variável no escopo do IF
 //if init
 //se variávelNova, recebe numero; outroNumero maior que 0, faço isso
 if outroNumero := numero; outroNumero > 0 {
  fmt.Println("Número é maior que zero")
 } else if numero < -10 {
  fmt.Println("Número é menor que menos dez")
 } else {
  fmt.Println("Número é menor que zero")
 }

 //fmt.Println(outroNumero) variável de escopo local do IF
}
```

### Switch

```go
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
```

### Loops

```go
package main

import (
 "fmt"
 "time"
)

func main() {
 fmt.Println("Loops")

 //Usar for em forma de while
 // i := 0

 // for i < 10 {
 //  i++
 //  fmt.Println("Incrementando i", i)
 //  time.Sleep(time.Millisecond * 50)
 // }

 //Usar for em forma de init
 for j := 0; j <= 10; j += 2 {
  fmt.Println("Incrementando j", j)
  time.Sleep(time.Millisecond * 100)
 }

 //For com cláusula range
 nomes := [3]string{"João", "Davi", "Lucas"}

 for indice, valor := range nomes {
  fmt.Println(indice, valor)
 }

 for _, valor := range nomes {
  fmt.Println(valor)
 }

 for indice, letra := range "PALAVRA" {
  fmt.Println(indice, string(letra))
 }

 usuario := map[string]string{
  "nome":      "Pedro",
  "sobremome": "Grachet",
 }

 for chave, valor := range usuario {
  fmt.Println(chave, valor)
 }

 //Não é possível utilizar o range em Structs

 //Loop Infinito
 // i := 0
 // for {
 //  fmt.Println(i)
 //  i++
 // }
}
```

### Funções Avançadas

### Retorno Nomeado

```go
package main

import "fmt"

func CalculoMatemáticos(n1, n2 int) (soma int, subtracao int) {
 soma = n1 + n2
 subtracao = n1 - n2
 return
}

func main() {
 soma, subtracao := CalculoMatemáticos(10, 20)
 fmt.Println(soma, subtracao)
}
```

### Funções  Variáticas

```go
package main

import "fmt"

func Soma(numeros ...int) int {
 //fmt.Println(numeros)
 total := 0
 for _, numero := range numeros {
  total += numero
 }
 return total
}

func Escrever(texto string, numeros ...int) {
 for _, numero := range numeros {
  fmt.Println(texto, numero)
 }
}

func main() {
 //resultado := Soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
 resultado := Soma()
 fmt.Println(resultado)

 Escrever("Contando até 6", 1, 2, 3, 4, 5, 6)
}
```

### Funções Anonimas

```go
package main

import "fmt"

func main() {

 func() {
  fmt.Println("Olá mundo!")
 }()

 func(texto string) {
  fmt.Println(texto)
 }("Batata")

 variavel := func(texto string) string {
  return fmt.Sprintf("Recebido -> %s", texto)
 }("Batata")

 fmt.Println(variavel)
}
```

### Funções Recursivas

```go
package main

import "fmt"

func Fibonnacci(posicao uint) uint {
 if posicao <= 1 {
  return posicao
 }

 return Fibonnacci(posicao-2) + Fibonnacci(posicao-1)
}

func main() {
 fmt.Println("Funções Recursivas")
 //Sequência de Fibonnacci
 // 1 1 2 3 5 8 13 21
 fmt.Println(Fibonnacci(uint(30)))
}
```

### Defer

```go
package main

import "fmt"

func funcao1() {
 fmt.Println("Executando a função 1")
}

func funcao2() {
 fmt.Println("Executando a função 2")
}

func AlunoEstaAprovado(n1, n2 float32) bool {
 defer fmt.Println("Média calculada, o resutaldo será retornado.")
 fmt.Println("Entrando na Função para verificar se o aluno está aprovado")
 media := (n1 + n2) / 2

 if media >= 6 {
  return true
 }
 return false
}

func main() {
 // defer funcao1()
 // //defer = adiar. Objetivo: adiar até o último momento possível.
 // Útiu para utilizar com banco de dados e etc
 // funcao2()

 fmt.Println(AlunoEstaAprovado(7, 8))
}
```

### Funções Panic e Recovery

```go
package main

import "fmt"

func recuperarExecucao() {
 //fmt.Println("Tentativa de recuperar a execução")
 if r := recover(); r != nil {
  fmt.Println("Execução recuperada com sucesso!")
 }
}

func alunoEstaAprovado(n1, n2 float64) bool {
 defer recuperarExecucao()

 media := (n1 + n2) / 2

 if media > 6 {
  return true
 } else if media < 6 {
  return false
 }

 panic("A média é exatamente 6!") //Para tudo e entra em pânico
}

func main() {
 fmt.Println(alunoEstaAprovado(8, 6))
 fmt.Println("Pós execução!")
}
```

### Funções Closure

```go
package main

import "fmt"

//Funções Closures são basicamente funções que referenciam variaveis fora do seu corpo

func closure() func() {
 texto := "Dentro da função closure"

 funcao := func() {
  fmt.Println(texto)
 }

 return funcao
}

func main() {
 texto := "Dentro da função main"
 fmt.Println(texto)

 funcaoNova := closure()
 funcaoNova()
}
```

### Funções com Ponteiros

```go
package main

import "fmt"

func inverterSinal(numero int) int {
 return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
 *numero *= -1
}

func main() {
 numero := 20
 numeroInvertido := inverterSinal(numero)
 fmt.Println(numeroInvertido)
 fmt.Println(numero)

 inverterSinalComPonteiro(&numero)
 fmt.Println(numero)
}
```

### Função Init

```go
package main

import "fmt"

func init() {
 fmt.Println("Executando a init")
}

func main() {
 fmt.Println("Executando")
}
```

### Interfaces

```go
package main

import (
 "fmt"
 "math"
)

type forma interface {
 area() float64
}

func escreverArea(f forma) {
 fmt.Printf("A área da forma é %0.2f\n", f.area())
}

type retangulo struct {
 altura  float64
 largura float64
}

func (r retangulo) area() float64 {
 return r.altura * r.largura
}

func (c circulo) area() float64 {
 return c.raio * c.raio * math.Pi
}

type circulo struct {
 raio float64
}

func main() {
 r := retangulo{10, 15}
 c := circulo{10}
 escreverArea(r)
 escreverArea(c)
}
```

### Interfaces Genéricas

```go
package main

import "fmt"

//Tudo atende essa interface
func generica(interf interface{}) {
 fmt.Println(interf)
}

func main() {
 generica("String")
 generica(1)
 generica(true)

 fmt.Println()

 mapa := map[interface{}]interface{}{}
}
```

### Goroutines

```go
package main

import (
 "fmt"
 "time"
)

func main() {
 //Concorrência é diferente de paralelismo
 go escrever("Olá, mundo!") // goroutine - Executa essa função, mas não espera ela terminar para continuar executando o programa
 escrever("Programando em Golang!")
}

func escrever(texto string) {
 for {
  fmt.Println(texto)
  time.Sleep(time.Second)
 }
}
```

### WaitGroup - Sincronizador de Goroutines

```go
package main

import (
 "fmt"
 "sync"
)

func main() {
 var waitGroup sync.WaitGroup

 waitGroup.Add(2) //Grupo de espera - Quantidade de goroutine que serão executadas ao mesmo tempo

 go func() {
  escrever("Programando em Golang!")
  waitGroup.Done() //Tira uma goroutine do contador do grupo de espera
 }()

 go func() {
  escrever("Olá, mundo!")
  waitGroup.Done() //Tira uma goroutine do contador do grupo de espera
 }()

 waitGroup.Wait() //Espera a contagem das goroutines chegar em 0

}

func escrever(texto string) {
 for i := 0; i < 250; i++ {
  fmt.Println(texto)
 }
}
```
