package services

import (
	"testing"

	"Backend/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBrokerScorer es una implementación mock de BrokerScorer.
type MockBrokerScorer struct {
	mock.Mock
}

func (m *MockBrokerScorer) GetScore(brokerage string) float64 {
	args := m.Called(brokerage)
	return args.Get(0).(float64)
}

func TestCalculateStockScore(t *testing.T) {
	// Configurar el mock del BrokerScorer
	mockScorer := new(MockBrokerScorer)
	mockScorer.On("GetScore", "The Goldman Sachs Group").Return(2.0)

	// Caso de prueba: éxito
	t.Run("Success", func(t *testing.T) {
		stock := models.Stock{
			Ticker:     "AAPL",
			TargetFrom: 100,
			TargetTo:   150,
			RatingFrom: "Sell",
			RatingTo:   "Buy",
			Brokerage:  "The Goldman Sachs Group",
		}

		score := CalculateStockScore(stock, mockScorer)
		assert.Equal(t, 8.0, score) // 3 (price) + 3 (rating) + 2 (brokerage)
	})
}

func TestCalculateStockRecommendations(t *testing.T) {
	// Configurar el mock del BrokerScorer
	mockScorer := new(MockBrokerScorer)
	mockScorer.On("GetScore", "The Goldman Sachs Group").Return(2.0)

	// Caso de prueba: éxito
	t.Run("Success", func(t *testing.T) {
		stocks := []models.Stock{
			{
				Ticker:     "AAPL",
				TargetFrom: 100,
				TargetTo:   150,
				RatingFrom: "Sell",
				RatingTo:   "Buy",
				Brokerage:  "The Goldman Sachs Group",
			},
		}

		recommendations := CalculateStockRecommendations(stocks, mockScorer)
		assert.Equal(t, 1, len(recommendations))
		assert.Equal(t, "Strong Buy", recommendations[0].Recommendation)
	})
}
