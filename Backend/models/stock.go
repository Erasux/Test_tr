package models

import (
	"fmt"
)

type Stock struct {
	ID         int64   `gorm:"primaryKey" json:"id"`
	Ticker     string  `json:"ticker"`
	Company    string  `json:"company"`
	TargetFrom float64 `json:"target_from"`
	TargetTo   float64 `json:"target_to"`
	Action     string  `json:"action"`
	Brokerage  string  `json:"brokerage"`
	RatingFrom string  `json:"rating_from"`
	RatingTo   string  `json:"rating_to"`
	Time       string  `json:"time"`
}

// Método para convertir a DTO
func (s *Stock) ToDTO() map[string]interface{} {
	return map[string]interface{}{
		"ticker":      s.Ticker,
		"company":     s.Company,
		"target_from": fmt.Sprintf("%.2f", s.TargetFrom),
		"target_to":   fmt.Sprintf("%.2f", s.TargetTo),
		"action":      s.Action,
		"brokerage":   s.Brokerage,
		"rating_from": s.RatingFrom,
		"rating_to":   s.RatingTo,
		"time":        s.Time,
	}
}

type StockResponse struct {
	Items    []map[string]interface{} `json:"items"`
	NextPage string                   `json:"next_page"`
}

type StockRecommendation struct {
	Stock          Stock   `json:"stock"`
	Score          float64 `json:"score"`
	Recommendation string  `json:"recommendation"`
}
