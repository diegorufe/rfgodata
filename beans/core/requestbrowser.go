package core

import "rfgodata/beans/query"

// RequestBrowser : class store config for request browser
type RequestBrowser struct {
	Fields      []query.Field  `json:"fields"`
	Joins       []query.Join   `json:"joins"`
	Filters     []query.Filter `json:"filters"`
	Orders      []query.Order  `json:"orders"`
	Groups      []query.Order  `json:"groups"`
	First       int            `json:"first"`
	RecordsPage int            `json:"recordsPage"`
}
