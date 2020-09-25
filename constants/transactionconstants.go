package constants

// ParamTransaction : for get transaction in map values
const ParamTransaction string = "ParamTransaction"

// ErrorTransactionType : code erros for transaction
type ErrorTransactionType int

const (
	// ErrorTransactionNotFound : transaction not found error
	ErrorTransactionNotFound ErrorTransactionType = 0xE000001
)
