package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"Backend/models"

	"github.com/gin-gonic/gin"
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

// MockStockRepository es una implementación mock de StockRepository.
type MockStockRepository struct {
	mock.Mock
}

func (m *MockStockRepository) GetStocks(db *gorm.DB, ticker, company, brokerage string) ([]models.Stock, error) {
	args := m.Called(db, ticker, company, brokerage)
	return args.Get(0).([]models.Stock), args.Error(1)
}

func (m *MockStockRepository) GetAllStocks(db *gorm.DB) ([]models.Stock, error) {
	args := m.Called(db)
	return args.Get(0).([]models.Stock), args.Error(1)
}

func TestGetStocks(t *testing.T) {
	// Configurar el mock del repositorio
	mockRepo := new(MockStockRepository)
	mockDB := new(MockDB)

	// Configurar el handler
	handler := NewStockHandler(mockDB)
	handler.repo = mockRepo

	// Configurar el contexto de Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/stocks", handler.GetStocks)

	// Caso de prueba: éxito
	t.Run("Success", func(t *testing.T) {
		mockRepo.On("GetStocks", mockDB, "AAPL", "", "").Return([]models.Stock{
			{Ticker: "AAPL", Company: "Apple Inc."},
		}, nil)

		req, _ := http.NewRequest(http.MethodGet, "/stocks?ticker=AAPL", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.Code)
		}
	})

	// Caso de prueba: error
	t.Run("Error", func(t *testing.T) {
		mockRepo.On("GetStocks", mockDB, "INVALID", "", "").Return([]models.Stock{}, gorm.ErrRecordNotFound)

		req, _ := http.NewRequest(http.MethodGet, "/stocks?ticker=INVALID", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.Code)
		}
	})
}
