package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type user struct {
	Id    uint32 `json:"id"`
	Name  string `json:"nome"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var user user

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struc"))
		return
	}

	db, err := banco.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	statement, err := db.Prepare("INSERT INTO users(name, email) VALUES (?,?)")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserted, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("Erro ao obter o id inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserted)))
}
