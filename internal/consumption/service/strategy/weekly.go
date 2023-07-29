package service

import (
	"fmt"
	"sort"
	"time"

	"github.com/energy/internal/consumption/model"
)

type WeeklyDataRetriever struct {
	repo Irepository
}

type DataGraph struct {
	Active             []float64 `json:"active"`
	Exported           []float64 `json:"exported"`
	ReactiveCapacitive []float64 `json:"reactive_capacitive"`
	ReactiveInductive  []float64 `json:"reactive_inductive"`
}

func NewWeeklyDataRetriever(repository Irepository) *WeeklyDataRetriever {
	return &WeeklyDataRetriever{repo: repository}
}

func (r *WeeklyDataRetriever) GetConsumption(meterID int, start, end time.Time) (map[string]interface{}, error) {
	consumption, err := r.repo.GetConsumptionByMeterIDAndDateRange(meterID, start, end)

	if err != nil {
		return nil, fmt.Errorf("failed to get consumption data: %w", err)
	}

	// Organizar los datos de consumo en semanas
	weeklyData := make(map[string][]model.Consumption)
	for _, data := range consumption {
		weekStart := data.Date.AddDate(0, 0, -int(data.Date.Weekday()))
		weekEnd := weekStart.AddDate(0, 0, 6)
		weekKey := fmt.Sprintf("%s - %s", weekStart.Format("Jan 2"), weekEnd.Format("Jan 2"))
		weeklyData[weekKey] = append(weeklyData[weekKey], data)
	}

	// Ordenar las semanas por fecha
	var weeks []string
	for week := range weeklyData {
		weeks = append(weeks, week)
	}
	sort.Strings(weeks)

	// Calcular la sumatoria de los consumos semanales
	responseData := make(map[string]interface{})
	responseData["address"] = "Dirección mock"
	responseData["meter_id"] = meterID

	dataGraph := DataGraph{
		Active:             make([]float64, 0),
		Exported:           make([]float64, 0),
		ReactiveCapacitive: make([]float64, 0),
		ReactiveInductive:  make([]float64, 0),
	}

	for _, week := range weeks {
		weeklyConsumption := weeklyData[week]
		activeSum, exportedSum, reactiveCapSum, reactiveIndSum := 0.0, 0.0, 0.0, 0.0
		for _, data := range weeklyConsumption {
			activeSum += data.ActiveEnergy
			exportedSum += data.Solar
			reactiveCapSum += data.CapacitiveReactive
			reactiveIndSum += data.ReactiveEnergy
		}
		dataGraph.Active = append(dataGraph.Active, activeSum)
		dataGraph.Exported = append(dataGraph.Exported, exportedSum)
		dataGraph.ReactiveCapacitive = append(dataGraph.ReactiveCapacitive, reactiveCapSum)
		dataGraph.ReactiveInductive = append(dataGraph.ReactiveInductive, reactiveIndSum)
	}

	responseData["data_graph"] = dataGraph
	responseData["period"] = weeks

	return responseData, nil
}
