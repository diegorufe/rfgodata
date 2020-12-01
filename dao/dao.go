package dao

import (
	"reflect"
	"rfgodata/beans/query"
)

// InstanceModelArrayFunc For genereate instace array model
type InstanceModelArrayFunc func(executeFunction func(containerData interface{}) (interface{}, error)) (interface{}, error)

// InstanceModelFunc For genereate instace model
type InstanceModelFunc func(executeFunction func(containerData interface{}) (interface{}, error)) (interface{}, error)

// IDao : Interface define dao data operations
type IDao interface {
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
	TableName              string
	InstanceFindModelArray InstanceModelArrayFunc
	InstanceFindModel      InstanceModelFunc
	TypeModel              reflect.Type
}
