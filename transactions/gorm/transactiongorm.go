package gorm

import (
	"rfgodata/transactions"

	"gorm.io/gorm"

	"rfgodata/beans/query"

	"rfgodata/models"
)

// TransactionGorm transaction type gorm
type TransactionGorm struct {
	transactions.BaseTransaction
	Transaction *gorm.DB
}

// Count : method for count data
func (transactionGorm TransactionGorm) Count(tableName string, filters []query.Filter, joins []query.Join, groups []query.Group) (int64, error) {
	var returnCount int64 = 0
	var returnError error = nil

	//  TODO apply wheres, joins ...
	transactionGorm.Transaction.Table(tableName).Count(&returnCount)

	return returnCount, returnError
}

// List : method for get list of data
func (transactionGorm TransactionGorm) List(instaceArrayMode []models.IModel, fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit) ([]models.IModel, error) {
	var returnData []models.IModel = instaceArrayMode
	var returnError error = nil

	//  TODO apply wheres, order, limit ...
	transactionGorm.Transaction.Find(&returnData)

	return returnData, returnError
}

// RollBack : Method for execute rollback
func (transactionGorm TransactionGorm) RollBack() {
	transactionGorm.Transaction.Rollback()
}
