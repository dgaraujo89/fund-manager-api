package request

import (
	"time"

	"github.com/diegogomesaraujo/fund-manager-api/internal/domain"
)

// CreateStockOrder response dto
type CreateStockOrder struct {
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

// ToStockOrder convert to domain.StockOrder
func (s *CreateStockOrder) ToStockOrder() domain.StockOrder {
	stockOrder := domain.StockOrder{}

	stockOrder.Stock = s.Stock
	stockOrder.PurchasePrice = uint64(s.PurchasePrice * 100.0)
	stockOrder.TargetPrice = uint64(s.TargetPrice * 100.0)
	stockOrder.PurchaseDate = s.PurchaseDate
	stockOrder.SalePrice = uint64(s.SalePrice * 100.0)
	stockOrder.SaleDate = s.SaleDate
	stockOrder.StopPercentage = s.StopPercentage
	stockOrder.Amount = s.Amount
	stockOrder.Finished = s.Finished
	stockOrder.OrderType = s.OrderType

	return stockOrder
}
