package query

// Field : for apply in querys
type Field struct {
	Name        string `json:"name"`
	AliasTable  string `json:"aliasTable"`
	AliasField  string `json:"aliasField"`
	CustomField string `json:"customField"`
}
