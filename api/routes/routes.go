// routes/routes.go
package routes

import (
	"github.com/energy/internal/consumption/delivery/http"
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, handler *http.Handler) {
	router.HandleFunc("/consumption", handler.GetConsumptionData).Methods("GET")
}
