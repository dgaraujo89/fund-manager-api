package usecases

import "github.com/diegogomesaraujo/fund-manager-api/internal/domain"

// CreateStockOrderPort a port to create a stock order
type CreateStockOrderPort interface {
	Create(stockOrder domain.StockOrder)
}

// CreateStockOrder register a order
func CreateStockOrder(stockOrder domain.StockOrder, createStockOrderPort CreateStockOrderPort) {
	createStockOrderPort.Create(stockOrder)
}
