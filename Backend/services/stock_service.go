package services

import (
	"sort"
	"strings"

	"Backend/models"
)

type Rating string

const (
	Sell           Rating = "Sell"
	MarketPerforme Rating = "Market Perform"
	Buy            Rating = "Buy"
	Neutral        Rating = "Neutral"
	Underweight    Rating = "Underweight"
	Overweight     Rating = "Overweight"
)

const (
	PriceIncreaseScore                  = 3
	LargePriceIncreaseBonus             = 2
	SellToBuyScore                      = 3
	NeutralToBuyScore                   = 2
	SellToNeutralScore                  = 1
	MarketPerformeToBuyScore            = 3
	MarketPerformeToNeutralScore        = 2
	MarketPerformeToSellScore           = 1
	MarketPerformeToMarketPerformeScore = 3
	MarketPerformeToUnderweightScore    = 2
	MarketPerformeToOverweightScore     = 1
	UnderweightToBuyScore               = 3
	UnderweightToNeutralScore           = 2
	UnderweightToSellScore              = 1
	UnderweightToMarketPerformeScore    = 3
	UnderweightToUnderweightScore       = 2
	UnderweightToOverweightScore        = 1
	OverweightToBuyScore                = 3
	OverweightToNeutralScore            = 2
	OverweightToSellScore               = 1
	OverweightToMarketPerformeScore     = 3
	OverweightToUnderweightScore        = 2
	OverweightToOverweightScore         = 1
)

type BrokerScorer interface {
	GetScore(brokerage string) float64
}

// DefaultBrokerScorer implementa BrokerScorer usando un mapa de brokers.
type DefaultBrokerScorer struct {
	TopBrokers map[string]float64
}

// GetScore devuelve el score de un broker.
func (s *DefaultBrokerScorer) GetScore(brokerage string) float64 {
	if value, exists := s.TopBrokers[brokerage]; exists {
		return value
	}
	return 0
}

// NewDefaultBrokerScorer crea una nueva instancia de DefaultBrokerScorer
func NewDefaultBrokerScorer(topBrokers map[string]float64) *DefaultBrokerScorer {
	return &DefaultBrokerScorer{
		TopBrokers: topBrokers,
	}
}

func calculatePriceImpact(stock models.Stock) float64 {
	priceIncrease := stock.TargetTo - stock.TargetFrom
	if priceIncrease > 0 {
		if priceIncrease > 10 {
			return PriceIncreaseScore + LargePriceIncreaseBonus
		}
		return PriceIncreaseScore
	}
	return 0
}

func calculateRatingImpact(stock models.Stock) float64 {
	if stock.RatingFrom == string(Sell) && stock.RatingTo == string(Buy) {
		return SellToBuyScore
	} else if stock.RatingFrom == string(Neutral) && stock.RatingTo == string(Buy) {
		return NeutralToBuyScore
	} else if stock.RatingFrom == string(Sell) && stock.RatingTo == string(Neutral) {
		return SellToNeutralScore
	} else if stock.RatingFrom == string(MarketPerforme) && stock.RatingTo == string(Buy) {
		return MarketPerformeToBuyScore
	} else if stock.RatingFrom == string(MarketPerforme) && stock.RatingTo == string(Neutral) {
		return MarketPerformeToNeutralScore
	} else if stock.RatingFrom == string(MarketPerforme) && stock.RatingTo == string(Sell) {
		return MarketPerformeToSellScore
	} else if stock.RatingFrom == string(MarketPerforme) && stock.RatingTo == string(MarketPerforme) {
		return MarketPerformeToMarketPerformeScore
	} else if stock.RatingFrom == string(MarketPerforme) && stock.RatingTo == string(Underweight) {
		return MarketPerformeToUnderweightScore
	} else if stock.RatingFrom == string(MarketPerforme) && stock.RatingTo == string(Overweight) {
		return MarketPerformeToOverweightScore
	} else if stock.RatingFrom == string(Underweight) && stock.RatingTo == string(Buy) {
		return UnderweightToBuyScore
	} else if stock.RatingFrom == string(Underweight) && stock.RatingTo == string(Neutral) {
		return UnderweightToNeutralScore
	} else if stock.RatingFrom == string(Underweight) && stock.RatingTo == string(Sell) {
		return UnderweightToSellScore
	} else if stock.RatingFrom == string(Underweight) && stock.RatingTo == string(MarketPerforme) {
		return UnderweightToMarketPerformeScore
	} else if stock.RatingFrom == string(Underweight) && stock.RatingTo == string(Underweight) {
		return UnderweightToUnderweightScore
	} else if stock.RatingFrom == string(Underweight) && stock.RatingTo == string(Overweight) {
		return UnderweightToOverweightScore
	} else if stock.RatingFrom == string(Overweight) && stock.RatingTo == string(Buy) {
		return OverweightToBuyScore
	} else if stock.RatingFrom == string(Overweight) && stock.RatingTo == string(Neutral) {
		return OverweightToNeutralScore
	} else if stock.RatingFrom == string(Overweight) && stock.RatingTo == string(Sell) {
		return OverweightToSellScore
	} else if stock.RatingFrom == string(Overweight) && stock.RatingTo == string(MarketPerforme) {
		return OverweightToMarketPerformeScore
	} else if stock.RatingFrom == string(Overweight) && stock.RatingTo == string(Underweight) {
		return OverweightToUnderweightScore
	} else if stock.RatingFrom == string(Overweight) && stock.RatingTo == string(Overweight) {
		return OverweightToOverweightScore
	}
	return 0
}

func CalculateStockScore(stock models.Stock, scorer BrokerScorer) float64 {
	return calculatePriceImpact(stock) + calculateRatingImpact(stock) + scorer.GetScore(stock.Brokerage)
}

func CalculateStockRecommendations(stocks []models.Stock, scorer BrokerScorer) []models.StockRecommendation {
	recommendations := make([]models.StockRecommendation, 0)

	for _, stock := range stocks {
		score := CalculateStockScore(stock, scorer)
		recommendation := ""

		if score >= 7 {
			recommendation = "Strong Buy"
		} else if score >= 5 {
			recommendation = "Buy"
		} else if score >= 3 {
			recommendation = "Hold"
		} else {
			recommendation = "Sell"
		}

		recommendations = append(recommendations, models.StockRecommendation{
			Stock:          stock,
			Score:          score,
			Recommendation: recommendation,
		})
	}

	// Ordenar por score más alto
	sort.SliceStable(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})

	return recommendations
}

// SanitizeInput limpia y sanitiza el input del usuario
func SanitizeInput(input string) string {
	// Eliminar espacios en blanco al inicio y final
	input = strings.TrimSpace(input)
	// Convertir a minúsculas para normalizar
	input = strings.ToLower(input)
	// Eliminar caracteres especiales peligrosos
	input = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == ' ' || r == '-' || r == '.' {
			return r
		}
		return -1
	}, input)
	return input
}
