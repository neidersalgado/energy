package service

import (
	"fmt"
	"sort"
	"time"
)

type DailyDataRetriever struct {
	repo Irepository
}

func NewDailyDataRetriever(repository Irepository) *DailyDataRetriever {
	return &DailyDataRetriever{repo: repository}
}

func (r *DailyDataRetriever) GetConsumption(meterID int, start, end time.Time) (map[string]interface{}, error) {
	consumption, err := r.repo.GetConsumptionByMeterIDAndDateRange(meterID, start, end)

	if err != nil {
		return nil, fmt.Errorf("failed to get consumption data: %w", err)
	}

	dataByDate := make(map[string]map[string]float64)

	for _, record := range consumption {
		dateStr := record.Date.Format("Jan 2")

		if _, ok := dataByDate[dateStr]; !ok {
			dataByDate[dateStr] = map[string]float64{
				"active":              0.0,
				"reactive_inductive":  0.0,
				"reactive_capacitive": 0.0,
				"exported":            0.0,
			}
		}

		dataByDate[dateStr]["active"] += record.ActiveEnergy
		dataByDate[dateStr]["reactive_inductive"] += record.ReactiveEnergy
		dataByDate[dateStr]["reactive_capacitive"] += record.CapacitiveReactive
		dataByDate[dateStr]["exported"] += record.Solar
	}

	var dataGraph = map[string][]float64{
		"active":              nil,
		"reactive_inductive":  nil,
		"reactive_capacitive": nil,
		"exported":            nil,
	}

	responseData := map[string]interface{}{
		"period":     getUniqueDates(dataByDate),
		"meter_id":   meterID,
		"address":    "mock address",
		"data_graph": dataGraph,
	}

	for _, v := range dataByDate {
		dataGraph["active"] = append(dataGraph["active"], v["active"])
		dataGraph["reactive_inductive"] = append(dataGraph["reactive_inductive"], v["reactive_inductive"])
		dataGraph["reactive_capacitive"] = append(dataGraph["reactive_capacitive"], v["reactive_capacitive"])
		dataGraph["exported"] = append(dataGraph["exported"], v["exported"])
	}

	return responseData, nil
}

func getUniqueDates(dataByDate map[string]map[string]float64) []string {
	uniqueDates := make([]string, 0, len(dataByDate))
	for date := range dataByDate {
		uniqueDates = append(uniqueDates, date)
	}
	sort.Strings(uniqueDates)
	return uniqueDates
}
