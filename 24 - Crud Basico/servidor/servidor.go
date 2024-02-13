package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	Id    uint32 `json:"id"`
	Name  string `json:"nome"`
	Email string `json:"email"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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

func FindUser(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro ID para inteiro"))
		return
	}

	db, err := banco.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	line, err := db.Query("SELECT * FROM users WHERE id = ?", ID)
	if err != nil {
		w.Write([]byte("Erro ao buscar usuário no banco de dados"))
		return
	}

	var user user
	if line.Next() {
		if err := line.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	if user.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para JSON!"))
		return
	}
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db, err := banco.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	lines, err := db.Query("SELECT * FROM users")
	if err != nil {
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}
	defer lines.Close()

	var users []user

	for lines.Next() {
		var user user

		if err := lines.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Erro ao escanear o usuário"))
			return
		}

		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro ID para inteiro"))
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var user user
	if err := json.Unmarshal(bodyRequest, &user); err != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, err := banco.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Erro ao atualizar o usuário"))
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, err := banco.Connect()
	if err != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from users where id = ?")
	if err != nil {
		w.Write([]byte("Erro ao criar statment!"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
