package tesoql

// TesoQL struct encapsulates the service layer for TesoQL.
// It holds a reference to the Service, which is responsible for
// handling the core querying operations.
type TesoQL struct {
	Service *Service
}

// NewTesoQL initializes a new instance of TesoQL based on the provided configuration.
// It determines the appropriate repository implementation (e.g., MongoDB, SQL)
// based on the engine specified in the Config struct.
// If the specified engine is not supported, the function panics.
//
// The method sets up the repository, creates a new TesoQL service with the
// repository and feature toggles, and returns a pointer to the TesoQL struct.
//
// Example usage:
//
//	cfg := &tesoql.Config{
//		Engine: tesoql.MONGO_ENGINE,
//		// Other config fields...
//	}
//	tesoql := cfg.NewTesoQL()
//
// Example usage with default config:
//
//	cfg := tesoql.DefaultConfig
//	cfg.Engine = tesoql.POSTGRES_ENGINE
//	cfg.ConnectionConfig = &tesoql.ConnectionConfig{
//		DBName:           "example_db",
//		ConnectionString: "postgres://user:password@localhost:5432/example_db",
//	}
//	tesoql := cfg.NewTesoQL()
//
// Returns:
//
// - *TesoQL: A pointer to the initialized TesoQL struct.
func (cfg *Config) NewTesoQL() *TesoQL {
	var repo iTesoQlRepo

	switch cfg.Engine {
	case "mongo":
		repo = newMongoRepository(cfg)
	case sqlDriverList[cfg.Engine]:
		repo = newSqlRepository(cfg)
	default:
		panic("DB connection could not have been established!")
	}

	service := newTesoQlService(&repo, cfg.Toggles)
	return &TesoQL{Service: service}
}
