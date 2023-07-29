package consumption

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Consumption struct {
	gorm.Model
	ID                 string    `gorm:"type:varchar(36);primary_key"`
	MeterID            int       `gorm:"type:int"`
	ActiveEnergy       float64   `gorm:"type:double"`
	ReactiveEnergy     float64   `gorm:"type:double"`
	CapacitiveReactive float64   `gorm:"type:double"`
	Solar              float64   `gorm:"type:double"`
	Date               time.Time `gorm:"type:datetime"`
}

func (Consumption) TableName() string {
	return "energy"
}
