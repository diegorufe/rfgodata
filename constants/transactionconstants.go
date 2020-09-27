package constants

// ParamTransaction : for get transaction in map values
const ParamTransaction string = "ParamTransaction"

// ErrorTransactionType : code erros for transaction
type ErrorTransactionType int

const (
	// ErrorTransactionNotFound : transaction not found error
	ErrorTransactionNotFound ErrorTransactionType = 0xE000001
)

// TransactionType : Indicate type of transaction
type TransactionType int

const (
	// TransactionGorm : transaction for gorm
	TransactionGorm TransactionType = 0
)
