package service

import (
	"time"
)

type ConsumptionRetriever interface {
	GetConsumption(meterID int, start, end time.Time) (map[string]interface{}, error)
}

type ConsumptionService struct {
	Retriever ConsumptionRetriever
}

func NewConsumptionService(retriever ConsumptionRetriever) *ConsumptionService {
	return &ConsumptionService{Retriever: retriever}
}

func (s *ConsumptionService) GetConsumptionData(meterID int, start, end time.Time) (map[string]interface{}, error) {
	return s.Retriever.GetConsumption(meterID, start, end)
}
