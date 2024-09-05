package tesoql

// Config holds the configuration settings for initializing a TesoQL instance.
// It includes database engine settings, connection configurations, feature toggles,
// field mappings, pagination settings, and a flag to print SQL queries.
type Config struct {
	Engine           string            // The database engine to use (e.g., "mongo", "mysql").
	ConnectionConfig *ConnectionConfig // The configuration for database connection details.
	Toggles          *ToggleConfig     // Feature toggles to enable or disable specific behaviors.
	FieldsMap        *FieldsMap        // Mappings for different fields like search, sorting, etc.
	Pagination       *PaginationConfig // Configuration for pagination settings.
	PrintSqlQuery    bool              // Flag to determine if SQL queries should be printed.
}

// FieldsMap defines the mappings for various field types.
// These include datetime fields, search fields, sorting fields,
// projection fields, and condition fields.
type FieldsMap struct {
	DateTimeFieldKeys map[string]string // Mappings for datetime fields.
	SearchFields      map[string]string // Mappings for search fields.
	SortingFields     map[string]string // Mappings for sorting fields.
	ProjectionFields  map[string]string // Mappings for projection fields.
	ConditionFields   map[string]string // Mappings for condition fields.
}

// ConnectionConfig holds the database connection details.
// It includes the database name, table name, client, and connection string.
type ConnectionConfig struct {
	DBName           string      // The name of the database.
	TableName        string      // The name of the table.
	Client           interface{} // The database client instance.
	ConnectionString string      // The connection string to the database.
}

// ToggleConfig defines the feature toggles that control
// the enabling or disabling of specific TesoQL functionalities.
type ToggleConfig struct {
	DisableSearch       bool                 // Toggle to disable search functionality.
	DisableProjection   bool                 // Toggle to disable projection functionality.
	DisableSorting      bool                 // Toggle to disable sorting functionality.
	DisableConditioning bool                 // Toggle to disable conditioning functionality.
	DisablePagination   bool                 // Toggle to disable pagination.
	DisableTotalCount   bool                 // Toggle to disable total count calculation.
	SortingToggles      *SortingToggles      // Nested toggles for sorting behavior.
	ConditioningToggles *ConditioningToggles // Nested toggles for conditioning behavior.
}

// SortingToggles holds the toggles related to sorting behavior.
// It allows enabling or disabling specific sorting options like
// high-to-low and low-to-high sorting.
type SortingToggles struct {
	DisableHighToLow bool // Toggle to disable high-to-low sorting.
	DisableLowToHigh bool // Toggle to disable low-to-high sorting.
}

// ConditioningToggles holds the toggles related to conditioning behavior.
// It allows enabling or disabling specific conditions like greater than,
// lower than, exact match, and exclusions.
type ConditioningToggles struct {
	DisableGreaterThan        bool // Toggle to disable "greater than" condition.
	DisableGreaterOrEqual     bool // Toggle to disable "greater or equal" condition.
	DisableValuesToExactMatch bool // Toggle to disable exact match condition.
	DisableLowerThan          bool // Toggle to disable "lower than" condition.
	DisableLowerOrEqual       bool // Toggle to disable "lower or equal" condition.
	DisableValuesToExclude    bool // Toggle to disable exclusion condition.
}

// PaginationConfig defines the settings related to pagination.
// It includes the upper bound limit for pagination results.
type PaginationConfig struct {
	LimitUpperBound int64 // The upper bound for the number of results per page.
}
