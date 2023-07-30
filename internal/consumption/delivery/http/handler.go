package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	repository "github.com/energy/internal/consumption/repository/mysql"
	"github.com/energy/internal/consumption/service"
	strategy "github.com/energy/internal/consumption/service/strategy"
)

type Handler struct {
	service service.ConsumptionService
	repo    *repository.ConsumptionRepository
}

func NewHandler(repository *repository.ConsumptionRepository) *Handler {
	return &Handler{
		repo: repository,
	}
}

func (h *Handler) GetConsumptionData(w http.ResponseWriter, r *http.Request) {
	meterID := r.URL.Query().Get("meters_ids")
	id, _ := strconv.Atoi(meterID)
	fmt.Printf(" param %v", id)

	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")
	kindPeriod := r.URL.Query().Get("kind_period")

	startDate, _ := time.Parse("2006-01-02", startDateStr)
	endDate, _ := time.Parse("2006-01-02", endDateStr)
	var retriever service.ConsumptionRetriever
	switch kindPeriod {
	case "monthly":
		retriever = strategy.NewMonthlyDataRetriever(h.repo)
	case "weekly":
		retriever = strategy.NewWeeklyDataRetriever(h.repo)
	case "daily":
		retriever = strategy.NewDailyDataRetriever(h.repo)
	default:
		http.Error(w, "Invalid kind_period", http.StatusBadRequest)
		return
	}

	h.service = *service.NewConsumptionService(retriever)

	data, err := h.service.GetConsumptionData(id, startDate, endDate)
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
