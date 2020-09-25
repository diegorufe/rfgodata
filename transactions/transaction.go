package transactions

// ITransaction : interface define methods for transactions
type ITransaction interface {
	// Edit : method for edit data
	Edit(data interface{}) (interface{}, error)
}
