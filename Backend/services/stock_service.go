package services

import (
	"sort"

	"Backend/models"
)

func CalculateStockScore(stock models.Stock) float64 {
	score := 0.0

	// Target price increase impact
	priceIncrease := stock.TargetTo - stock.TargetFrom
	if priceIncrease > 0 {
		score += 3
		if priceIncrease > 10 {
			score += 2
		}
	}

	// Rating changes impact
	if stock.RatingFrom == "Sell" && stock.RatingTo == "Buy" {
		score += 3
	} else if stock.RatingFrom == "Neutral" && stock.RatingTo == "Buy" {
		score += 2
	} else if stock.RatingFrom == "Sell" && stock.RatingTo == "Neutral" {
		score += 1
	}

	// Weight based on brokerage reputation
	topBrokers := map[string]float64{
		"The Goldman Sachs Group": 2,
		"JPMorgan Chase":          1.5,
		"Bank of America":         1,
	}

	if value, exists := topBrokers[stock.Brokerage]; exists {
		score += value
	}

	return score
}

func CalculateStockRecommendations(stocks []models.Stock) []models.StockRecommendation {
	recommendations := make([]models.StockRecommendation, 0)

	for _, stock := range stocks {
		score := CalculateStockScore(stock)
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
			Stock: stock, Score: score, Recommendation: recommendation,
		})
	}

	// Sort by highest score
	sort.Slice(recommendations, func(i, j int) bool {
		return recommendations[i].Score > recommendations[j].Score
	})

	return recommendations
}
