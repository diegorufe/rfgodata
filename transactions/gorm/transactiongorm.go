package gorm

import (
	"rfgodata/transactions"

	"gorm.io/gorm"
)

// TransactionGorm transaction type gorm
type TransactionGorm struct {
	*transactions.BaseTransaction
	Transaction *gorm.DB
}

// Edit : method for edit data in database
func (transactionGorm *TransactionGorm) Edit(data interface{}) (interface{}, error) {
	var returnError error = nil
	transactionGorm.Transaction.Save(&data)
	if transactionGorm.Transaction.Error != nil {
		returnError = transactionGorm.Transaction.Error
	}

	// Finish transaction
	transactionGorm.FinishTransaction(returnError)

	return data, returnError
}

// RollBack : Method for execute rollback
func (transactionGorm *TransactionGorm) RollBack() {
	transactionGorm.Transaction.Rollback()
}
