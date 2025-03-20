package models

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

type StockResponse struct {
	Items    []StockData `json:"items"`
	NextPage string      `json:"next_page"`
}

type StockData struct {
	Ticker     string `json:"ticker"`
	Company    string `json:"company"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type StockRecommendation struct {
	Stock          Stock   `json:"stock"`
	Score          float64 `json:"score"`
	Recommendation string  `json:"recommendation"`
}
