package routes

import (
	"github.com/energy/internal/consumption/delivery/http"
	"github.com/energy/internal/consumption/repository"
	"github.com/energy/internal/consumption/service"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	repo := repository.NewMySQLRepository() // Implementa la interfaz repository.ConsumptionRepository
	service := service.NewConsumptionService(repo) // Implementa la interfaz service.ConsumptionService
	handler := http.NewHandler(service)

	router.HandleFunc("/consumption", handler.GetConsumptionData).Methods("GET")
}
