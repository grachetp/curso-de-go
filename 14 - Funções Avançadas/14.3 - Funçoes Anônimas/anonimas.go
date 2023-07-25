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
