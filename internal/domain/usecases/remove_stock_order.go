package usecases

// RemoveStockOrderPort remove a stock order
type RemoveStockOrderPort interface {
	Remove(id string) error
}

// Remove a stock order
func Remove(id string, removePort RemoveStockOrderPort) error {
	return removePort.Remove(id)
}
