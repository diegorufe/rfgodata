package gorm

import (
	"gorm.io/gorm"
)

// TransactionGorm transaction type gorm
type TransactionGorm struct {
	Transaciont *gorm.DB
}

// Edit : method for edit data in database
func (transactionGorm TransactionGorm) Edit(data interface{}) (interface{}, error) {
	var returnError error = nil
	transactionGorm.Transaciont.Edit(&data)
	if transactionGorm.Transaciont.Error != nil {
		returnError = transactionGorm.Transaciont.Error
	}
	return data, returnError
}
