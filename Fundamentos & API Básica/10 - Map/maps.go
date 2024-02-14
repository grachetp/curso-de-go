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
