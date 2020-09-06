package response

import (
	"time"

	"github.com/diegogomesaraujo/fund-manager-api/internal/domain"
)

// StockOrder response dto
type StockOrder struct {
	ID             string    `json:"id,omitempty"`
	Stock          string    `json:"stock,omitempty"`
	PurchasePrice  float64   `json:"purchasePrice,omitempty"`
	TargetPrice    float64   `json:"targetPrice,omitempty"`
	PurchaseDate   time.Time `json:"purchaseDate,omitempty"`
	SalePrice      float64   `json:"salePrice,omitempty"`
	SaleDate       time.Time `json:"saleDate,omitempty"`
	StopPercentage float64   `json:"stopPercentage,omitempty"`
	Amount         uint64    `json:"amount,omitempty"`
	Finished       bool      `json:"finished,omitempty"`
	OrderType      string    `json:"orderType,omitempty"`
}

// FromStockOrder convert a domain.StockOrder to dto
func (s *StockOrder) FromStockOrder(stockOrder domain.StockOrder) {
	s.ID = stockOrder.ID
	s.Stock = stockOrder.Stock
	s.PurchasePrice = float64(stockOrder.PurchasePrice) / 100.0
	s.TargetPrice = float64(stockOrder.TargetPrice) / 100.0
	s.PurchaseDate = stockOrder.PurchaseDate
	s.SalePrice = float64(stockOrder.SalePrice) / 100.0
	s.SaleDate = stockOrder.SaleDate
	s.StopPercentage = stockOrder.StopPercentage
	s.Amount = stockOrder.Amount
	s.Finished = stockOrder.Finished
	s.OrderType = stockOrder.OrderType
}
