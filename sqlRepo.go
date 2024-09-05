package tesoql

import (
	"database/sql"
	"fmt"
	"sync"
)

type sqlRepository struct {
	sql           *sql.DB
	tableName     string
	fieldsMap     *FieldsMap
	printSqlQuery bool
}

func newSqlRepository(cfg *Config) *sqlRepository {
	var db *sql.DB
	client, ok := cfg.ConnectionConfig.Client.(*sql.DB)
	if ok && client != nil {
		db = client
	} else {
		var err error
		db, err = sql.Open(cfg.Engine, cfg.ConnectionConfig.ConnectionString)
		if err != nil {
			panic(fmt.Sprintf("Error opening database connection: %v", err))
		}

		err = db.Ping()
		if err != nil {
			panic(fmt.Sprintf("Error pinging database: %v", err))
		}
	}

	return &sqlRepository{
		sql:           db,
		tableName:     cfg.ConnectionConfig.TableName,
		fieldsMap:     cfg.FieldsMap,
		printSqlQuery: cfg.PrintSqlQuery,
	}
}

func (r *sqlRepository) repository(jsonMap *JsonMap) ([]map[string]interface{}, int, int, *ErrorResponseDTO) {

	query, whereClause, queryArgs := jsonMap.GetSqlQuery(r.fieldsMap, r.tableName, r.printSqlQuery)
	var results []map[string]interface{}

	if !jsonMap.SuppressDataResponse {
		rows, err := r.sql.Query(query, queryArgs...)
		if err != nil {
			return nil, 0, 0, newResponse(TESOQL_SQL_ERROR, err.Error(), SQL_QUERYEXEC_ERR_CODE)
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			return nil, 0, 0, newResponse(TESOQL_SQL_ERROR, err.Error(), SQL_COLUMNS_ERR_CODE)
		}

		for rows.Next() {
			columnPointers := make([]interface{}, len(columns))
			row := make(map[string]interface{})
			for i, col := range columns {
				var colValue interface{}
				columnPointers[i] = &colValue
				row[col] = &colValue
			}
			if err := rows.Scan(columnPointers...); err != nil {
				return nil, 0, 0, newResponse(TESOQL_SQL_ERROR, err.Error(), SQL_SCAN_ERR_CODE)
			}

			for i, col := range columns {
				row[col] = *(columnPointers[i].(*interface{}))
			}

			results = append(results, row)
		}
	}

	size := len(results)

	var tesoQlErr *ErrorResponseDTO
	var totalCount int
	var wg sync.WaitGroup
	if jsonMap.TotalCount {
		wg.Add(1)

		go func() {
			defer wg.Done()
			totalCount, tesoQlErr = r.countTotal(whereClause, queryArgs)
		}()
	}
	wg.Wait()
	return results, totalCount, size, tesoQlErr
}

func (r *sqlRepository) countTotal(whereClause string, queryArgs []interface{}) (int, *ErrorResponseDTO) {
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE 1=1", r.tableName) + whereClause

	row := r.sql.QueryRow(countQuery, queryArgs...)

	var count int
	if err := row.Scan(&count); err != nil {
		return 0, newResponse(TESOQL_SQL_ERROR, err.Error(), SQL_COUNT_QUERYEXEC_ERR_CODE)
	}

	return count, nil
}
