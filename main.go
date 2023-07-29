package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/energy/pkg/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := db.NewDBConn()
	if err != nil {
		fmt.Println("Error al intentar abrir la conexión a la base de datos:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conexión a la base de datos exitosa!")

	router := mux.NewRouter()
	http.Handle()
}
