package query

import "rfgodata/constants/query"

// Join : for apply in querys
type Join struct {
	Field           string         `json:"field"`
	Alias           string         `json:"alias"`
	JoinType        query.JoinType `json:"joinType"`
	CustomQueryJoin string         `json:"customQueryJoin"`
	AliasJoinField  string         `json:"aliasJoinField"`
}
