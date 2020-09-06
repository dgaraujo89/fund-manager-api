package repository

func calculateTotalPages(size int64, total int64) int64 {
	if total <= size {
		return 1
	}

	pages := total / size

	if total%size > 0 {
		pages++
	}

	return pages
}
