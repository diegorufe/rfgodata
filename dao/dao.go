package dao

import (
	"rfgodata/beans/query"
)

// IDao : Interface define dao data operations
type IDao interface {
	// Edit : method for edit data
	Edit(data interface{}, mapParams *map[string]interface{}) (interface{}, error)

	// Add : method for save data
	Add(data interface{}, mapParams *map[string]interface{}) (interface{}, error)

	// Delete: method for delete data
	Delete(data interface{}, mapParams *map[string]interface{}) error

	// Count: method for count data
	Count(filters []query.Filter, joins []query.Join, limit rfdatalimit.Limit, mapParams *map[string]interface{}) (uint64, error)

	// List: method for get list of data
	List(fields []rfdatafield.Field, filters []query.Filter, joins []query.Join, limit query.Limit, mapParams *map[string]interface{}) ([]interface{}, error)
}
