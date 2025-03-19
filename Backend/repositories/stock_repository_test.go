package repositories

import (
	"testing"

	"Backend/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockDB es una implementación mock de gorm.DB para pruebas.
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	m.Called(query, args)
	return &gorm.DB{}
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	m.Called(dest, conds)
	return &gorm.DB{}
}

func TestGetStocks(t *testing.T) {
	// Configurar el mock de la base de datos
	mockDB := new(MockDB)
	mockDB.On("Where", "ticker = ?", []interface{}{"AAPL"}).Return(mockDB)
	mockDB.On("Find", &[]models.Stock{}).Return(mockDB)

	// Caso de prueba: éxito
	t.Run("Success", func(t *testing.T) {
		stocks, err := GetStocks(mockDB, "AAPL", "", "")
		assert.NoError(t, err)
		assert.Equal(t, 0, len(stocks)) // El mock no devuelve datos
	})
}
