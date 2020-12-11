package utils

import (
	"database/sql"
	"reflect"
	"rfgodata/beans/query"
	querycst "rfgodata/constants/query"
	"time"
)

// GetFilterSelectPk : method for get field select with pk from table.
func GetFilterSelectPk(pk interface{}) query.Filter {
	return query.Filter{FilterType: querycst.LiteralEqual, FilterOperationType: querycst.And, Field: "id", Value: pk}
}

// GetFieldSelectAll : method for get field select all values. DONT include joins asociations
func GetFieldSelectAll() query.Field {
	return query.Field{Name: "*"}
}

// AddCreatedAt Function for added created at if exist field
func AddCreatedAt(data interface{}) {
	if data != nil {
		instanceData := reflect.ValueOf(data).Elem()

		// find created at
		valueField := instanceData.FieldByName("CreatedAt")

		if valueField.IsValid() && valueField.CanSet() {
			timeToSet := sql.NullTime{Time: time.Now(), Valid: true}
			vallueToSet := reflect.ValueOf(timeToSet)
			valueField.Set(vallueToSet)
		}
	}
}

// AddUpdatedAt Function for added updated at at if exist field
func AddUpdatedAt(data interface{}) {
	if data != nil {
		instanceData := reflect.ValueOf(data).Elem()

		// find created at
		valueField := instanceData.FieldByName("UpdatedAt")

		if valueField.IsValid() && valueField.CanSet() {
			timeToSet := sql.NullTime{Time: time.Now(), Valid: true}
			vallueToSet := reflect.ValueOf(timeToSet)
			valueField.Set(vallueToSet)
		}
	}
}
