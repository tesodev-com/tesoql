package tesoql

const (
	MONGO_ENGINE      = "mongo"
	SQLITE_ENGINE     = "sqlite3"
	MYSQL_ENGINE      = "mysql"
	POSTGRES_ENGINE   = "postgres"
	SQLSERVER_ENGINE  = "sqlserver"
	ORACLE_ENGINE     = "godror"
	CLICKHOUSE_ENGINE = "clickhouse"
	ATHENA_ENGINE     = "athena"
	DYNAMODB_ENGINE   = "dynamodb"
	AVATICA_ENGINE    = "avatica"
	H2_ENGINE         = "h2"
	HIVE_ENGINE       = "hive"
	IGNITE_ENGINE     = "ignite"
	IMPALA_ENGINE     = "impala"
	COSMOSDB_ENGINE   = "cosmos"
	COUCHBASE_ENGINE  = "couchbase"
	DB2_ENGINE        = "db2"
	DATABRICKS_ENGINE = "databricks"
	DUCKDB_ENGINE     = "duckdb"
	EXASOL_ENGINE     = "exasol"
	FIREBIRD_ENGINE   = "firebirdsql"
	GENJI_ENGINE      = "genji"
	BIGQUERY_ENGINE   = "bigquery"
	SPANNER_ENGINE    = "spanner"
	ADODB_ENGINE      = "adodb"
	MAXCOMPUTE_ENGINE = "maxcompute"
	MYSQL_ENGINE_ALT1 = "mysql"
	MYSQL_ENGINE_ALT2 = "mysql"
	ODBC_ENGINE       = "odbc"
	QL_ENGINE         = "ql"
	ASE_ENGINE        = "ase"
	HANA_ENGINE       = "hdb"
	RESTSQL_ENGINE    = "restsql"
	SNOWFLAKE_ENGINE  = "snowflake"
	TRINO_ENGINE      = "trino"
	VERTICA_ENGINE    = "vertica"
	VITESS_ENGINE     = "vitess"
	YDB_ENGINE        = "ydb"
	YQL_ENGINE        = "yql"
)

// Error types
const (
	BINDING_ERR             = "BINDING_ERROR"
	TESOQL_MONGO_ERROR      = "MONGO_ERROR"
	TESOQL_SQL_ERROR        = "TESOQL_SQL_ERROR"
	TESOQL_TOGGLE_ERROR     = "TESOQL_TOGGLE_ERROR"
	TESOQL_VALIDATION_ERROR = "TESOQL_VALIDATION_ERROR"
)

// Validation Error Codes
const (
	BINDING_ERR_CODE    = 400000
	SORTABLE_ERR_CODE   = 400001
	SEARCHABLE_ERR_CODE = 400002
	PROJECTION_ERR_CODE = 400003
	CONDITION_ERR_CODE  = 400004
)

// Toggle Validation Error Codes
const (
	SORTABLE_TOGGLE_ERR_CODE                     = 400005
	SEARCHABLE_TOGGLE_ERR_CODE                   = 400006
	PROJECTION_TOGGLE_ERR_CODE                   = 400007
	CONDITION_TOGGLE_ERR_CODE                    = 400008
	PAGINATION_TOGGLE_ERR_CODE                   = 400009
	GREATERTHAN_CONDITION_TOGGLE_ERR_CODE        = 400010
	GREATEROREQUAL_CONDITION_TOGGLE_ERR_CODE     = 400011
	LOWERTHAN_CONDITION_TOGGLE_ERR_CODE          = 400012
	LOWEROREQUAL_CONDITION_TOGGLE_ERR_CODE       = 400013
	VALUESTOEXCLUDE_CONDITION_TOGGLE_ERR_CODE    = 400014
	VALUESTOEXACTMATCH_CONDITION_TOGGLE_ERR_CODE = 400015
	LOWTOHIGH_CONDITION_TOGGLE_ERR_CODE          = 400016
	HIGHTOLOW_CONDITION_TOGGLE_ERR_CODE          = 400017
)

// Repository Level Error Codes
const (
	SQL_QUERYEXEC_ERR_CODE       = 500001
	SQL_COLUMNS_ERR_CODE         = 500002
	SQL_SCAN_ERR_CODE            = 500003
	SQL_COUNT_QUERYEXEC_ERR_CODE = 500004

	MONGO_FIND_ERR_CODE   = 500005
	MONGO_CURSOR_ERR_CODE = 500006
)

//	MONGO_EMPTY_QUERY_ERR_CODE                   = 404018
//
// Sql Driver list
var sqlDriverList = map[string]string{
	SQLITE_ENGINE:     SQLITE_ENGINE,
	MYSQL_ENGINE:      MYSQL_ENGINE,
	POSTGRES_ENGINE:   POSTGRES_ENGINE,
	SQLSERVER_ENGINE:  SQLSERVER_ENGINE,
	ORACLE_ENGINE:     ORACLE_ENGINE,
	CLICKHOUSE_ENGINE: CLICKHOUSE_ENGINE,
	ATHENA_ENGINE:     ATHENA_ENGINE,
	DYNAMODB_ENGINE:   DYNAMODB_ENGINE,
	AVATICA_ENGINE:    AVATICA_ENGINE,
	H2_ENGINE:         H2_ENGINE,
	HIVE_ENGINE:       HIVE_ENGINE,
	IGNITE_ENGINE:     IGNITE_ENGINE,
	IMPALA_ENGINE:     IMPALA_ENGINE,
	COSMOSDB_ENGINE:   COSMOSDB_ENGINE,
	COUCHBASE_ENGINE:  COUCHBASE_ENGINE,
	DB2_ENGINE:        DB2_ENGINE,
	DATABRICKS_ENGINE: DATABRICKS_ENGINE,
	DUCKDB_ENGINE:     DUCKDB_ENGINE,
	EXASOL_ENGINE:     EXASOL_ENGINE,
	FIREBIRD_ENGINE:   FIREBIRD_ENGINE,
	GENJI_ENGINE:      GENJI_ENGINE,
	BIGQUERY_ENGINE:   BIGQUERY_ENGINE,
	SPANNER_ENGINE:    SPANNER_ENGINE,
	ADODB_ENGINE:      ADODB_ENGINE,
	MAXCOMPUTE_ENGINE: MAXCOMPUTE_ENGINE,
	ODBC_ENGINE:       ODBC_ENGINE,
	QL_ENGINE:         QL_ENGINE,
	ASE_ENGINE:        ASE_ENGINE,
	HANA_ENGINE:       HANA_ENGINE,
	RESTSQL_ENGINE:    RESTSQL_ENGINE,
	SNOWFLAKE_ENGINE:  SNOWFLAKE_ENGINE,
	TRINO_ENGINE:      TRINO_ENGINE,
	VERTICA_ENGINE:    VERTICA_ENGINE,
	VITESS_ENGINE:     VITESS_ENGINE,
	YDB_ENGINE:        YDB_ENGINE,
	YQL_ENGINE:        YQL_ENGINE,
}

// DefaultConfig provides a default configuration for TesoQL.
// It is set to use MongoDB as the database engine with feature toggles
// that enable all features by default. It also includes a default pagination
// limit of 50 results per page and disables the printing of SQL queries.
//
// The default configuration can be modified as needed before passing it to
// the TesoQL initialization function.
//
// Example usage:
//
//	cfg := tesoql.DefaultConfig
//	cfg.Engine = tesoql.POSTGRES_ENGINE
//	cfg.ConnectionConfig = &tesoql.ConnectionConfig{
//		DBName:           "example_db",
//		ConnectionString: "postgres://user:password@localhost:5432/example_db",
//	}
//	tesoql := cfg.NewTesoQL()
//
// Fields:
// - Engine: The database engine to use, default is MONGO_ENGINE.
// - ConnectionConfig: Connection details, default is nil (must be set by the user).
// - Toggles: Feature toggles to control search, projection, sorting, etc., all enabled by default.
// - FieldsMap: Field mappings, default is nil (can be customized).
// - Pagination: Pagination settings, default is 50 results per page.
// - PrintSqlQuery: Flag to determine if SQL queries should be printed, default
var DefaultConfig = Config{
	Engine:           MONGO_ENGINE,
	ConnectionConfig: nil,
	Toggles: &ToggleConfig{
		DisableSearch:       false,
		DisableProjection:   false,
		DisableSorting:      false,
		DisableConditioning: false,
		DisablePagination:   false,
		SortingToggles: &SortingToggles{
			DisableHighToLow: false,
			DisableLowToHigh: false,
		},
		ConditioningToggles: &ConditioningToggles{
			DisableGreaterThan:        false,
			DisableGreaterOrEqual:     false,
			DisableValuesToExactMatch: false,
			DisableLowerThan:          false,
			DisableLowerOrEqual:       false,
			DisableValuesToExclude:    false,
		},
	},
	FieldsMap: nil,
	Pagination: &PaginationConfig{
		LimitUpperBound: 50,
	},
	PrintSqlQuery: false,
}
