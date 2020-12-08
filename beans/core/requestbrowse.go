package core

import "rfgodata/beans/query"

// RequestBrowse : class store config for request browser
type RequestBrowse struct {
	Limit   query.Limit    `json:"limit"`
	Fields  []query.Field  `json:"fields"`
	Joins   []query.Join   `json:"joins"`
	Filters []query.Filter `json:"filters"`
	Orders  []query.Order  `json:"orders"`
}
