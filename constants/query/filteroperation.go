package query

// FilterOperation indicate type operation for filter
type FilterOperationType string

const (
	And FilterOperationType = "AND"

	Or FilterOperationType = "OR"
)
