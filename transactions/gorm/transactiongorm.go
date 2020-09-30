package gorm

import (
	"rfgodata/transactions"

	"gorm.io/gorm"

	"rfgodata/beans/query"
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
func (transactionGorm TransactionGorm) List(tableName string, instaceModel func(dbContext interface{}) (interface{}, error), fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit) (interface{}, error) {
	return instaceModel(transactionGorm.Transaction.Table(tableName))
}

// RollBack : Method for execute rollback
func (transactionGorm TransactionGorm) RollBack() {
	transactionGorm.Transaction.Rollback()
}
