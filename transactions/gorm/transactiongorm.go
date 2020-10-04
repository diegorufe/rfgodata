package gorm

import (
	"rfgodata/transactions"

	"gorm.io/gorm"

	"rfgodata/beans/query"

	"rfgocore/utils/utilsstring"

	queryconstants "rfgodata/constants/query"
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

func (transactionGorm TransactionGorm) applyWhereQueryBuilder(filters []query.Filter, firstLevel bool, valuesQuery *[]interface{}) (string, []interface{}) {
	var queryBuilder string = ""

	if filters != nil && len(filters) > 0 {
		for index, filter := range filters {
			queryBuilder = queryBuilder + " "
			if !(firstLevel && index == 0) {
				// And / or operation type
				switch filter.FilterOperationType {

				case queryconstants.And:
					queryBuilder = queryBuilder + " AND "
					break

				case queryconstants.Or:
					queryBuilder = queryBuilder + " OR "
					break
				}

			}

			// open brackets
			if filter.OpenBrackets > 0 {
				for i := 0; i < filter.OpenBrackets; i++ {
					queryBuilder = queryBuilder + " ( "
				}
			}

			// Alias
			if utilsstring.IsNotEmpty(filter.Alias) {
				queryBuilder = queryBuilder + " " + filter.Alias + "."
			}

			// Field
			if utilsstring.IsNotEmpty(filter.Field) {
				queryBuilder = queryBuilder + filter.Field
			}

			// Filter type operation
			switch filter.FilterType {
			case queryconstants.Equal:
				queryBuilder = queryBuilder + " = ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
			}

			// close brackets
			if filter.CloseBrackets > 0 {
				for i := 0; i < filter.CloseBrackets; i++ {
					queryBuilder = queryBuilder + " ) "
				}
			}
		}

	}

	return queryBuilder, *valuesQuery
}

func (transactionGorm TransactionGorm) applyWhere(db *gorm.DB, filters []query.Filter) *gorm.DB {
	var dbReturn *gorm.DB
	var queryBuilder string = ""
	var valuesQuery []interface{}

	queryBuilder, valuesQuery = transactionGorm.applyWhereQueryBuilder(filters, true, &valuesQuery)

	// If have query call where
	if utilsstring.IsNotEmpty(queryBuilder) {
		dbReturn = db.Where(queryBuilder, valuesQuery)
	}
	return dbReturn
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
