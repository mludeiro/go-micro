package database

type ResultSet struct {
	Data  []interface{}
	Page  uint
	Pages uint
	Total uint
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
}
