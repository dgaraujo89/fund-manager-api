package usecases

import "github.com/diegogomesaraujo/fund-manager-api/internal/domain"

// FindStockOrderPort find a stock order port
type FindStockOrderPort interface {
	Find(id string) *domain.StockOrder
}

// Find a stock order
func Find(id string, findPort FindStockOrderPort) *domain.StockOrder {
	return findPort.Find(id)
}
