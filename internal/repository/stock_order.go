package repository

import (
	"log"

	"github.com/google/uuid"

	"github.com/diegogomesaraujo/fund-manager-api/internal/db"
	"github.com/diegogomesaraujo/fund-manager-api/internal/domain"
	"github.com/diegogomesaraujo/fund-manager-api/internal/repository/entity"
)

// StockOrderRepositoryPort port impl
type StockOrderRepositoryPort struct{}

// Create a stock order
func (r *StockOrderRepositoryPort) Create(stockOrder domain.StockOrder) error {
	stockOrderEntity := entity.StockOrder{}
	stockOrderEntity.FromStockOrder(stockOrder)

	db, err := db.GetDb()
	if err != nil {
		log.Println("Erro when getting the database reference", err)
		return err
	}

	stockOrderEntity.ID = uuid.New().String()

	db.Create(stockOrderEntity)

	return nil
}

// List stock orders
func (r *StockOrderRepositoryPort) List(pageable domain.Pageable) (domain.Pagination, error) {
	db, err := db.GetDb()
	if err != nil {
		log.Println("Erro when getting the database reference", err)
		return domain.Pagination{}, err
	}

	var total int64
	db.Model(&entity.StockOrder{}).Count(&total)

	var stockOrders []entity.StockOrder

	db.Model(&entity.StockOrder{}).
		Offset(pageable.Page * pageable.Size).
		Limit(pageable.Size).
		Find(&stockOrders)

	pagination := domain.Pagination{
		Page:       pageable.Page,
		TotalPages: calculateTotalPages(int64(pageable.Size), total),
		Total:      total,
		Content:    toStockOrders(stockOrders),
	}

	return pagination, nil
}

// Find by id
func (r *StockOrderRepositoryPort) Find(id string) *domain.StockOrder {
	db, err := db.GetDb()
	if err != nil {
		log.Println("Erro when getting the database reference", err)
		return nil
	}

	var stockOrders []entity.StockOrder

	db.Find(&stockOrders, "id = ?", id)

	if len(stockOrders) == 0 {
		return nil
	}

	domain := stockOrders[0].ToStockOrder()

	return &domain
}

// Remove by id
func (r *StockOrderRepositoryPort) Remove(id string) error {
	db, err := db.GetDb()
	if err != nil {
		log.Println("Erro when getting the database reference", err)
		return err
	}

	var stockOrder entity.StockOrder
	db.Delete(&stockOrder, "id = ?", id)

	return nil
}

func toStockOrders(entities []entity.StockOrder) []domain.StockOrder {
	stocks := make([]domain.StockOrder, len(entities))

	for i, entity := range entities {
		stocks[i] = entity.ToStockOrder()
	}

	return stocks
}
