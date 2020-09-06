package domain

import "time"

// StockOrder domain
type StockOrder struct {
	ID             string
	Stock          string
	PurchasePrice  uint64
	TargetPrice    uint64
	PurchaseDate   time.Time
	SalePrice      uint64
	SaleDate       time.Time
	StopPercentage float64
	Amount         uint64
	Finished       bool
	OrderType      string
}
