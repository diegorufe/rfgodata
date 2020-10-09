package gorm

import (
	"rfgocore/utils/utilsstring"
	"rfgodata/beans/query"
	queryconstants "rfgodata/constants/query"

	"gorm.io/gorm"
)

// DefaultAliasQuery indicates default alias for use in querys if not pass alias
const DefaultAliasQuery string = "defaultAliasQuery"

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
	var dbReturn *gorm.DB
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
				case queryconstants.RigthJoinFetch:
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

				case queryconstants.RigthJoin:
				case queryconstants.RigthJoinFetch:
					queryBuilder = queryBuilder + " RIGTH JOIN " + join.Field
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
	var queryBuilder string = ""

	if fields != nil && len(fields) > 0 {

	} else {
		queryBuilder = queryBuilder + " " + DefaultAliasQuery + ".* "
	}

	dbReturn = dbReturn.Select(queryBuilder)

	return dbReturn
}
