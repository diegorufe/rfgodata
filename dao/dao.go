package dao

import (
	"reflect"
	"rfgodata/beans/query"
)

// IDao : Interface define dao data operations
type IDao interface {
	Test() string
	// Edit : method for edit data
	Edit(data interface{}, mapParams *map[string]interface{}) (interface{}, error)

	// Add : method for save data
	Add(data interface{}, mapParams *map[string]interface{}) (interface{}, error)

	// Delete : method for delete data
	Delete(data interface{}, mapParams *map[string]interface{}) error

	// Count : method for count data
	Count(filters []query.Filter, joins []query.Join, groups []query.Group, mapParams *map[string]interface{}) (int64, error)

	// List : method for get list of data
	List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit, mapParams *map[string]interface{}) (interface{}, error)
}

// BaseDao is  base struct for daos
type BaseDao struct {
	TableName string
	TypeModel reflect.Type
}

// List : method for get list of data
func (dao BaseDao) List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit, mapParams *map[string]interface{}) (interface{}, error) {
	return nil, nil
}

// Count : method for count data
func (dao BaseDao) Count(filters []query.Filter, joins []query.Join, groups []query.Group, mapParams *map[string]interface{}) (int64, error) {
	return 0, nil
}

// Delete : method for delete data
func (dao BaseDao) Delete(data interface{}, mapParams *map[string]interface{}) error {
	return nil
}

// Edit : method for edit data in database
func (dao BaseDao) Edit(data interface{}, mapParams *map[string]interface{}) (interface{}, error) {
	return nil, nil
}

// Add : method for save data
func (dao BaseDao) Add(data interface{}, mapParams *map[string]interface{}) (interface{}, error) {
	return nil, nil
}
