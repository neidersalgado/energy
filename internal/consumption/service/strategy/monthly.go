package service

import (
	"fmt"
	"sort"
	"time"

	"github.com/energy/internal/consumption/model"
)

type Irepository interface {
	GetConsumptionByMeterIDAndDateRange(meterID int, start, end time.Time) ([]model.Consumption, error)
}

type MonthlyDataRetriever struct {
	repo Irepository
}

func NewMonthlyDataRetriever(repository Irepository) *MonthlyDataRetriever {
	return &MonthlyDataRetriever{repo: repository}
}

func (r *MonthlyDataRetriever) GetConsumption(meterID int, start, end time.Time) (map[string]interface{}, error) {
	consumption, err := r.repo.GetConsumptionByMeterIDAndDateRange(meterID, start, end)

	if err != nil {
		return nil, fmt.Errorf("failed to get consumption data: %w", err)
	}

	dataByMonth := make(map[string]map[string]float64)

	for _, record := range consumption {
		yearMonthStr := record.Date.Format("Jan 2006")

		if _, ok := dataByMonth[yearMonthStr]; !ok {
			dataByMonth[yearMonthStr] = map[string]float64{
				"active":              0.0,
				"reactive_inductive":  0.0,
				"reactive_capacitive": 0.0,
				"exported":            0.0,
			}
		}

		dataByMonth[yearMonthStr]["active"] += record.ActiveEnergy
		dataByMonth[yearMonthStr]["reactive_inductive"] += record.ReactiveEnergy
		dataByMonth[yearMonthStr]["reactive_capacitive"] += record.CapacitiveReactive
		dataByMonth[yearMonthStr]["exported"] += record.Solar
	}

	var dataGraph = map[string][]float64{
		"active":              nil,
		"reactive_inductive":  nil,
		"reactive_capacitive": nil,
		"exported":            nil,
	}

	responseData := map[string]interface{}{
		"period":     getUniqueMonths(dataByMonth),
		"meter_id":   meterID,
		"address":    "mock address",
		"data_graph": dataGraph,
	}

	for _, v := range dataByMonth {
		dataGraph["active"] = append(dataGraph["active"], v["active"])
		dataGraph["reactive_inductive"] = append(dataGraph["reactive_inductive"], v["reactive_inductive"])
		dataGraph["reactive_capacitive"] = append(dataGraph["reactive_capacitive"], v["reactive_capacitive"])
		dataGraph["exported"] = append(dataGraph["exported"], v["exported"])
	}

	return responseData, nil
}

func getUniqueMonths(dataByMonth map[string]map[string]float64) []string {
	uniqueMonths := make([]string, 0, len(dataByMonth))
	for yearMonth := range dataByMonth {
		uniqueMonths = append(uniqueMonths, yearMonth)
	}
	sort.Strings(uniqueMonths)
	return uniqueMonths
}
