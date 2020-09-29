package service

import (
	"rfgodata/beans/query"
	"rfgodata/dao"
)

// IService : interface for all services
type IService interface {

	// Count : method for count data
	Count(filters []query.Filter, joins []query.Join, groups []query.Group, mapParams *map[string]interface{}) (int64, error)

	// List : method for get list of data
	List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit, mapParams *map[string]interface{}) ([]interface{}, error)
}

// BaseService is  base struct for services
type BaseService struct {
	Dao dao.IDao
}

// Count : method for count data
func (service BaseService) Count(filters []query.Filter, joins []query.Join, groups []query.Group, mapParams *map[string]interface{}) (int64, error) {
	return service.Dao.Count(filters, joins, groups, mapParams)
}

// List : method for get list of data
func (service BaseService) List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit, mapParams *map[string]interface{}) ([]interface{}, error) {
	return service.Dao.List(fields, filters, joins, orders, groups, limit, mapParams)
}
