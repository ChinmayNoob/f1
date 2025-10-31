package utils

import "strconv"

const (
	DEFAULT_PAGE      = 1
	DEFAULT_PAGE_SIZE = 10
)

func ParsePagination(pageStr, limitStr string) (int, int) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = DEFAULT_PAGE
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = DEFAULT_PAGE_SIZE
	}

	return page, limit
}

func Paginate(query string, page, limit int) (string, error) {
	if page < 1 {
		page = DEFAULT_PAGE
	}
	if limit < 1 {
		limit = DEFAULT_PAGE_SIZE
	}
	offset := (page - 1) * limit
	return query + " LIMIT " + strconv.Itoa(limit) + " OFFSET " + strconv.Itoa(offset), nil
}
