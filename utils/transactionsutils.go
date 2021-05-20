package utils

import (
	"errors"
	"rfgocore/utils/utilsstring"
	"rfgodata/constants"
	"rfgodata/transactions"
)

var debug bool = false

// SetDebug : Method for set debug queries transaction
func SetDebug(isDebug bool) {
	debug = isDebug
}

// Method to check is debug queries transaction
func IsDebug() bool {
	return debug
}

// GetTransactionInParams : method for get transaction in params
func GetTransactionInParams(mapParams *map[string]interface{}) (transactions.ITransaction, error) {
	var returnTransaction transactions.ITransaction = nil
	var returnError error = nil
	var paramExists bool = true

	if mapParams != nil {
		returnTransaction = (*mapParams)[constants.ParamTransaction].(transactions.ITransaction)

		// If not exist in mapparams return nil and return code error
		if returnTransaction == nil || !paramExists {
			returnTransaction = nil
			returnError = errors.New(utilsstring.IntToString(int(constants.ErrorTransactionNotFound)))
		}
	}

	return returnTransaction, returnError
}
