package repository

type ResultSet[T any] struct {
	Page  uint
	Pages uint
	Total int64
	Query
	Data  []T
	Error error
}

type Condition struct {
	Field      string
	Comparator string
	Value      string
}

type Query struct {
	Fetchs     []string
	Conditions []Condition
	OrderBy    []string
	PageSize   uint
	PageNumber uint
}

func (query Query) GetPageSize() uint {
	if query.PageSize > 1000 || query.PageSize == 0 {
		return 100
	} else {
		return query.PageSize
	}
}
