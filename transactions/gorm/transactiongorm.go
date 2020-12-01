package gorm

import (
	"reflect"
	"rfgodata/transactions"

	"gorm.io/gorm"

	"rfgodata/beans/query"

	trxGormUtils "rfgodata/utils/gorm"
)

// TransactionGorm transaction type gorm
type TransactionGorm struct {
	transactions.BaseTransaction
	Transaction *gorm.DB
}

// Count : method for count data
func (transactionGorm TransactionGorm) Count(tableName string, filters []query.Filter, joins []query.Join, groups []query.Group) (int64, error) {
	var returnCount int64 = 0
	db := trxGormUtils.ApplyWhere(transactionGorm.Transaction.Table(tableName), filters)
	db = trxGormUtils.ApplyJoins(db, joins)
	res := db.Count(&returnCount)

	return returnCount, res.Error
}

// List : method for get list of data
func (transactionGorm TransactionGorm) List(tableName string, modelType reflect.Type, fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit) (interface{}, error) {
	db := trxGormUtils.ApplySelect(transactionGorm.Transaction.Table(tableName+" "+trxGormUtils.DefaultAliasQuery), fields)
	db = trxGormUtils.ApplyWhere(db, filters)
	db = trxGormUtils.ApplyJoins(db, joins)
	db = trxGormUtils.ApplyOrders(db, orders)
	db = trxGormUtils.ApplyLimit(db, limit)

	return trxGormUtils.RawData(db, modelType)
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
		errReturn = transactionGorm.Commit()
	}
	return errReturn
}
