package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	fmt.Println("API running!")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), r))
}
