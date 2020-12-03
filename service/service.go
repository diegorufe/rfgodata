package service

import (
	"rfgodata/beans/query"
	querycst "rfgodata/constants/query"
	"rfgodata/dao"
)

// IService : interface for all services
type IService interface {

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

	// LoadNew : method for load new data
	LoadNew(mapParams *map[string]interface{}) (interface{}, error)

	// Read: Method for read entity
	Read(pk interface{}, mapParams *map[string]interface{}) (interface{}, error)
}

// BaseService is  base struct for services
type BaseService struct {
	Dao dao.IDao
}

// Add : method for save data
func (service BaseService) Add(data interface{}, mapParams *map[string]interface{}) (interface{}, error) {
	return service.Dao.Add(data, mapParams)
}

// Edit : method for edit data in database
func (service BaseService) Edit(data interface{}, mapParams *map[string]interface{}) (interface{}, error) {
	return service.Dao.Edit(data, mapParams)
}

// Delete : method for delete data
func (service BaseService) Delete(data interface{}, mapParams *map[string]interface{}) error {
	return service.Dao.Delete(data, mapParams)
}

// Count : method for count data
func (service BaseService) Count(filters []query.Filter, joins []query.Join, groups []query.Group, mapParams *map[string]interface{}) (int64, error) {
	return service.Dao.Count(filters, joins, groups, mapParams)
}

// List : method for get list of data
func (service BaseService) List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit, mapParams *map[string]interface{}) (interface{}, error) {
	return service.Dao.List(fields, filters, joins, orders, groups, limit, mapParams)
}

// Read: Method for read entity
func (service BaseService) Read(pk interface{}, mapParams *map[string]interface{}) (interface{}, error) {
	var data interface{}
	filters := make([]query.Filter, 1)
	fields := make([]query.Field, 1)

	filters[0] = query.Filter{FilterType: querycst.Equal, FilterOperationType: querycst.And, Field: "id", Value: pk}
	fields[0] = query.Field{Name: "*"}

	arrayData, err := service.List(fields, filters, nil, nil, nil, query.Limit{Start: 0, End: 1}, mapParams)

	if err == nil && arrayData != nil && len((arrayData.([]interface{}))) > 0 {
		data = arrayData.([]interface{})[0]
	}

	return data, err
}

// LoadNew : method for load new data
func (service BaseService) LoadNew(mapParams *map[string]interface{}) (interface{}, error) {
	return nil, nil
}
