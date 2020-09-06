package response

import "github.com/diegogomesaraujo/fund-manager-api/internal/domain"

// Pagination pagination dto
type Pagination struct {
	Page       int         `json:"page"`
	TotalPages int64       `json:"totalPages"`
	Total      int64       `json:"total"`
	Content    interface{} `json:"content"`
}

// FromPagination convert domain to dto
func (p *Pagination) FromPagination(domain domain.Pagination, content interface{}) {
	p.Page = domain.Page
	p.TotalPages = domain.TotalPages
	p.Total = domain.Total
	p.Content = content
}
