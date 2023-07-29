package service

import (
	"time"

	"github.com/energy/internal/consumption/model"
)

type ConsumptionRepository interface {
	GetConsumptionByMeterIDAndDateRange(meterID int, start, end time.Time) ([]model.Consumption, error)
}

type MonthlyConsumptionService struct {
	Repo ConsumptionRepository
}

func NewMonthlyConsumptionService(repo ConsumptionRepository) MonthlyConsumptionService {
	return MonthlyConsumptionService{
		Repo: repo,
	}
}

func (s *MonthlyConsumptionService) GetMonthlyConsumption(meterID int, startDate, endDate time.Time) ([]MonthlyConsumptionResult, error) {

}
