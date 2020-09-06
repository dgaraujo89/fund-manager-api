package domain

// Pagination pagination
type Pagination struct {
	Page       int
	TotalPages int64
	Total      int64
	Content    interface{}
}

// Pageable pagination criteria
type Pageable struct {
	Page int
	Size int
}
