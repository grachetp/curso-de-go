package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnString = ""
	Port       = 0
)

func LoadConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	ConnString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("BD_USER"),
		os.Getenv("BD_PASSWORD"),
		os.Getenv("BD_NAME"),
	)
}

// DB_USUARIO=golang
// DB_SENHA=golang
// DB_NOME=devbook
// API_PORT=5000
