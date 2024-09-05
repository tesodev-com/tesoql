package tesoql

// JsonMap represents the structure for defining query parameters.
// It includes search filters, projection fields, sorting conditions,
// complex conditions, pagination, and options to control the response behavior.
type JsonMap struct {
	Search               map[string][]interface{}      `json:"search"`               // Search criteria mapped by field names.
	ProjectionFields     []string                      `json:"projectionFields"`     // Fields to include in the query result.
	SortConditions       []SortInput                   `json:"sortConditions"`       // Sorting conditions for the query results.
	Conditions           map[string]ConditionOperators `json:"conditions"`           // Complex conditions for filtering the data.
	Pagination           Pagination                    `json:"pagination"`           // Pagination settings for limiting and offsetting the results.
	TotalCount           bool                          `json:"totalCount"`           // Flag to determine whether to include the total count of records.
	SuppressDataResponse bool                          `json:"suppressDataResponse"` // Flag to suppress the data response (useful for count-only queries).
}

// ConditionOperators defines the various operators that can be applied
// to a particular field in the query, such as greater than, less than, and exact match.
type ConditionOperators struct {
	GreaterThan        interface{}   `json:"greaterThan"`        // Greater than condition.
	GreaterOrEqual     interface{}   `json:"greaterOrEqual"`     // Greater than or equal condition.
	ValuesToExactMatch []interface{} `json:"valuesToExactMatch"` // Exact match condition (array of values).
	LowerThan          interface{}   `json:"lowerThan"`          // Less than condition.
	LowerOrEqual       interface{}   `json:"lowerOrEqual"`       // Less than or equal condition.
	ValuesToExclude    []interface{} `json:"valuesToExclude"`    // Exclusion condition (array of values).
}

// Pagination defines the structure for paginating query results.
// It includes settings for limiting the number of results and skipping a certain number of records.
type Pagination struct {
	Limit  int64 `json:"limit"`  // Maximum number of results to return.
	Offset int64 `json:"offset"` // Number of results to skip before starting to return results.
}

// SortInput defines the structure for specifying sorting behavior in a query.
// It includes the field to sort by and the sorting condition (e.g., ascending or descending).
type SortInput struct {
	Field         string `json:"field"`         // The field by which to sort.
	SortCondition string `json:"sortCondition"` // The sorting condition (e.g., "ASC" for ascending, "DESC" for descending).
}

// ErrorResponseDTO is used to represent errors that occur during query processing.
// It includes the type of error, a descriptive message, and an error code.
type ErrorResponseDTO struct {
	ErrorType string // The type of error (e.g., validation error, repository error).
	ErrorMsg  string // A detailed error message.
	ErrorCode int    // A numeric code representing the specific error.
}
