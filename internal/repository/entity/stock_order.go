package entity

import (
	"time"

	"github.com/diegogomesaraujo/fund-manager-api/internal/domain"
)

var layout = "2006-01-01 00:00:00"
var defaultDate = "1900-01-01 00:00:00"

// StockOrder entity
type StockOrder struct {
	ID             string  `gorm:"column:id;primaryKey"`
	Stock          string  `gorm:"column:stock"`
	PurchasePrice  uint64  `gorm:"column:purchase_price"`
	TargetPrice    uint64  `gorm:"column:target_price"`
	PurchaseDate   string  `gorm:"column:purchase_date"`
	SalePrice      uint64  `gorm:"column:sale_price"`
	SaleDate       string  `gorm:"column:sale_date"`
	StopPercentage float64 `gorm:"column:stop_percentage"`
	Amount         uint64  `gorm:"column:amount"`
	Finished       string  `gorm:"column:finished"`
	OrderType      string  `gorm:"column:order_type"`
}

// TableName table name
func (s *StockOrder) TableName() string {
	return "stock_order"
}

// FromStockOrder convert from domain
func (s *StockOrder) FromStockOrder(stockOrder domain.StockOrder) {
	s.ID = stockOrder.ID
	s.Stock = stockOrder.Stock
	s.PurchasePrice = stockOrder.PurchasePrice
	s.TargetPrice = stockOrder.TargetPrice
	s.PurchaseDate = stockOrder.PurchaseDate.Format(layout)
	s.SalePrice = stockOrder.SalePrice
	s.SaleDate = stockOrder.SaleDate.Format(layout)
	s.StopPercentage = stockOrder.StopPercentage
	s.Amount = stockOrder.Amount
	s.OrderType = string(stockOrder.OrderType)

	s.Finished = "N"

	if stockOrder.Finished {
		s.Finished = "Y"
	}
}

// ToStockOrder convert to domain
func (s *StockOrder) ToStockOrder() domain.StockOrder {
	stockOrder := domain.StockOrder{}

	finished := false
	if s.Finished == "Y" {
		finished = true
	}

	stockOrder.ID = s.ID
	stockOrder.Stock = s.Stock
	stockOrder.PurchasePrice = s.PurchasePrice
	stockOrder.TargetPrice = s.TargetPrice
	stockOrder.SalePrice = s.SalePrice
	stockOrder.StopPercentage = s.StopPercentage
	stockOrder.Amount = s.Amount
	stockOrder.Finished = finished
	stockOrder.OrderType = s.OrderType

	purchaseDate, err := time.Parse(layout, s.PurchaseDate)
	if err == nil {
		stockOrder.PurchaseDate = purchaseDate
	}

	if s.SaleDate != defaultDate {
		saleDate, err := time.Parse(layout, s.SaleDate)
		if err == nil {
			stockOrder.SaleDate = saleDate
		}
	}

	return stockOrder
}
