package transactions

import "rfgodata/beans/query"

// ITransaction : interface define methods for transactions
type ITransaction interface {

	// Add : method for add data
	Add(data interface{}) (interface{}, error)

	// Edit : method for edit data
	Edit(data interface{}) (interface{}, error)

	// Count : method for count data
	Count(tableName string, filters []query.Filter, joins []query.Join, groups []query.Group) (int64, error)

	// List : method for get list of data
	List(tableName string, instaceModel func(func(containerData interface{}) (interface{}, error)) (interface{}, error), fields []query.Field, filters []query.Filter, joins []query.Join, orders []query.Order, groups []query.Group, limit query.Limit) (interface{}, error)

	// RollBack : Method for execute rollback
	RollBack() error

	// Commit : Method for commit
	Commit() error

	// FinishTransaction: Method for finish transaction
	FinishTransaction(err error) error
}

// BaseTransaction : base struct for all transactions
type BaseTransaction struct {
}

// RollBack : Method for execute RollBack
func (baseTransaction BaseTransaction) RollBack() error {
	return nil
}

// Commit : Method for execute Commit
func (baseTransaction BaseTransaction) Commit() error {
	return nil
}

// Add : method for add data
func (baseTransaction BaseTransaction) Add(data interface{}) (interface{}, error) {
	return nil, nil
}

// Edit : method for edit data
func (baseTransaction BaseTransaction) Edit(data interface{}) (interface{}, error) {
	return nil, nil
}

// FinishTransaction : Method for finish transaction
func (baseTransaction BaseTransaction) FinishTransaction(err error) error {
	var errReturn error
	if err != nil {
		errReturn = baseTransaction.RollBack()
	} else {
		// TODO COMMIT if needed
		errReturn = baseTransaction.Commit()
	}
	return errReturn
}
