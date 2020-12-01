package gorm

import (
	"container/list"
	"reflect"
	"rfgodata/transactions"
	"strings"

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

	var arrayData []interface{}

	rows, err := db.Rows()

	if err == nil {
		colsRows, _ := rows.Columns()

		var containerList list.List
		var instanceModel interface{}

		for rows.Next() {

			// Instace ponter of model
			resultInstanceModel := reflect.New(modelType)
			instanceModel = resultInstanceModel.Interface()

			// Create a slice of interface{}'s to represent each column,
			// and a second slice to contain pointers to each item in the columns slice.
			//columns := make([]interface{}, len(colsRows))
			columnPointers := make([]interface{}, len(colsRows))
			instaceModelColumnReflect := reflect.ValueOf(instanceModel).Elem()

			for index, columKey := range colsRows {

				titleKey := strings.Title(columKey)

				if titleKey == "Id" {
					titleKey = "ID"
				}

				fieldColum := instaceModelColumnReflect.FieldByName(titleKey)
				interfacePointer := fieldColum.Addr().Interface()
				columnPointers[index] = interfacePointer
			}

			// Scan the result into the column pointers...
			// TODO time error
			if err = rows.Scan(columnPointers...); err != nil {
				break
			}

			for index, columKey := range colsRows {

				titleKey := strings.Title(columKey)

				if titleKey == "Id" {
					titleKey = "ID"
					//reflect.ValueOf(instanceModel).Elem().FieldByName(titleKey).SetUint(returnVal.(uint64))
				}

				valueColumn := columnPointers[index].(*interface{})

				//fmt.Print(typeof(*val))

				returnVal := *valueColumn

				switch returnVal.(type) {
				case []uint8:
					returnVal = string((returnVal).([]uint8))
					break
				default:
					break
				}

				valueReflect := reflect.ValueOf(instanceModel).Elem().FieldByName(titleKey)

				switch returnVal.(type) {
				case string:
					valueReflect.SetString(returnVal.(string))
					break
				case int:
				case int16:
				case int32:
				case int64:
					valueReflect.SetInt(returnVal.(int64))
					break

				case bool:
					valueReflect.SetBool(returnVal.(bool))
					break
				}

			}

			containerList.PushBack(instanceModel)
		}

		// If list contain data create array return
		if containerList.Len() > 0 {

		}
	}

	return arrayData, err
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
