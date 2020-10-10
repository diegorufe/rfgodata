package query

import "rfgodata/constants/query"

// Order : for apply in querys
type Order struct {
	Field     string          `json:"field"`
	Alias     string          `json:"alias"`
	OrderType query.OrderType `json:"orderType"`
}
