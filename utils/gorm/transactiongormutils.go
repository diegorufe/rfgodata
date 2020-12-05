package gorm

import (
	"container/list"
	"reflect"
	"rfgocore/utils/utilsstring"
	"rfgodata/beans/query"
	queryconstants "rfgodata/constants/query"
	"strings"

	"gorm.io/gorm"
)

// DefaultAliasQuery indicates default alias for use in querys if not pass alias
const DefaultAliasQuery string = "defaultAliasQuery"

// FieldSeparator for find sub entities
const FieldSeparator string = "_FIELD_SEPARATOR_"

func applyWhereQueryBuilder(filters []query.Filter, firstLevel bool, valuesQuery *[]interface{}) (string, []interface{}) {
	var queryBuilder string = ""

	if filters != nil && len(filters) > 0 {
		for index, filter := range filters {
			queryBuilder = queryBuilder + " "
			if !(firstLevel && index == 0) {
				// And / or operation type
				switch filter.FilterOperationType {

				case queryconstants.And:
					queryBuilder = queryBuilder + " AND "
					break

				case queryconstants.Or:
					queryBuilder = queryBuilder + " OR "
					break
				}

			}

			// open brackets
			if filter.OpenBrackets > 0 {
				for i := 0; i < filter.OpenBrackets; i++ {
					queryBuilder = queryBuilder + " ( "
				}
			}

			// Alias
			if utilsstring.IsNotEmpty(filter.Alias) {
				queryBuilder = queryBuilder + " " + filter.Alias + "."
			} else {
				queryBuilder = queryBuilder + " " + DefaultAliasQuery + "."
			}

			// Field
			if utilsstring.IsNotEmpty(filter.Field) {
				queryBuilder = queryBuilder + filter.Field
			}

			// Filter type operation
			switch filter.FilterType {

			case queryconstants.Equal:
				queryBuilder = queryBuilder + " = ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
				break

			case queryconstants.NotEqual:
				queryBuilder = queryBuilder + " != ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
				break

			case queryconstants.Ge:
				queryBuilder = queryBuilder + " > ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
				break

			case queryconstants.Gt:
				queryBuilder = queryBuilder + " >= ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
				break

			case queryconstants.In:
				queryBuilder = queryBuilder + " IN ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
				break

			case queryconstants.NotIn:
				queryBuilder = queryBuilder + " NOT IN ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
				break

			case queryconstants.Like:
				queryBuilder = queryBuilder + " LIKE ? "
				*valuesQuery = append(*valuesQuery, "%"+filter.Value.(string)+"%")
				break

			case queryconstants.LikeStart:
				queryBuilder = queryBuilder + " LIKE ? "
				*valuesQuery = append(*valuesQuery, "%"+filter.Value.(string))
				break

			case queryconstants.LikeEnd:
				queryBuilder = queryBuilder + " LIKE ? "
				*valuesQuery = append(*valuesQuery, filter.Value.(string)+"%")
				break

			case queryconstants.Le:
				queryBuilder = queryBuilder + " <= ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
				break

			case queryconstants.Lt:
				queryBuilder = queryBuilder + " < ? "
				*valuesQuery = append(*valuesQuery, filter.Value)
				break

			}

			// close brackets
			if filter.CloseBrackets > 0 {
				for i := 0; i < filter.CloseBrackets; i++ {
					queryBuilder = queryBuilder + " ) "
				}
			}
		}

	}

	return queryBuilder, *valuesQuery
}

// ApplyWhere : for apply where query
func ApplyWhere(db *gorm.DB, filters []query.Filter) *gorm.DB {
	var dbReturn *gorm.DB = db
	var queryBuilder string = ""
	var valuesQuery []interface{}

	queryBuilder, valuesQuery = applyWhereQueryBuilder(filters, true, &valuesQuery)

	// If have query call where
	if utilsstring.IsNotEmpty(queryBuilder) {
		dbReturn = db.Where(queryBuilder, valuesQuery)
	}

	return dbReturn
}

// ApplyLimit for query
func ApplyLimit(db *gorm.DB, limit query.Limit) *gorm.DB {
	return db.Offset(limit.Start).Limit(limit.End)
}

// ApplyJoins query
func ApplyJoins(db *gorm.DB, joins []query.Join) *gorm.DB {
	var dbReturn *gorm.DB = db
	var queryBuilder string = ""

	if joins != nil && len(joins) > 0 {
		for _, join := range joins {

			if utilsstring.IsNotEmpty(join.JoinFetchPreload) {
				switch join.JoinType {
				case queryconstants.InnerJoinFetch:
					dbReturn = dbReturn.Preload(join.JoinFetchPreload)
					break
				case queryconstants.LeftJoinFetch:
					dbReturn = dbReturn.Preload(join.JoinFetchPreload)
					break
				case queryconstants.RightJoinFetch:
					dbReturn = dbReturn.Preload(join.JoinFetchPreload)
					break
				}
			}

			queryBuilder = ""

			if utilsstring.IsNotEmpty(join.CustomQueryJoin) {
				queryBuilder = queryBuilder + " " + join.CustomQueryJoin
			} else {
				switch join.JoinType {

				case queryconstants.InnerJoin:
				case queryconstants.InnerJoinFetch:
					queryBuilder = queryBuilder + " JOIN " + join.Field
					break

				case queryconstants.LeftJoin:
				case queryconstants.LeftJoinFetch:
					queryBuilder = queryBuilder + " LEFT JOIN " + join.Field
					break

				case queryconstants.RightJoin:
				case queryconstants.RightJoinFetch:
					queryBuilder = queryBuilder + " RIGHT JOIN " + join.Field
					break

				}
			}

			if utilsstring.IsNotEmpty(join.Alias) {
				queryBuilder = queryBuilder + " as " + join.Alias
			} else {
				queryBuilder = queryBuilder + " as " + join.Field
			}

			if utilsstring.IsNotEmpty(join.JoinCondiction) {
				queryBuilder = queryBuilder + "  " + join.JoinCondiction
			}

			// Apply joins condiction
			dbReturn = dbReturn.Joins(queryBuilder)

		}
	}
	return dbReturn
}

// ApplySelect for query
func ApplySelect(db *gorm.DB, fields []query.Field) *gorm.DB {
	var dbReturn *gorm.DB = db
	//var dbReturn *gorm.DB = db
	var queryBuilder string = ""

	if fields != nil && len(fields) > 0 {
		for index, field := range fields {
			if index > 0 {
				queryBuilder = queryBuilder + " , "
			} else {
				queryBuilder = queryBuilder + " "
			}
			if utilsstring.IsNotEmpty(field.CustomField) {
				queryBuilder = queryBuilder + field.CustomField
			} else if utilsstring.IsNotEmpty(field.AliasTable) {
				queryBuilder = queryBuilder + field.AliasTable + "."
			} else {
				queryBuilder = queryBuilder + DefaultAliasQuery + "."
			}

			if utilsstring.IsEmpty(field.CustomField) {
				queryBuilder = queryBuilder + field.Name + " "
			}

			if utilsstring.IsNotEmpty(field.AliasField) {
				queryBuilder = queryBuilder + field.AliasField + " "
			}
		}
	} else {
		queryBuilder = queryBuilder + " " + DefaultAliasQuery + ".* "
	}

	dbReturn = dbReturn.Select(queryBuilder)

	return dbReturn
}

// ApplyOrders method for apply orders
func ApplyOrders(db *gorm.DB, orders []query.Order) *gorm.DB {
	var dbReturn *gorm.DB = db
	var queryBuilder string = ""

	if orders != nil && len(orders) > 0 {
		for _, order := range orders {

			if utilsstring.IsNotEmpty(order.Alias) {
				queryBuilder = queryBuilder + "  " + order.Alias + "."
			} else {
				queryBuilder = queryBuilder + "  " + DefaultAliasQuery + "."
			}

			queryBuilder = queryBuilder + order.Field

			switch order.OrderType {

			case queryconstants.Asc:
				queryBuilder = queryBuilder + " ASC "
				break

			case queryconstants.Desc:
				queryBuilder = queryBuilder + " DESC "
				break
			}
		}

		dbReturn = dbReturn.Order(queryBuilder)
	}

	return dbReturn
}

// RawData : method for raw data into array for model type
func RawData(db *gorm.DB, modelType reflect.Type) ([]interface{}, error) {
	var arrayData []interface{}

	rows, err := db.Rows()

	if err == nil {
		colsRows, err := rows.Columns()

		// If has error return this
		if err != nil {
			return nil, err
		}

		var containerList list.List
		var instanceModel interface{}
		var fieldColum reflect.Value
		var instaceModelColumnReflectColumn reflect.Value

		for rows.Next() {

			// Instace ponter of model
			resultInstanceModel := reflect.New(modelType)
			instanceModel = resultInstanceModel.Interface()

			// Create a slice of interface{}'s to represent each column,
			// and a second slice to contain pointers to each item in the columns slice.
			columnPointers := make([]interface{}, len(colsRows))
			instaceModelColumnReflect := reflect.ValueOf(instanceModel).Elem()

			for index, columKey := range colsRows {

				instaceModelColumnReflectColumn = instaceModelColumnReflect

				titleKeySplit := strings.Split(columKey, FieldSeparator)
				lenTitleSPlit := len(titleKeySplit)

				for indexTitleSplit, titleKey := range titleKeySplit {
					fieldColum = instaceModelColumnReflectColumn.FieldByName(strings.Title(titleKey))

					if indexTitleSplit >= 0 && indexTitleSplit < (lenTitleSPlit-1) && lenTitleSPlit > 0 {
						instaceModelColumnReflectColumn = instaceModelColumnReflectColumn.FieldByName(strings.Title(titleKey)).Addr().Elem()
					}

				}

				interfacePointer := fieldColum.Addr().Interface()
				columnPointers[index] = interfacePointer
			}

			// Scan the result into the column pointers...
			// TODO time error solved by NullTime
			if err = rows.Scan(columnPointers...); err != nil {
				break
			}

			containerList.PushBack(instanceModel)
		}

		// If list contain data create array return
		if containerList.Len() > 0 {
			arrayData = make([]interface{}, containerList.Len())
			var counter uint64 = 0
			for element := containerList.Front(); element != nil; element = element.Next() {
				// do something with element.Value
				arrayData[counter] = element.Value
				counter = counter + 1
			}
		}
	}

	return arrayData, err
}
