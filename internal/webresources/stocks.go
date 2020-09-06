package webresources

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/diegogomesaraujo/fund-manager-api/internal/domain"
	"github.com/diegogomesaraujo/fund-manager-api/internal/domain/usecases"
	"github.com/diegogomesaraujo/fund-manager-api/internal/repository"
	"github.com/diegogomesaraujo/fund-manager-api/internal/webresources/dto/request"
	"github.com/diegogomesaraujo/fund-manager-api/internal/webresources/dto/response"
)

// StocksRoutes to home enpoints
var StocksRoutes = Routes{
	Route{
		Name:        "List stock orders",
		Method:      "GET",
		Path:        "/stock-orders",
		HandlerFunc: list,
	},
	Route{
		Name:        "Find a stock order",
		Method:      "GET",
		Path:        "/stock-orders/{id}",
		HandlerFunc: find,
	},
	Route{
		Name:        "Create a stock order",
		Method:      "POST",
		Path:        "/stock-orders",
		HandlerFunc: create,
	},
	Route{
		Name:        "Delete a stock order",
		Method:      "DELETE",
		Path:        "/stock-orders/{id}",
		HandlerFunc: delete,
	},
}

func list(w http.ResponseWriter, r *http.Request) {
	var err error
	stockOrderRepository := &repository.StockOrderRepositoryPort{}

	pageable := domain.Pageable{}
	pageable.Page = 0
	pageable.Size = 20

	size := r.URL.Query().Get("size")
	page := r.URL.Query().Get("page")

	if size != "" {
		pageable.Size, err = strconv.Atoi(size)
	}

	if page != "" {
		pageable.Page, err = strconv.Atoi(page)
	}

	pagination, err := usecases.ListStockOrders(pageable, stockOrderRepository)

	if err != nil {
		log.Println("ERROR:", err)
		handleError(w, "Error", 500)
		return
	}

	content := pagination.Content.([]domain.StockOrder)
	paginationDto := response.Pagination{}
	paginationDto.FromPagination(pagination, toStockOrders(content))

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(paginationDto)
}

func create(w http.ResponseWriter, r *http.Request) {
	var createStockOrderDto request.CreateStockOrder

	if readBodyFromJSON(w, r, &createStockOrderDto) != nil {
		return
	}

	stockOrderRepository := &repository.StockOrderRepositoryPort{}

	usecases.CreateStockOrder(createStockOrderDto.ToStockOrder(), stockOrderRepository)

	w.WriteHeader(201)
}

func find(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	stockOrderRepository := &repository.StockOrderRepositoryPort{}
	stockOrder := usecases.Find(id, stockOrderRepository)

	if stockOrder != nil {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(stockOrder)
		return
	}

	w.WriteHeader(404)
}

func delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	stockOrderRepository := &repository.StockOrderRepositoryPort{}
	if err := usecases.Remove(id, stockOrderRepository); err != nil {
		log.Println("ERROR:", err)
		handleError(w, "Error", 500)
		return
	}

	w.WriteHeader(204)
}

func toStockOrders(stockOrders []domain.StockOrder) []response.StockOrder {
	stocks := make([]response.StockOrder, len(stockOrders))

	for i, stockOrder := range stockOrders {
		stocks[i] = response.StockOrder{}
		stocks[i].FromStockOrder(stockOrder)
	}

	return stocks
}
