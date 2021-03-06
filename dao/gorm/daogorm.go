package daogorm

import (
	"rfgodata/utils"

	"gorm.io/gorm"

	"rfgodata/beans/query"

	"rfgodata/dao"
)

// DaoGorm : dao for gorm
type DaoGorm struct {
	dao.BaseDao
}

// Add : method for save data
func (daoGorm DaoGorm) Add(data interface{}, mapParams *map[string]interface{}) (interface{}, error) {
	var returnData interface{} = nil
	var returnError error = nil

	if data != nil {
		// Added updatedat and createdat
		utils.AddCreatedAt(data)
		utils.AddUpdatedAt(data)

		// find transaction
		transaction, returnError := utils.GetTransactionInParams(mapParams)

		if returnError == nil {
			returnData, returnError = (transaction).Add(data)
		}

	} else {
		returnError = gorm.ErrInvalidData
	}

	return returnData, returnError
}

// Edit : method for edit data in database
func (daoGorm DaoGorm) Edit(data interface{}, mapParams *map[string]interface{}) (interface{}, error) {
	var returnData interface{} = nil
	var returnErrorEdit error = nil

	if data != nil {

		// Added updated at
		utils.AddUpdatedAt(data)

		// find transaction
		transaction, returnError := utils.GetTransactionInParams(mapParams)

		// If has not error edit data
		if returnError == nil {
			returnData, returnErrorEdit = (transaction).Edit(data)
		} else {
			returnErrorEdit = returnError
		}

	} else {
		returnErrorEdit = gorm.ErrInvalidData
	}

	return returnData, returnErrorEdit
}

// Delete : method for delete data
func (daoGorm DaoGorm) Delete(data interface{}, mapParams *map[string]interface{}) error {
	var returnError error = nil

	if data != nil {
		// find transaction
		transaction, returnError := utils.GetTransactionInParams(mapParams)

		// If has not error edit data
		if returnError == nil {
			returnError = (transaction).Delete(data)
		}

	} else {
		returnError = gorm.ErrInvalidData
	}

	return returnError
}

// Count : method for count data
func (daoGorm DaoGorm) Count(filters []query.Filter, joins []query.Join, groups []query.Group, mapParams *map[string]interface{}) (int64, error) {
	var returnCount int64 = 0
	var returnError error = nil

	// find transaction
	transaction, returnError := utils.GetTransactionInParams(mapParams)

	if transaction != nil && returnError == nil {
		// Count
		returnCount, returnError = (transaction).Count(daoGorm.TableName, filters, joins, groups)
	}

	return returnCount, returnError
}

// List : method for get list of data
func (daoGorm DaoGorm) List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit, mapParams *map[string]interface{}) (interface{}, error) {
	var returnData interface{} = nil
	var returnError error = nil

	// find transaction
	transaction, returnError := utils.GetTransactionInParams(mapParams)

	if transaction != nil && returnError == nil {
		returnData, returnError = (transaction).List(daoGorm.TableName, daoGorm.TypeModel, fields, filters, joins, orders, groups, limit)
	} else {
		returnError = gorm.ErrInvalidTransaction
	}

	return returnData, returnError
}
