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
