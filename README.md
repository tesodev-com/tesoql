# tesoql
tesoql is a Go package designed to simplify the process of building SQL and MongoDB queries. It provides a service and repository layer that abstracts the complexities of database interactions, allowing you to connect directly to your database by selecting the appropriate engine (e.g., MongoDB, SQLite3, MySQL) or use its exposed public methods to build queries. This package is ideal for developers who need a flexible and powerful tool for managing database queries across multiple database engines.

------------


### Features
- **Multi-Engine Support:** Seamlessly switch between SQL and MongoDB based on your application's needs.
- **Service Layer:** Handles data processing.
- **Repository Layer:** Defines an interface with methods tailored for SQL-based databases and MongoDB.
- **Flexible Query Building:** Use the JsonMap struct to build complex queries effortlessly.
- **Validation:** Built-in validation mechanisms ensure the correctness of queries.
- **Feature Toggles:** Enable or disable specific features at runtime.
- **SQL Injection Prevention:** Defensive approach on query building to prevent SQL injection attacks.

------------
### Installation
To install *tesoql*, use the following `go get` command:
`go get github.com/tesodev-com/tesoql`

Then, import the package in your Go project:
`import "github.com/tesodev/tesoql"`

------------

### Configuration in TesoQL
The Config struct in *tesoql* is essential for setting up and customizing the behavior of your TesoQL instance. It encompasses various configuration settings, including database engine selection, connection details, feature toggles, field mappings, pagination, and debugging options.

#### 1. Config Struct
The *Config* struct is the primary configuration structure that you'll use to initialize a TesoQL instance. It includes several key components:

```go
type Config struct {
   Engine           string            
   ConnectionConfig *ConnectionConfig 
   Toggles          *ToggleConfig     
   FieldsMap        *FieldsMap        
   Pagination       *PaginationConfig 
   PrintSqlQuery    bool              
}
```

##### Fields:
- **Engine:** Specifies the database engine that TesoQL will use, such as MongoDB, MySQL, or SQLite.
- **ConnectionConfig:** Holds the details necessary to establish a connection to the database, including the database name, table name, client, and connection string.
- **Toggles:** Provides a set of feature toggles that control whether certain functionalities (like search or sorting) are enabled or disabled.
- **FieldsMap:** Maps field names used in the code to the actual field names in the database, ensuring that queries are correctly formed.
- **Pagination:** Defines the settings for pagination, including an upper bound on the number of results per page.
- **PrintSqlQuery:** A debugging flag that, when set to true, prints the generated SQL queries to the console.

#### 2. FieldsMap Struct
The *FieldsMap* struct defines how various types of fields (such as search, sorting, and projection fields) are mapped to their corresponding database fields. This mapping is crucial for ensuring that queries are correctly formed according to the database schema. (see detailed explanation on the ‘*FieldsMap*’ section)

```go
type FieldsMap struct {
   DateTimeFieldKeys map[string]string 
   SearchFields      map[string]string 
   SortingFields     map[string]string 
   ProjectionFields  map[string]string 
   ConditionFields   map[string]string 
}
```
#### 3. ConnectionConfig Struct

The *ConnectionConfig* struct holds the necessary information to connect to a database. This includes the name of the database, the name of the table, the database client, and the connection string.

```go
type ConnectionConfig struct {
   DBName           string      
   TableName        string      
   Client           interface{} 
   ConnectionString string      
}
```

#### 4. ToggleConfig Struct

The *ToggleConfig* struct allows you to enable or disable specific functionalities in TesoQL. It provides a fine-grained level of control over features such as search, projection, sorting, and more.

```go
type ToggleConfig struct {
   DisableSearch       bool                 
   DisableProjection   bool                 
   DisableSorting      bool                 
   DisableConditioning bool                 
   DisablePagination   bool                 
   DisableTotalCount   bool                 
   SortingToggles      *SortingToggles      
   ConditioningToggles *ConditioningToggles 
}
```

#### 5. SortingToggles Struct

The *SortingToggles* struct within *ToggleConfig* provides additional control over sorting behaviors. It allows you to disable specific sorting options such as high-to-low or low-to-high sorting.

```go
type SortingToggles struct {
   DisableHighToLow bool 
   DisableLowToHigh bool 
}
```

#### 6. ConditioningToggles Struct

The *ConditioningToggles* struct allows you to control the behavior of conditioning options such as greater than, less than, exact match, and exclusion.

```go
type ConditioningToggles struct {
   DisableGreaterThan        bool 
   DisableGreaterOrEqual     bool 
   DisableValuesToExactMatch bool 
   DisableLowerThan          bool 
   DisableLowerOrEqual       bool 
   DisableValuesToExclude    bool 
}
```

#### 7. PaginationConfig Struct
The PaginationConfig struct defines the settings related to pagination, including the maximum number of results that can be returned per page.

```go
    type PaginationConfig struct {
       LimitUpperBound int64 
    }
```

------------

### Usage

#### Basic Setup
Opening a connection with a Mongodb and an SQL based database differs.To get started with *tesoql*, you need to configure it with your desired database engine and connection details. The engine list of the supported databases by *tesoql*;

|  tesoql DB Engine Variable | SQL Driver Name  |
| ------------ | ------------ |
|  MONGO_ENGINE |  mongo |
| SQLITE_ENGINE  | sqlite3  |
| MYSQL_ENGINE | mysql |
| POSTGRES_ENGINE | postgres |
| SQLSERVER_ENGINE | sqlserver |
| ORACLE_ENGINE | godror |
| CLICKHOUSE_ENGINE | clickhouse |
| ATHENA_ENGINE | athena |
| DYNAMODB_ENGINE | dynamodb |
| AVATICA_ENGINE | avatica |
| H2_ENGINE | h2 |
| HIVE_ENGINE | hive |
| IGNITE_ENGINE | ignite |
| IMPALA_ENGINE | impala |
| COSMOSDB_ENGINE | cosmos |
| COUCHBASE_ENGINE | couchbase |
| DB2_ENGINE | db2 |
| DATABRICKS_ENGINE | databricks |
| DUCKDB_ENGINE | duckdb |
| EXASOL_ENGINE | exasol |
| FIREBIRD_ENGINE | firebirdsql |
| GENJI_ENGINE | genji |
| BIGQUERY_ENGINE | bigquery |
| SPANNER_ENGINE | spanner |
| ADODB_ENGINE | adodb |
| MAXCOMPUTE_ENGINE | maxcompute |
| MYSQL_ENGINE_ALT1 | mysql |
| MYSQL_ENGINE_ALT2 | mysql |
| ODBC_ENGINE | odbc |
| QL_ENGINE | ql |
| ASE_ENGINE | ase |
| HANA_ENGINE | hdb |
| RESTSQL_ENGINE | restsql |
| SNOWFLAKE_ENGINE | snowflake |
| TRINO_ENGINE | trino |
| VERTICA_ENGINE | vertica |
| VITESS_ENGINE | vitess |
| YDB_ENGINE | ydb |
| YQL_ENGINE | yql |

### Setting up TesoQL with Mongo
As it stated on the figure below tesoql looks for a client or a connection config to establish a connection with Mongo.

![](https://raw.githubusercontent.com/tesodev-com/tesoql/main/documentation/mongo.jpeg)

Connection logic of TesoQL to Mongodb

#### Usage Examples With Mongodb
###### Usage Example 1 : 
```go
tesoqlConfig := &tesoql.Config{
   Engine:    tesoql.MONGO_ENGINE,
   FieldsMap: &fieldsMap, //should be pre-defined by the user
   ConnectionConfig: &tesoql.ConnectionConfig{
      DBName:    "OMS",
      TableName: "OrderCollection",
      Client: *mongo.Client,
   },
   Pagination:    &tesoql.PaginationConfig{LimitUpperBound: 50},
   Toggles:       nil,
   PrintSqlQuery: false,
}


tesoQL := tesoqlConfig.NewTesoQL()
```
###### Usage Example 2 : 
```go
tesoqlConfig := &tesoql.Config{
   Engine:    tesoql.MONGO_ENGINE,
   FieldsMap: &fieldsMap,//should be pre-defined by the user
   ConnectionConfig: &tesoql.ConnectionConfig{
      DBName:    "OMS",
      TableName: "OrderCollection",
      ConnectionString: "mongodb://localhost:27017/",
   },
   Pagination:    &tesoql.PaginationConfig{LimitUpperBound: 50},
   Toggles:       nil,
   PrintSqlQuery: false,
}


tesoQL := tesoqlConfig.NewTesoQL()
```
###### Usage Example 3 (with DefaultConfig) :
```go
tesoqlConfig := tesoql.DefaultConfig
tesoqlConfig.FieldsMap = &fieldsMap //should be pre-defined by the user
tesoqlConfig.ConnectionConfig.ConnectionString = "mongodb://localhost:27017/"
tesoqlConfig.ConnectionConfig.DBName = "OMS"
tesoqlConfig.ConnectionConfig.TableName = "OrderCollection"
tesoQL := tesoqlConfig.NewTesoQL()
```
### Setting Up TesoQL with SQL Based Databases
![](https://raw.githubusercontent.com/tesodev-com/tesoql/main/documentation/sql.jpeg)

Connection logic of TesoQL to SQL based databases

#### Usage Examples With SQL Based Databases (Sqlite3)
###### Usage Example 1 : 

```go
tesoqlConfig := &tesoql.Config{
   Engine:    tesoql.SQLITE_ENGINE,
   FieldsMap: &fieldsMap, //should be pre-defined by the user
   ConnectionConfig: &tesoql.ConnectionConfig{
      DBName:    "OMS",
      TableName: "OrderTable",
      ConnectionString: "./myProject.sqlite",
   },
   Pagination:    &tesoql.PaginationConfig{LimitUpperBound: 50},
   Toggles:       nil,
   PrintSqlQuery: true,
}

tesoQL := tesoqlConfig.NewTesoQL()
```
###### Usage Example 2 : 
```go
db, err := sql.Open("sqlite3", "./myProject.sqlite")
if err!=nil{
   panic(err)
}


tesoqlConfig := &tesoql.Config{
   Engine:    tesoql.SQLITE_ENGINE,
   FieldsMap: &fieldsMap, //should be pre-defined by the user
   ConnectionConfig: &tesoql.ConnectionConfig{
      DBName:    "OMS",
      TableName: "OrderTable",
      Client:    db,
   },
   Pagination:    &tesoql.PaginationConfig{LimitUpperBound: 50},
   Toggles:       nil,
   PrintSqlQuery: true,
}

tesoQL := tesoqlConfig.NewTesoQL()
```

###### Usage Example 3 (with DefaultConfig) :

```go
tesoqlConfig := tesoql.DefaultConfig
tesoqlConfig.Engine = tesoql.SQLITE_ENGINE
tesoqlConfig.FieldsMap = &fieldsMap //should be pre-defined by the user
tesoqlConfig.ConnectionConfig.ConnectionString = "./myProject.sqlite"
tesoqlConfig.ConnectionConfig.DBName = "OMS"
tesoqlConfig.ConnectionConfig.TableName = "OrderTable"
tesoQL := tesoqlConfig.NewTesoQL()
```

------------

### FieldsMap
The *FieldsMap* is a ‘must’ component in *tesoql* that allows you to map the real field names in your database to more user-friendly or localized names -optionally-. This is particularly useful when you want to create a more intuitive interface for users who may not be familiar with the actual database field names.

For example, if you have a field in your database named ‘*orderStatus*’, you can map it to ‘*order_status*’ or any other alias you prefer, and use this alias in your queries. The *FieldsMap* ensures that these aliases are correctly translated back to the actual field names during query building and validation.

###### **Important:** 
*FieldsMap* declares a validation scheme for the `JsonMap.Validate()`. **For instance** let's consider a case like; a field in your database named ‘*orderStatus*’ mapped to ‘*order_status*’ in *FieldsMap* under *SearchFields*, but **NOT** under *ConditionFields*. An incoming querying parameter in the payload under “conditions” ; `JsonMap.Validate()` will be returning an error.

###### **Important:** 
With an incorrectly defined *FieldsMap*, *tesoql's* query builder will not function correctly as it will be looking for the incorrect field name. 

###### Usage Example : 
```go
var fieldsMap = &tesoql.FieldsMap{
   DateTimeFieldKeys: map[string]string{
      "order_date": "orderDate",
   },
   SearchFields: map[string]string{
      "productName": "product_name",
   },
   ProjectionFields: map[string]string{
      "id":              "id",
      "productName":     "product_name",
      "productId":       "product_id",
      "order_status":    "orderStatus",
      "orderId":         "order_id",
      "quantity":        "quantity",
      "amount":          "amount",
      "userCount":       "user_count",
      "remainingStock":  "remaining_stock",
      "order_date":      "orderDate",
   },
   SortingFields: map[string]string{
      "quantity": "quantity",
      "amount":   "amount",
   },
   ConditionFields: map[string]string{
      "remainingStock": "remaining_stock",
      "quantity":       "quantity",
      "order_date":     "orderDate",
      "userCount":      "user_count",
      //missing ‘order_status’
   },
}
tesoqlConfig := &tesoql.Config{
   Engine:    tesoql.MONGO_ENGINE,
   FieldsMap: fieldsMap, //should be pre-defined by the user
   ConnectionConfig: &tesoql.ConnectionConfig{
      DBName:    "OMS",
      TableName: "OrderCollection",
      Client: *mongo.Client,
   },
   Pagination:    &tesoql.PaginationConfig{LimitUpperBound: 50},
   Toggles:       nil,
   PrintSqlQuery: false,
}
tesoQL := tesoqlConfig.NewTesoQL()
payload := tesoql.JsonMap{
   Search: map[string][]interface{}{
      "order_status": {"pend", "Complet"},
   },
   ProjectionFields: []string{
      "order_status",
   },
   Conditions: map[string]tesoql.ConditionOperators{
      "order_status": {
         ValuesToExclude: []interface{}{"inactive"},
      },
   },
}


err := payload.Validate(tesoqlConfig)
```

The error (**ErrorResponseDTO*) will be resulted in 
```json
{
  "ErrorType": "TESOQL_VALIDATION_ERROR",
  "ErrorMsg": "Field : 'order_status' is not compatible to apply a condition query.",
  "ErrorCode": 400003
}
```
because ‘*order_status*’ was not declared in the *FieldsMap*, but it is still exist (incoming) in the payload (*JsonMap*).

On the other hand, the fields -that exist in the database- are mapped to their aliases. For instance a field named ‘*remaining_stock*’ is mapped to ‘*remainingStock*’  to query.

------------


### Building a JsonMap Payload

```go
payload := tesoql.JsonMap{
   Search: map[string][]interface{}{
      "productName": {"piz", "slic"},
   },
   ProjectionFields: []string{
      "id",
      "productName",
      "productId",
      "order_status",
      "orderId",
      "quantity",
      "amount",
      "userCount",
      "remainingStock",
      "order_date",
   },
   SortConditions: []tesoql.SortInput{
      {
         Field:         "quantity",
         SortCondition: "ASC",
      },
      {
         Field:         "amount",
         SortCondition: "DESC",
      },
   },
   Conditions: map[string]tesoql.ConditionOperators{
      "remainingStock": {
         GreaterThan: 25,
      },
      "order_status": {
         ValuesToExclude: []interface{}{"inactive"},
      },
      "quantity": {
         GreaterThan: 2,
         LowerThan:   10,
      },
      "order_date": {
         GreaterOrEqual: time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
         LowerOrEqual:   time.Now(),
      },
      "userCount": {
         ValuesToExactMatch: []interface{}{2},
      },
   },
   Pagination: tesoql.Pagination{
      Limit:  48,
      Offset: 2,
   },
   TotalCount:           true,
   SuppressDataResponse: false,
}

```
In this example, we’re constructing a *JsonMap* payload that defines the criteria for querying a database, including search filters, field projections, sorting rules, complex conditions, and pagination settings.
1. **Search:** The Search field is used to specify the criteria for filtering the data. In this case, we are searching for records where the *productName* field contains either "piz" or "slic". This allows the query to return records that match any of these product names.
2. **ProjectionFields:** The *ProjectionFields* slice lists the specific fields that should be included in the query result. Here, we are projecting the following fields: *id, productName, productId, order_status, orderId, quantity, amount, userCount, remainingStock*, and *order_date*. This means that only these fields will be returned in the results, which can help optimize performance by excluding unnecessary data.
3. **SortConditions**: The *SortConditions* field defines how the results should be ordered. In this example, we are sorting the results first by *quantity* in ascending order ("*ASC*"), and then by *amount* in descending order ("*DESC*"). This ensures that the data is organized in a specific and predictable manner.
4. **Conditions**: The *Conditions* field allows for more complex filtering of the data. Here are the conditions applied:
	5. *remainingStock*: We only want records where the remaining stock is greater than 25.
	6. *order_status*: We exclude records where the order status is "inactive".
	7. *quantity*: We are looking for records where the quantity is greater than 2 and less than 10.
	8. *order_date*: The order date must be between January 1, 2023, and the current date.
	9. *userCount*: We are specifically filtering for records where the user count is exactly 2.
10. **Pagination**: The *Pagination* field is used to manage how the results are paginated. In this case, we are limiting the results to 48 records (Limit: 48) and starting the results from the third record (Offset: 2). This is useful for controlling the amount of data returned and navigating through large datasets.
11. **TotalCount**: The *TotalCount* flag is set to true, meaning that the total count of matching records will be included in the response. This is often used for understanding the scope of the data that matches the query criteria.
12. **SuppressDataResponse**: The *SuppressDataResponse* flag is set to false, indicating that the actual data should be returned in the response. This flag is useful in scenarios where you may only want metadata without the full data payload.

------------

### Constructing Queries
#### SQL
**1. NewSqlQuery**

```go
sqlQuery := payload.NewSqlQuery(&fieldsMap)
```

This method constructs a type **SqlQuery* with given *JsonMap* payload and *FieldsMap*.

**2. GetSqlQuery**

```go
query, whereClause, args := payload.GetSqlQuery(&fieldsMap, "OrderTable", false)
```

This method calls `JsonMap.NewSqlQuery` inside, constructs the query and returns the query in the type of string with placeholders of ‘?’ to prevent SQL Injections. Returned arguments (args) is sorted respectively to the placeholders. 

------------

#### Mongodb

**1. NewMongoQuery**
```go
mongoQuery := payload.NewMongoQuery(&fieldsMap)
```
This method constructs a type **MongoQuery* with given *JsonMap* payload and *FieldsMap*.

- MongoQuery.Filter (*bson.D) 

Filter is constructed from JsonMap, and can be used directly in Mongodb Go Driver’s DB operation methods as a filter. For instance;
```go
totalCount, err := r.mongo.CountDocuments(ctx, mongoQuery.Filter)
```
- MongoQuery.Projection (*bson.D)

Projection is constructed from JsonMap, and can be used directly in Mongodb Go Driver’s options.FindOptions methods as a projection interface. For instance;
```go
opts = opts.SetProjection(mongoQuery.Projection)
cur, err := r.mongo.Find(ctx, filter, opts)
```
- MongoQuery.Sort (*bson.D)

Sort is constructed from JsonMap, and can be used directly in Mongodb Go Driver’s options.FindOptions methods as a sort interface. For instance;
```go
opts = opts.SetSort(mongoQuery.Sort)
cur, err := r.mongo.Find(ctx, filter, opts)
```
- MongoQuery.Limit & MongoQuery.Offset (int64)

Limit and Offset can be directly used in Mongodb Go Driver’s options.FindOptions methods as limit and skip variables.
```go
opts = options.Find().SetLimit(mongoQuery.Limit).SetSkip(mongoQuery.Offset)
cur, err := r.mongo.Find(ctx, filter, opts)
```

------------

### Handler Layer
You can use tesoql-echo package to implement handler-level integration. Relevant documentation is exists on [github.com/tesodev-com/tesoql-echo](https://github.com/tesodev-com/tesoql-echo "github.com/tesodev-com/tesoql-echo")
 

------------

### Using TesoQL as a Service

```go
tesoqlConfig := &tesoql.Config{
   Engine:    tesoql.MONGO_ENGINE,
   FieldsMap: fieldsMap, //should be pre-defined by the user
   ConnectionConfig: &tesoql.ConnectionConfig{
      DBName:    "OMS",
      TableName: "OrderCollection",
      Client: *mongo.Client,
   },
   Pagination:    &tesoql.PaginationConfig{LimitUpperBound: 50},
   Toggles:       nil,
   PrintSqlQuery: false,
}
tesoQL := tesoqlConfig.NewTesoQL()
//Payload variable had declared above examples (JsonMap)
results, totalCount, size, err := tesoQL.Service.Get(&payload)
```


*tesoQL.Service.Get* method returns; 
- **results** as a type of []map[string]any
- **totalCount** as int
- **size** as int
- **err** as type of **ErrorResponseDTO* which allows flexibility on use cases for any specific client that uses TesoQL.


------------

### Repository Layer

The repository layer provides an interface for database interactions. The implementation varies depending on the selected database engine.
- **MongoDB Implementation:** Provides methods tailored for MongoDB operations.
- **SQL Implementation:** Provides methods tailored for SQL database operations.

------------

### Types
##### 1. JsonMap
*JsonMap* represents the structure for defining query parameters. It includes search filters, projection fields, sorting conditions, complex conditions, pagination, and options to control the response behavior.

```go
type JsonMap struct {
   Search               map[string][]interface{}      `json:"search"`               
   ProjectionFields     []string                      `json:"projectionFields"` 
   SortConditions       []SortInput                   `json:"sortConditions"`
   Conditions           map[string]ConditionOperators `json:"conditions"`           
   Pagination           Pagination                    `json:"pagination"`           
   TotalCount           bool                          `json:"totalCount"`           
   SuppressDataResponse bool                          `json:"suppressDataResponse"`
}
```

###### Fields:
- **Search:** A map where the key is a field name and the value is a slice of interface{} representing the search values.
- **ProjectionFields:** A slice of strings that specifies which fields to return in the query result.
- **SortConditions:** A slice of SortInput structs that define the sorting rules for the query.
- **Conditions:** A map where the key is a field name and the value is a ConditionOperators struct, allowing for complex condition-based filtering.
- **Pagination:** A Pagination struct that defines how to paginate the results.
- **TotalCount:** A boolean flag that, if true, includes the total count of results in the response.
- **SuppressDataResponse:** A boolean flag that, if true, suppresses the data in the response (used in cases where only metadata is needed).

##### 2. SortInput
The SortInput struct is used within JsonMap to define sorting conditions for the query results.
```go
type SortInput struct {
   Field         string `json:"field"`    
   SortCondition string `json:"sortCondition"`
}
```
###### Fields:
- **Field:** The name of the field to sort by.
- **SortCondition:** The sorting order, typically "ASC" for ascending or "DESC" for descending.

##### 3. ConditionOperators
The ConditionOperators struct is used to define complex conditions for filtering query results. It supports multiple operators such as greater than, less than, exact match, etc.
```go
type ConditionOperators struct {
   GreaterThan        interface{}   `json:"greaterThan"`   
   GreaterOrEqual     interface{}   `json:"greaterOrEqual"`
   ValuesToExactMatch []interface{} `json:"valuesToExactMatch"`
   LowerThan          interface{}   `json:"lowerThan"`        
   LowerOrEqual       interface{}   `json:"lowerOrEqual"`  
   ValuesToExclude    []interface{} `json:"valuesToExclude"` 
}
```

###### Fields:
- **GreaterThan:** Used to filter results where the field is greater than the specified value.
- **GreaterOrEqual:** Used to filter results where the field is greater than or equal to the specified value.
- **LowerThan:** Used to filter results where the field is less than the specified value.
- **LowerOrEqual:** Used to filter results where the field is less than or equal to the specified value.
- **ValuesToMatch:** A slice of values to match exactly.
- **ValuesToExclude:** A slice of values to exclude from the results.
- **ValuesToExactMatch:** A slice of values for exact matching.

##### 4. Pagination
The Pagination struct is used to control the pagination of query results.
```go
type Pagination struct {
   Limit  int64 `json:"limit"` 
   Offset int64 `json:"offset"` 
}
```
###### Fields:
- **Limit:** The maximum number of items to return in the query results.
- **Offset:** The starting index from which to return results (useful for pagination).

These types are integral to the functionality of tesoql, providing a robust framework for constructing complex database queries in a structured and type-safe manner.

##### 5. ErrorResponseDTO
The ErrorResponseDTO struct is used to represent errors that occur during query processing in the tesoql package. It provides detailed information about the error, including the type, a descriptive message, and a specific error code.
```go
type ErrorResponseDTO struct {
   ErrorType string 
   ErrorMsg  string 
   ErrorCode int    
}
```

###### Fields:
- **ErrorType:** A string that categorizes the type of error, such as a validation error or a repository error.
- **ErrorMsg:** A detailed error message that explains what went wrong.
- **ErrorCode:** A numeric code that represents the specific error, which can be used for error handling or logging purposes.

This struct is crucial for providing clear and actionable feedback when something goes wrong during the processing of queries in tesoql. It helps developers quickly identify and respond to issues in their code.

##### 5.1 ErrorType List

| tesoql Error Code  |  string equivalent  |
| ------------ | ------------ |
| BINDING_ERR  |  "BINDING_ERROR" |
|  TESOQL_MONGO_ERROR |  "MONGO_ERROR" |
| TESOQL_SQL_ERROR  |  "TESOQL_SQL_ERROR" |
| TESOQL_TOGGLE_ERROR  | "TESOQL_TOGGLE_ERROR"  |
|  TESOQL_VALIDATION_ERROR | "TESOQL_VALIDATION_ERROR"  |

##### 5.2 ErrorCode List

###### 5.2.1 Validation Error Codes
| tesoql Error Code  |  integer equivalent |
| ------------ | ------------ |
| BINDING_ERR_CODE  |  400000 |
|SORTABLE_ERR_CODE| 400001  |
| SEARCHABLE_ERR_CODE  | 400002  |
| PROJECTION_ERR_CODE  |  400003 |
| CONDITION_ERR_CODE  |  400004 |

###### 5.2.2 Toggle Validation Error Codes

| tesoql Error Code  |  integer equivalent |
| ------------ | ------------ |
| SORTABLE_TOGGLE_ERR_CODE | 400005 |
| SEARCHABLE_TOGGLE_ERR_CODE | 400006 |
| PROJECTION_TOGGLE_ERR_CODE | 400007 |
| CONDITION_TOGGLE_ERR_CODE | 400008 |
| PAGINATION_TOGGLE_ERR_CODE | 400009 |
| GREATERTHAN_CONDITION_TOGGLE_ERR_CODE | 400010 |
| GREATEROREQUAL_CONDITION_TOGGLE_ERR_CODE | 400011 |
| LOWERTHAN_CONDITION_TOGGLE_ERR_CODE | 400012 |
| LOWEROREQUAL_CONDITION_TOGGLE_ERR_CODE | 400013 |
| VALUESTOEXCLUDE_CONDITION_TOGGLE_ERR_CODE | 400014 |
| VALUESTOEXACTMATCH_CONDITION_TOGGLE_ERR_CODE | 400015 |
| LOWTOHIGH_CONDITION_TOGGLE_ERR_CODE | 400016 |
| HIGHTOLOW_CONDITION_TOGGLE_ERR_CODE | 400017 |

###### 5.2.3 Repository Level Error Codes
| tesoql Error Code  |  integer equivalent |
| ------------ | ------------ |
| SQL_QUERYEXEC_ERR_CODE | 500001 |
| SQL_COLUMNS_ERR_CODE | 500002 |
| SQL_SCAN_ERR_CODE | 500003 |
| SQL_COUNT_QUERYEXEC_ERR_CODE | 500004 |
| MONGO_FIND_ERR_CODE | 500005 |
| MONGO_CURSOR_ERR_CODE | 500006 |

##### 6. MongoQuery
The *MongoQuery* struct represents a MongoDB query structure, including filter criteria, projection, sorting, limit, and offset options. It is used to construct queries that are specific to MongoDB databases.
```go
type MongoQuery struct {
   Filter     *bson.D 
   Projection *bson.D 
   Sort       *bson.D 
   Limit      int64   
   Offset     int64   
}
```

###### Fields:
- **Filter:** A BSON document that defines the criteria to filter the MongoDB documents.
- **Projection:** A BSON document that specifies the fields to include or exclude in the query result.
- **Sort:** A BSON document that defines the sorting order of the query results.
- **Limit:** The maximum number of documents to return.
- **Offset:** The number of documents to skip before starting to return the results.

##### 7. SqlQuery
The SqlQuery struct represents an SQL query structure, including clauses like SELECT, WHERE, ORDER BY, LIMIT, and OFFSET, as well as the arguments for query placeholders. It is used to build queries that are specific to SQL-based databases.

```go
type SqlQuery struct {
   Select  string        
   Where   string        
   OrderBy string       
   Limit   string        
   Offset  string        
   Args    []interface{} 
}
```

###### Fields:
- **Select:** A string representing the fields to be selected in the SQL query.
- **Where:** A string that defines the conditions for filtering the SQL query results.
- **OrderBy:** A string that specifies the sorting order for the SQL query results.
- **Limit:** A string that defines the maximum number of rows to return.
- **Offset:** A string that specifies the number of rows to skip before starting to return the results.
- **Args:** A slice of interfaces that holds the arguments for the query's placeholders (e.g., values for ? placeholders in the SQL query).

These structs, *MongoQuery* and *SqlQuery*, are essential for building database-specific queries in *tesoql*, providing flexibility and control over how data is queried and retrieved from MongoDB and SQL databases.

------------


## Contributing
Contributions are welcome! Please open an issue or submit a pull request if you have any improvements or bug fixes.
