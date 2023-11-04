package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //driver connection
)

// Open connection with database
func Connect() (*sql.DB, error) {
	connString := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// func Close(conn *sql.DB) {
// 	conn.Close()
// }
