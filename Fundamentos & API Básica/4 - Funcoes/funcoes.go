package main

import "fmt"

func main() {
	var result = sum(10, 20)
	resul1 := sum(10, 30)
	fmt.Println(result)
	fmt.Println(resul1)

	r1, r2 := calc(1, 2)
	fmt.Println(r1, r2)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("batata")

	var f1 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result2 := f1("batata")
	fmt.Println(result2)

	result3, result4 := calc(1, 2)
	fmt.Println(result3, result4)

	//Não declarar as duas variáveis de retorno
	result5, _ := calc(10, 30)
	fmt.Println(result5)
}

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calc(n1, n2 int8) (int8, string) {
	return n1 + n2, "batatao"
}
