package service_test

import (
	"testing"
	"time"

	service "github.com/energy/internal/consumption/service/strategy"
	"github.com/energy/internal/consumption/service/strategy/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDailyDataRetriever_GetConsumption(t *testing.T) {
	repoMock := mocks.NewMockRepository()
	start := time.Date(2023, time.July, 25, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, time.August, 15, 0, 0, 0, 0, time.UTC)

	repoMock.On("GetConsumptionByMeterIDAndDateRange", 1, start, end).Return(mocks.MockConsumptionData(), nil).Times(1)
	srv := service.NewDailyDataRetriever(repoMock)

	data, err := srv.GetConsumption(1, start, end)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	repoMock.AssertExpectations(t)
}
