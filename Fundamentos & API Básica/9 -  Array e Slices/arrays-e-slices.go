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
