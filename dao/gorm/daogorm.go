package rfdatadaogorm

import (
	"rfgodata/utils"

	"gorm.io/gorm"
)

// DaoGorm : dao for gorm
type DaoGorm struct {
}

// Edit : method for edit data in database
func (dao *DaoGorm) Edit(data interface{}, mapParams *map[string]interface{}) (interface{}, error) {
	var returnData interface{} = nil
	var returnError error = nil

	if data != nil {
		// find transaction
		transaction, returnError := utils.GetTransactionInParams(mapParams)

		// If has not error edit data
		if returnError != nil {
			returnData, returnError = transaction.Edit(data)
		}

	} else {
		returnError = gorm.ErrInvalidData
	}

	return returnData, returnError
}
