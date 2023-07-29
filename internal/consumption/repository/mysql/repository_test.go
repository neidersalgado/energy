package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	repository "github.com/energy/internal/consumption/repository/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestGetConsumptionByMeterIDAndDateRange(t *testing.T) {
	// Arrange
	sqlDB, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer sqlDB.Close()

	db, err := gorm.Open("mysql", sqlDB)
	assert.Nil(t, err)

	start := time.Now()
	end := start.Add(24 * time.Hour)

	rows := sqlmock.NewRows([]string{"ID", "METER_ID", "ACTIVE_ENERGY", "REACTIVE_ENERGY", "CAPACITIVE_REACTIVE", "SOLAR", "DATE"}).
		AddRow("1", 1, 1.0, 1.0, 1.0, 1.0, start).
		AddRow("2", 1, 2.0, 2.0, 2.0, 2.0, end)

	query := "^SELECT \\* FROM  \\`ENERGY\\` WHERE \\(meter_id = \\? AND date BETWEEN \\? AND \\?\\) ORDER BY date ASC"

	mock.ExpectQuery(query).
		WithArgs(1, start, end).
		WillReturnRows(rows)

	r := repository.New(db)

	// Act
	data, err := r.GetConsumptionByMeterIDAndDateRange(1, start, end)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 2, len(data))
	assert.Equal(t, "1", data[0].ID)
	assert.Equal(t, 1, data[0].MeterID)
	assert.Equal(t, 1.0, data[0].ActiveEnergy)
	assert.Equal(t, 1.0, data[0].ReactiveEnergy)
	assert.Equal(t, 1.0, data[0].CapacitiveReactive)
	assert.Equal(t, 1.0, data[0].Solar)
	assert.Equal(t, start, data[0].Date)

	assert.Equal(t, "2", data[1].ID)
	assert.Equal(t, 1, data[1].MeterID)
	assert.Equal(t, 2.0, data[1].ActiveEnergy)
	assert.Equal(t, 2.0, data[1].ReactiveEnergy)
	assert.Equal(t, 2.0, data[1].CapacitiveReactive)
	assert.Equal(t, 2.0, data[1].Solar)
	assert.Equal(t, end, data[1].Date)
}
