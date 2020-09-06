package usecases

import "github.com/diegogomesaraujo/fund-manager-api/internal/domain"

// ListStocksOrderPort a port to create a stock order
type ListStocksOrderPort interface {
	List(pageable domain.Pageable) (domain.Pagination, error)
}

// ListStockOrders list stock orders
func ListStockOrders(pageable domain.Pageable, port ListStocksOrderPort) (domain.Pagination, error) {
	return port.List(pageable)
}
