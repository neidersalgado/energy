package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Consumption struct {
	gorm.Model
	ID                 string    `gorm:"column:ID;type:varchar(36);primary_key"`
	MeterID            int       `gorm:"column:METER_ID;type:int"`
	ActiveEnergy       float64   `gorm:"column:ACTIVE_ENERGY;type:double"`
	ReactiveEnergy     float64   `gorm:"column:REACTIVE_ENERGY;type:double"`
	CapacitiveReactive float64   `gorm:"column:CAPACITIVE_REACTIVE;type:double"`
	Solar              float64   `gorm:"column:SOLAR;type:double"`
	Date               time.Time `gorm:"column:DATE;type:datetime"`
}

func (Consumption) TableName() string {
	return "ENERGY"
}