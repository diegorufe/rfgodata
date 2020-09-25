package transactions

// ITransaction : interface define methods for transactions
type ITransaction interface {
	// Edit : method for edit data
	Edit(data interface{}) (interface{}, error)

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
