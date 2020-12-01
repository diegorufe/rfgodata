package query

// JoinType indicate type for join
type JoinType string

const (
	InnerJoin JoinType = "INNER JOIN"

	InnerJoinFetch JoinType = "INNER JOIN FETCH"

	LeftJoin JoinType = "LEFT JOIN"

	LeftJoinFetch JoinType = "LEFT JOIN FETCH"

	RightJoin JoinType = "RIGHT JOIN"

	RightJoinFetch JoinType = "RIGHT JOIN FETCH"
)
