package repository

import (
	"time"

	"github.com/energy/internal/consumption/model"
	"github.com/jinzhu/gorm"
)

type ConsumptionRepository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *ConsumptionRepository {
	return &ConsumptionRepository{
		DB: db,
	}
}

func (r *ConsumptionRepository) GetConsumptionByMeterIDAndDateRange(meterID int, start, end time.Time) ([]model.Consumption, error) {
	var data []model.Consumption

	r.DB.LogMode(true)

	err := r.DB.Where("meter_id = ? AND date BETWEEN ? AND ?", meterID, start, end).
		Unscoped().
		Order("date ASC").
		Find(&data).Error

	return data, err
}
