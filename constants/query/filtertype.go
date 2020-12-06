package query

// FilterType indicate type for filter
type FilterType string

const (
	Equal FilterType = "="

	LiteralEqual FilterType = "EQUAL"

	Ge FilterType = ">"

	Gt FilterType = ">="

	In FilterType = "IN"

	NotEqual FilterType = "!="

	NotIn FilterType = "NOT_IN"

	Le FilterType = "<="

	Like FilterType = "LIKE"

	LikeStart FilterType = "LIKE_START"

	LikeEnd FilterType = "LIKE_END"

	Lt FilterType = "<"
)
