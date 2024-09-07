package tesoql

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

// MongoQuery represents a MongoDB query structure, including filter, projection, sort, limit, and offset options.
type MongoQuery struct {
	Filter     *bson.D // Filter criteria for the MongoDB query.
	Projection *bson.D // Fields to include or exclude in the result set.
	Sort       *bson.D // Sorting criteria for the query results.
	Limit      int64   // Maximum number of documents to return.
	Offset     int64   // Number of documents to skip.
}

// NewMongoQuery creates a new MongoQuery based on the provided FieldsMap and JsonMap.
// It sets up the filter, sort, projection, limit, and offset for the query.
//
// Example usage:
//
//	fm := &tesoql.FieldsMap{ /* field mappings */ }
//	jm := &tesoql.JsonMap{ /* JSON input */ }
//	query := jm.NewMongoQuery(fm)
//
// Returns:
//
// - *MongoQuery: A pointer to the initialized MongoQuery struct.
func (jm *JsonMap) NewMongoQuery(fm *FieldsMap) *MongoQuery {
	query := new(MongoQuery)
	query.Filter = getMongoFilter(fm, jm)
	query.Sort = getMongoSortCondition(fm, jm)
	query.Projection = getMongoProjection(fm, jm)
	query.Limit = jm.Pagination.Limit
	query.Offset = jm.Pagination.Offset
	return query
}

func getMongoFilter(fm *FieldsMap, jm *JsonMap) *bson.D {
	var filterArr bson.A
	var condArr bson.A

	filterArr = addMongoSearchFilter(filterArr, jm, fm)

	condArr = addMongoConditionFilter(condArr, jm, fm)

	if len(condArr) > 0 {
		combinedFilter := bson.D{{"$and", condArr}}
		filterArr = append(filterArr, combinedFilter)
	}

	if filterArr != nil {
		filter := bson.D{{"$and", filterArr}}
		return &filter
	}
	return nil
}

func getMongoProjection(fm *FieldsMap, jm *JsonMap) *bson.D {

	var projection bson.D

	for _, field := range jm.ProjectionFields {
		projection = append(projection, bson.E{Key: fm.ProjectionFields[field], Value: 1})
	}
	if projection != nil {
		return &projection
	}
	return nil
}

func getMongoSortCondition(fm *FieldsMap, jm *JsonMap) *bson.D {

	var sort bson.D

	for _, sortInput := range jm.SortConditions {
		var sortCondition int
		switch sortInput.SortCondition {
		case "ASC":
			sortCondition = 1
			break
		case "DESC":
			sortCondition = -1
			break
		default:
			sortCondition = 1
		}
		sort = append(sort, bson.E{Key: fm.SortingFields[sortInput.Field], Value: sortCondition})
	}
	if sort != nil {
		return &sort
	}
	return nil
}

func addMongoSearchFilter(filterArr bson.A, jm *JsonMap, fm *FieldsMap) bson.A {
	for key, values := range jm.Search {
		var orFilters bson.A
		for _, value := range values {
			orFilters = append(orFilters, bson.D{
				{fm.SearchFields[key], primitive.Regex{Pattern: fmt.Sprintf("%v", value), Options: "i"}},
			})
		}
		if orFilters != nil {
			filterArr = append(filterArr, bson.D{{"$or", orFilters}})
		}
	}
	return filterArr
}

func addMongoConditionFilter(condArr bson.A, jm *JsonMap, fm *FieldsMap) bson.A {
	var condition bson.D
	for k, v := range jm.Conditions {
		if v.ValuesToExactMatch != nil {
			if isDateTimeFieldKey(k, fm.DateTimeFieldKeys) {
				for _, value := range v.ValuesToExactMatch {
					value = parseDateTimeOrReturn(value)
				}
			}
			condition = bson.D{{fm.ConditionFields[k], bson.D{{"$in", v.ValuesToExactMatch}}}}

			condArr = append(condArr, condition)
		}
		if v.GreaterOrEqual != nil {

			v.GreaterOrEqual = treatDateTime(k, fm.DateTimeFieldKeys, v.GreaterOrEqual)

			condition = bson.D{{fm.ConditionFields[k], bson.D{{"$gte", v.GreaterOrEqual}}}}

			condArr = append(condArr, condition)
		}
		if v.GreaterThan != nil {

			v.GreaterThan = treatDateTime(k, fm.DateTimeFieldKeys, v.GreaterThan)

			condition = bson.D{{fm.ConditionFields[k], bson.D{{"$gt", v.GreaterThan}}}}

			condArr = append(condArr, condition)
		}
		if v.LowerOrEqual != nil {

			v.LowerOrEqual = treatDateTime(k, fm.DateTimeFieldKeys, v.LowerOrEqual)

			condition = bson.D{{fm.ConditionFields[k], bson.D{{"$lte", v.LowerOrEqual}}}}

			condArr = append(condArr, condition)
		}
		if v.LowerThan != nil {

			v.LowerThan = treatDateTime(k, fm.DateTimeFieldKeys, v.LowerThan)

			condition = bson.D{{fm.ConditionFields[k], bson.D{{"$lt", v.LowerThan}}}}

			condArr = append(condArr, condition)
		}
		if v.ValuesToExclude != nil {
			if isDateTimeFieldKey(k, fm.DateTimeFieldKeys) {
				for _, item := range v.ValuesToExclude {
					item = parseDateTimeOrReturn(item)
				}
			}

			condition = bson.D{{fm.ConditionFields[k], bson.D{{"$nin", v.ValuesToExclude}}}}

			condArr = append(condArr, condition)
		}
	}
	return condArr
}

// SqlQuery represents an SQL query structure, including select, where, order by, limit, offset clauses, and arguments.
type SqlQuery struct {
	Select  string        // Fields to select in the SQL query.
	Where   string        // Filter conditions for the SQL query.
	OrderBy string        // Sorting criteria for the SQL query.
	Limit   string        // Maximum number of rows to return.
	Offset  string        // Number of rows to skip.
	Args    []interface{} // Arguments for the query's placeholders.
}

// NewSqlQuery creates a new SqlQuery based on the provided FieldsMap and JsonMap.
// It sets up the select fields, where conditions, sorting, limit, and offset for the query.
//
// Example usage:
//
//	fm := &tesoql.FieldsMap{ /* field mappings */ }
//	jm := &tesoql.JsonMap{ /* JSON input */ }
//	query := jm.NewSqlQuery(fm)
//
// Returns:
// - *SqlQuery: A pointer to the initialized SqlQuery struct.
func (jm *JsonMap) NewSqlQuery(fm *FieldsMap) *SqlQuery {
	query := new(SqlQuery)
	query.Select = getSqlProjection(fm, jm)
	query.Where, query.Args = getSqlFilter(fm, jm)
	query.OrderBy = getSqlSortCondition(fm, jm)
	query.Limit = fmt.Sprintf("LIMIT %d", jm.Pagination.Limit)
	query.Offset = fmt.Sprintf("OFFSET %d", jm.Pagination.Offset)
	return query
}

func getSqlProjection(fm *FieldsMap, jm *JsonMap) string {
	if len(jm.ProjectionFields) > 0 {
		var fields []string
		for _, field := range jm.ProjectionFields {
			if value, exists := fm.ProjectionFields[field]; exists {
				fields = append(fields, value)
			}
		}
		return strings.Join(fields, ", ")
	}
	return "*"
}

func getSqlFilter(fm *FieldsMap, jm *JsonMap) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	conditions, args = addSqlSearchFilter(fm, jm, conditions, args)

	conditions, args = addSqlConditionFilters(fm, jm, conditions, args)

	return strings.Join(conditions, " AND "), args
}

func getSqlSortCondition(fm *FieldsMap, jm *JsonMap) string {
	var orderBy []string
	for _, sortInput := range jm.SortConditions {
		orderBy = append(orderBy, fmt.Sprintf("%s %s", fm.SortingFields[sortInput.Field], sortInput.SortCondition))
	}
	if len(orderBy) > 0 {
		return fmt.Sprintf(" ORDER BY %s", strings.Join(orderBy, ", "))
	}
	return ""
}

func addSqlSearchFilter(fm *FieldsMap, jm *JsonMap, conditions []string, args []interface{}) ([]string, []interface{}) {
	for key, values := range jm.Search {
		var orConditions []string
		for _, value := range values {
			orConditions = append(orConditions, fmt.Sprintf("%s LIKE ?", fm.SearchFields[key]))
			args = append(args, fmt.Sprintf("%%%v%%", value))
		}
		if orConditions != nil {
			conditions = append(conditions, fmt.Sprintf("(%s)", strings.Join(orConditions, " OR ")))
		}
	}
	return conditions, args
}

func addSqlConditionFilters(fm *FieldsMap, jm *JsonMap, conditions []string, args []interface{}) ([]string, []interface{}) {
	for key, condition := range jm.Conditions {
		if condition.ValuesToExactMatch != nil && len(condition.ValuesToExactMatch) > 0 {
			placeholders := make([]string, len(condition.ValuesToExactMatch))
			for i, value := range condition.ValuesToExactMatch {
				placeholders[i] = "?"
				args = append(args, value)
			}
			conditions = append(conditions, fmt.Sprintf("%s IN (%s)", fm.ConditionFields[key], strings.Join(placeholders, ", ")))
		}
		if condition.GreaterOrEqual != nil {
			conditions = append(conditions, fmt.Sprintf("%s >= ?", fm.ConditionFields[key]))
			args = append(args, condition.GreaterOrEqual)
		}
		if condition.GreaterThan != nil {
			conditions = append(conditions, fmt.Sprintf("%s > ?", fm.ConditionFields[key]))
			args = append(args, condition.GreaterThan)
		}
		if condition.LowerOrEqual != nil {
			conditions = append(conditions, fmt.Sprintf("%s <= ?", fm.ConditionFields[key]))
			args = append(args, condition.LowerOrEqual)
		}
		if condition.LowerThan != nil {
			conditions = append(conditions, fmt.Sprintf("%s < ?", fm.ConditionFields[key]))
			args = append(args, condition.LowerThan)
		}
		if condition.ValuesToExclude != nil && len(condition.ValuesToExclude) > 0 {
			placeholders := make([]string, len(condition.ValuesToExclude))
			for i, value := range condition.ValuesToExclude {
				placeholders[i] = "?"
				args = append(args, value)
			}
			conditions = append(conditions, fmt.Sprintf("%s NOT IN (%s)", fm.ConditionFields[key], strings.Join(placeholders, ", ")))
		}
	}
	return conditions, args
}

// GetSqlQuery generates a full SQL query string, including the select, where, order by, limit, and offset clauses.
// It also returns the where clause and the arguments for the query.
//
// The reason for returning the where clause and arguments separately is to prevent SQL injection attacks.
// By using placeholders in the query and passing the arguments separately, it is ensured that user input
// is properly sanitized and doesn't lead to security vulnerabilities.
//
// Example usage:
//
//	fm := &tesoql.FieldsMap{ /* field mappings */ }
//	jm := &tesoql.JsonMap{ /* JSON input */ }
//	sqlQuery, whereClause, args := jm.GetSqlQuery(fm, "table_name", true)
//
// Returns:
//
// - string: The full SQL query string.
//
// - string: The where clause of the SQL query.
//
// - []interface{}: The arguments for the query's placeholders.
func (jm *JsonMap) GetSqlQuery(fieldsMap *FieldsMap, tableName string, printQuery bool) (string, string, []interface{}) {

	query := jm.NewSqlQuery(fieldsMap)

	sqlQuery := fmt.Sprintf("SELECT %s FROM %s WHERE 1=1", query.Select, tableName)

	var whereClause string

	if query.Where != "" {
		whereClause = fmt.Sprintf(" AND %s", query.Where)
	}
	whereClause += fmt.Sprintf("%s %s %s", query.OrderBy, query.Limit, query.Offset)

	if printQuery {
		fmt.Printf("Query: %s\nWith Arguments: %v\n", sqlQuery+whereClause, query.Args)
	}

	return sqlQuery + whereClause, whereClause, query.Args
}
