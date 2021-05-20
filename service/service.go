package service

import (
	"math"
	"reflect"
	"rfgodata/beans/core"
	databcore "rfgodata/beans/core"
	"rfgodata/beans/query"
	"rfgodata/dao"
	"rfgodata/utils"
)

// IService : interface for all services
type IService interface {

	// Edit : method for edit data
	Edit(data interface{}, mapParams *map[string]interface{}) core.ResponseService

	// Add : method for save data
	Add(data interface{}, mapParams *map[string]interface{}) core.ResponseService

	// Delete : method for delete data
	Delete(data interface{}, mapParams *map[string]interface{}) core.ResponseService

	// Count : method for count data
	Count(filters []query.Filter, joins []query.Join, groups []query.Group, mapParams *map[string]interface{}) core.ResponseService

	// List : method for get list of data
	List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit, mapParams *map[string]interface{}) core.ResponseService

	// Browse : method for get count and list  data
	Browser(requestBrowse core.RequestBrowser, mapParams *map[string]interface{}) core.ResponseService

	// LoadNew : method for load new data
	LoadNew(mapParams *map[string]interface{}) core.ResponseService

	// Read: Method for read entity
	Read(pk interface{}, mapParams *map[string]interface{}) core.ResponseService

	// GetTypeModel: Method for get type molde
	GetTypeModel() reflect.Type

	// FindByPk : method for find entity by pk
	FindByPk(pk interface{}, fields []query.Field, joins []query.Join, mapParams *map[string]interface{}) core.ResponseService
}

// BaseService is  base struct for services
type BaseService struct {
	Dao dao.IDao
}

// Add : method for save data
func (service BaseService) Add(data interface{}, mapParams *map[string]interface{}) core.ResponseService {
	var responseService core.ResponseService = core.ResponseService{}
	dataResponseDao, err := service.Dao.Add(data, mapParams)
	responseService.Data = dataResponseDao
	responseService.ResponseError = err
	return responseService
}

// Edit : method for edit data in database
func (service BaseService) Edit(data interface{}, mapParams *map[string]interface{}) core.ResponseService {
	var responseService core.ResponseService = core.ResponseService{}
	dataResponseDao, err := service.Dao.Edit(data, mapParams)
	// Set data in response service
	responseService.Data = dataResponseDao
	responseService.ResponseError = err
	return responseService
}

// Delete : method for delete data
func (service BaseService) Delete(data interface{}, mapParams *map[string]interface{}) core.ResponseService {
	var responseService core.ResponseService = core.ResponseService{}
	responseService.ResponseError = service.Dao.Delete(data, mapParams)
	return responseService
}

// Count : method for count data
func (service BaseService) Count(filters []query.Filter, joins []query.Join, groups []query.Group, mapParams *map[string]interface{}) core.ResponseService {
	var responseService core.ResponseService = core.ResponseService{}
	dataResponseDao, err := service.Dao.Count(filters, joins, groups, mapParams)
	responseService.Data = dataResponseDao
	responseService.ResponseError = err
	return responseService
}

// List : method for get list of data
func (service BaseService) List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit, mapParams *map[string]interface{}) core.ResponseService {
	var responseService core.ResponseService = core.ResponseService{}
	dataResponseDao, err := service.Dao.List(fields, filters, joins, orders, groups, limit, mapParams)
	responseService.Data = dataResponseDao
	responseService.ResponseError = err
	return responseService
}

// Browser : method for get count and list  data
func (service BaseService) Browser(requestBrowse core.RequestBrowser, mapParams *map[string]interface{}) core.ResponseService {
	var responseService core.ResponseService = core.ResponseService{}
	var responseBrowser databcore.ResponseBrowser = databcore.ResponseBrowser{}

	// Firt step count data
	dataResponseDao, err := service.Dao.Count(requestBrowse.Filters, requestBrowse.Joins, nil, mapParams)

	responseService.ResponseError = err

	if err == nil && dataResponseDao > 0 {

		responseBrowser.Count = dataResponseDao

		// Second step list data
		dataResponseDao, err := service.Dao.List(requestBrowse.Fields, requestBrowse.Filters, requestBrowse.Joins, requestBrowse.Orders, nil, service.CalculateLimitBrowser(int(responseBrowser.Count), requestBrowse.First, requestBrowse.RecordsPage), mapParams)

		responseBrowser.Data = dataResponseDao
		responseService.Data = responseBrowser
		responseService.ResponseError = err
	}

	return responseService
}

// Read: Method for read entity
func (service BaseService) Read(pk interface{}, mapParams *map[string]interface{}) core.ResponseService {
	var responseService core.ResponseService = core.ResponseService{}
	var data interface{}
	filters := make([]query.Filter, 1)
	fields := make([]query.Field, 1)

	filters[0] = utils.GetFilterSelectPk(pk)
	fields[0] = utils.GetFieldSelectAll()

	responseServiceList := service.List(fields, filters, nil, nil, nil, query.Limit{Start: 0, End: 1}, mapParams)

	if responseServiceList.ResponseError == nil && responseServiceList.Data != nil {
		data = responseServiceList.Data
	}

	responseService.Data = data
	responseService.ResponseError = responseServiceList.ResponseError

	return responseService
}

// LoadNew : method for load new data
func (service BaseService) LoadNew(mapParams *map[string]interface{}) core.ResponseService {
	return core.ResponseService{}
}

// CalculateLimitBrowser : method for calculate limit browser
func (service BaseService) CalculateLimitBrowser(totalRecors int, first int, recordsPage int) query.Limit {
	var numberOfPages float64 = math.Ceil(float64(totalRecors) / float64(recordsPage))
	var page float64 = 0

	if first > -1 {
		page = math.Ceil(float64(first) / float64(recordsPage))
	}

	if page > numberOfPages {
		page = 1
	}

	if page < 1 {
		page = 1
	}

	return query.Limit{Start: ((int(page) - 1) * recordsPage), End: recordsPage}
}

// FindByPk : method for find entity by pk
func (service BaseService) FindByPk(pk interface{}, fields []query.Field, joins []query.Join, mapParams *map[string]interface{}) core.ResponseService {
	var data interface{}

	if len(fields) == 0 {
		fields := make([]query.Field, 1)
		// all fields serie
		fields[0] = utils.GetFieldSelectAll()
	}

	filters := make([]query.Filter, 1)

	filters[0] = utils.GetFilterSelectPk(pk)

	responseService := service.List(fields, filters, joins, nil, nil, query.Limit{Start: 0, End: 1}, mapParams)

	if responseService.ResponseError == nil && responseService.Data != nil && len((responseService.Data.([]interface{}))) > 0 {
		data = responseService.Data.([]interface{})[0]
	}

	return databcore.ResponseService{Data: data, ResponseError: responseService.ResponseError}
}

// GetTypeModel : Method for get type molde
func (service BaseService) GetTypeModel() reflect.Type {
	return service.Dao.GetTypeModel()
}
