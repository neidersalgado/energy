package mocks

import (
	"time"

	"github.com/energy/internal/consumption/model"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func (m *MockRepository) GetConsumptionByMeterIDAndDateRange(meterID int, start, end time.Time) ([]model.Consumption, error) {
	args := m.Called(meterID, start, end)
	return args.Get(0).([]model.Consumption), args.Error(1)
}

func MockConsumptionData() []model.Consumption {
	consumptionData := []model.Consumption{
		{
			ID:                 "8bf1c442-1e6e-4720-8f38-cf6322e719db",
			MeterID:            1,
			ActiveEnergy:       310.51294,
			ReactiveEnergy:     61.3086,
			CapacitiveReactive: 61.3086,
			Solar:              4.064753395119119,
			Date:               time.Date(2022, 8, 25, 0, 8, 39, 0, time.UTC),
		},
		{
			ID:                 "c854bd02-2a39-439e-9524-397e667e31da",
			MeterID:            1,
			ActiveEnergy:       310.38232,
			ReactiveEnergy:     61.30225,
			CapacitiveReactive: 61.30225,
			Solar:              4.0631472743659485,
			Date:               time.Date(2022, 8, 24, 23, 57, 8, 0, time.UTC),
		},
		{
			ID:                 "bb067163-3439-4ef8-ae09-4b29f216c044",
			MeterID:            1,
			ActiveEnergy:       309.27429,
			ReactiveEnergy:     61.21935,
			CapacitiveReactive: 61.21935,
			Solar:              4.05190417735569,
			Date:               time.Date(2022, 8, 26, 22, 43, 19, 0, time.UTC),
		},
		{
			ID:                 "505ed30e-9386-4801-95b3-f67c69822969",
			MeterID:            1,
			ActiveEnergy:       308.64334,
			ReactiveEnergy:     61.1635,
			CapacitiveReactive: 61.1635,
			Solar:              4.04620141097223,
			Date:               time.Date(2022, 8, 24, 21, 57, 11, 0, time.UTC),
		},
		{
			ID:                 "eb72a75a-2705-4e23-b108-bfbfbb0ffa21",
			MeterID:            1,
			ActiveEnergy:       307.90561,
			ReactiveEnergy:     61.07565,
			CapacitiveReactive: 61.07565,
			Solar:              4.041380812156727,
			Date:               time.Date(2022, 8, 27, 20, 58, 41, 0, time.UTC),
		},
		{
			ID:                 "6012d160-8a66-4d14-88ed-4b28d79873bd",
			MeterID:            1,
			ActiveEnergy:       307.18256,
			ReactiveEnergy:     60.80365,
			CapacitiveReactive: 60.80365,
			Solar:              4.052041448169642,
			Date:               time.Date(2022, 8, 24, 19, 49, 28, 0, time.UTC),
		},
		{
			ID:                 "c9feff05-3983-4610-bdab-d893bd415cb5",
			MeterID:            1,
			ActiveEnergy:       306.29083,
			ReactiveEnergy:     60.73145,
			CapacitiveReactive: 60.73145,
			Solar:              4.043364352407196,
			Date:               time.Date(2022, 8, 25, 18, 53, 8, 0, time.UTC),
		},
		{
			ID:                 "276fda18-0893-4c77-8e6d-3f72378ea6af",
			MeterID:            1,
			ActiveEnergy:       304.70724,
			ReactiveEnergy:     60.6508,
			CapacitiveReactive: 60.6508,
			Solar:              4.023960772157993,
			Date:               time.Date(2022, 8, 29, 17, 54, 7, 0, time.UTC),
		},
		{
			ID:                 "872241ac-bf4c-4264-96ed-f03a24270f69",
			MeterID:            1,
			ActiveEnergy:       303.82883,
			ReactiveEnergy:     60.40475,
			CapacitiveReactive: 60.40475,
			Solar:              4.02988308038689,
			Date:               time.Date(2022, 8, 30, 16, 57, 5, 0, time.UTC),
		},
		{
			ID:                 "1be63fa4-1c32-475d-90a6-e85ccd2c834e",
			MeterID:            1,
			ActiveEnergy:       303.0769,
			ReactiveEnergy:     60.02965,
			CapacitiveReactive: 60.02965,
			Solar:              4.048786724560281,
			Date:               time.Date(2022, 8, 20, 15, 59, 44, 0, time.UTC),
		},
		{
			ID:                 "53b9e101-176c-4cb9-a3d7-020277f46d78",
			MeterID:            1,
			ActiveEnergy:       302.31982,
			ReactiveEnergy:     5.61665,
			CapacitiveReactive: 5.61665,
			Solar:              52.82564696037674,
			Date:               time.Date(2022, 8, 20, 14, 52, 19, 0, time.UTC),
		},
		{
			ID:                 "e3fdeedb-5124-49eb-99e4-1c045f3ca5c8",
			MeterID:            1,
			ActiveEnergy:       301.56497,
			ReactiveEnergy:     5.39865,
			CapacitiveReactive: 5.39865,
			Solar:              52.81777923762402,
			Date:               time.Date(2022, 8, 26, 13, 44, 55, 0, time.UTC),
		},
		{
			ID:                 "7740c7b9-6aa2-42f5-8643-9e3a6c949a14",
			MeterID:            1,
			ActiveEnergy:       301.1106,
			ReactiveEnergy:     5.35655,
			CapacitiveReactive: 5.35655,
			Solar:              52.81247213546051,
			Date:               time.Date(2022, 8, 30, 12, 54, 34, 0, time.UTC),
		},
	}

	return consumptionData
}
