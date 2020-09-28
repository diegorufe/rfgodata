package transactions

import "rfgodata/beans/query"

// ITransaction : interface define methods for transactions
type ITransaction interface {
	// Edit : method for edit data
	Edit(data interface{}) (interface{}, error)

	// Count : method for count data
	Count(tableName string, filters []query.Filter, joins []query.Join, groups []query.Group) (int64, error)

	// List : method for get list of data
	List(fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit) ([]interface{}, error)

	// RollBack : Method for execute rollback
	RollBack()

	// FinishTransaction: Method for finish transaction
	FinishTransaction(err error)
}

// BaseTransaction: base struct for all transactions
type BaseTransaction struct {
}

// RollBack : Method for execute RollBack
func (baseTransaction *BaseTransaction) RollBack() {

}

// FinishTransaction: Method for finish transaction
func (baseTransaction *BaseTransaction) FinishTransaction(err error) {
	if err != nil {
		baseTransaction.RollBack()
	} else {
		// TODO COMMIT if needed
	}
}
