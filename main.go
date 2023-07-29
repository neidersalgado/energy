// main.go
package main

import (
	"fmt"
	"log"
	net "net/http"

	"github.com/energy/api/routes"
	"github.com/energy/internal/config"
	"github.com/energy/internal/consumption/delivery/http"
	repository "github.com/energy/internal/consumption/repository/mysql"
	"github.com/energy/pkg/db"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqlDB, err := db.NewDBConn()
	if err != nil {
		fmt.Println("Error al intentar abrir la conexión a la base de datos:", err)
		return
	}
	defer sqlDB.Close()

	gormDB, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetConfig().DBUser,
		config.GetConfig().DBPassword,
		config.GetConfig().DBHost,
		config.GetConfig().DBPort,
		config.GetConfig().DBName,
	))
	if err != nil {
		fmt.Println("Error al convertir *sql.DB a *gorm.DB:", err)
		return
	}

	repo := repository.New(gormDB)
	handler := http.NewHandler(repo)

	router := mux.NewRouter()
	routes.RegisterRoutes(router, handler)

	port := "8080" // Cambia el puerto si lo deseas
	fmt.Printf("Servidor iniciado en el puerto %s\n", port)
	log.Fatal(net.ListenAndServe(":"+port, router))
}
