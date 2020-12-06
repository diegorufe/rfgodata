package utils

import (
	"rfgodata/beans/query"
	querycst "rfgodata/constants/query"
)

// GetFilterSelectPk : method for get field select with pk from table.
func GetFilterSelectPk(pk interface{}) query.Filter {
	return query.Filter{FilterType: querycst.LiteralEqual, FilterOperationType: querycst.And, Field: "id", Value: pk}
}

// GetFieldSelectAll : method for get field select all values. DONT include joins asociations
func GetFieldSelectAll() query.Field {
	return query.Field{Name: "*"}
}
