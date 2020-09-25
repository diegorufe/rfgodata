package query

// Limit : for apply in querys
type Limit struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// NewLimit : method for create limit
func NewLimit() *Limit {
	var bean *Limit = new(Limit)
	bean.Start = 0
	bean.End = 0
	return bean
}
