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

	//  TODO apply wheres, joins ...
	res := transactionGorm.Transaction.Table(tableName).Count(&returnCount)

	return returnCount, res.Error
}

// List : method for get list of data
func (transactionGorm TransactionGorm) List(tableName string, instaceModel func(func(containerData interface{}) (interface{}, error)) (interface{}, error), fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit) (interface{}, error) {
	return instaceModel(func(containerData interface{}) (interface{}, error) {
		res := transactionGorm.Transaction.Table(tableName).Offset(limit.Start).Limit(limit.End).Find(containerData)
		return containerData, res.Error
	})
}

// RollBack : Method for execute rollback
func (transactionGorm TransactionGorm) RollBack() error {
	return transactionGorm.Transaction.Rollback().Error
}

// Commit : Method for execute Commit
func (transactionGorm TransactionGorm) Commit() error {
	return transactionGorm.Transaction.Commit().Error
}

// FinishTransaction : Method for finish transaction
func (transactionGorm TransactionGorm) FinishTransaction(err error) error {
	var errReturn error
	if err != nil {
		errReturn = transactionGorm.RollBack()
	} else {
		// TODO COMMIT if needed
		errReturn = transactionGorm.Commit()
	}
	return errReturn
}
