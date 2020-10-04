package query

import (
	"rfgodata/constants/query"
)

// Filter :  for apply in querys
type Filter struct {
	FilterType          query.FilterType          `json:"filterType"`
	FilterOperationType query.FilterOperationType `json:"filterOperationType"`
	Field               string                    `json:"filed"`
	Alias               string                    `json:"alias"`
	Value               interface{}               `json:"value"`
	CollecionFilters    []Filter                  `json:"collecionFilters"`
	OpenBrackets        int                       `json:"openBrackets"`
	CloseBrackets       int                       `json:"closeBrackets"`
}
