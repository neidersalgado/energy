package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	repository "github.com/energy/internal/consumption/repository/mysql"
	"github.com/energy/internal/consumption/service"
	strategy "github.com/energy/internal/consumption/service/strategy"
	"github.com/gorilla/mux"
)

type Handler struct {
	service service.ConsumptionService
	repo    *repository.ConsumptionRepository
}

func NewHandler(service service.ConsumptionService, repository repository.ConsumptionRepository) *Handler {
	return &Handler{
		service: service,
		repo:    &repository,
	}
}

func (h *Handler) GetConsumptionData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	meterID, err := strconv.Atoi(params["meter_id"])
	if err != nil {
		http.Error(w, "Failed cas meterID", http.StatusBadRequest)
		return
	}
	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")
	kindPeriod := r.URL.Query().Get("kind_period")

	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)
	var retriever service.ConsumptionRetriever
	switch kindPeriod {
	case "monthly":
		retriever = strategy.NewMonthlyDataRetriever(h.repo) // Pasar h.repo en lugar de &h.repo
	case "weekly":
		retriever = strategy.NewWeeklyDataRetriever(h.repo) // Pasar h.repo en lugar de &h.repo
	case "daily":
		retriever = strategy.NewDailyDataRetriever(h.repo) // Pasar h.repo en lugar de &h.repo
	default:
		http.Error(w, "Invalid kind_period", http.StatusBadRequest)
		return
	}

	h.service = *service.NewConsumptionService(retriever)

	data, err := h.service.GetConsumptionData(meterID, startDate, endDate)
	if err != nil {
		http.Error(w, "Failed to get consumption data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed encode response", http.StatusInternalServerError)
	}
}