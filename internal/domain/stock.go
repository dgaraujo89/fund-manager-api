package domain

import "time"

type orderType string

// SELL order type
const SELL orderType = "SELL"

// BUY order type
const BUY orderType = "BUY"

// StockOrder domain
type StockOrder struct {
	Stock          string
	PurchasePrice  uint64
	targetPrice    uint64
	PurchaseDate   time.Time
	SalePrice      uint64
	SaleDate       time.Time
	StopPercentage float64
	Amount         uint64
	Finished       bool
	OrderType      orderType
}
