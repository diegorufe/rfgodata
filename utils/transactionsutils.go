package rfdatatransactionsutils

import (
	"errors"
	"rfgocore/utils/utilsstring"
	"rfgodata/constants"
	"rfgodata/transactions"
)

// GetTransactionInParams : method for get transaction in params
func GetTransactionInParams(mapParams *map[string]interface{}) (transactions.ITransaction, error) {
	var returnTransaction transactions.ITransaction = nil
	var returnError error = nil
	var paramExists bool = false

	if mapParams != nil {
		returnTransaction, paramExists = (*mapParams)[constants.ParamTransaction].(transactions.ITransaction)

		// If not exist in mapparams return nil and return code error
		if !paramExists {
			returnTransaction = nil
			returnError = errors.New(utilsstring.IntToString(int(constants.ErrorTransactionNotFound)))
		}
	}

	return returnTransaction, returnError
}
