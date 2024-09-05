package tesoql

// Service provides the core functionality for interacting with the repository.
// It is initialized with a repository interface and a set of feature toggles.
type Service struct {
	repo    *iTesoQlRepo  // The repository interface for database interactions.
	toggles *ToggleConfig // Feature toggles that control the behavior of the service.
}

func newTesoQlService(repo *iTesoQlRepo, toggles *ToggleConfig) *Service {
	return &Service{repo: repo, toggles: toggles}
}

// Get retrieves data from the repository based on the provided JsonMap.
// It performs validation against the service's toggles before querying the repository.
//
// Example usage:
//
//	tesoqlConfig := tesoql.Config{
//		 Configuration setup...
//	}
//	tesoQL := tesoqlConfig.NewTesoQL()
//
//	jsonMapVariable := tesoql.JsonMap{
//		 Payload...
//	}
//
//	//It is possible to use Validate function (Assuming fieldsmap and paginationconfig are previously defined)
//	jsonMapVariable.Validate(fieldsmap, paginationconfig)
//
//	results, totalCount, size, err := tesoQL.Service.Get(&jsonMapVariable)
//
// Returns:
//
// - []map[string]any: The data retrieved from the repository.
//
// - int: The total count of records that match the query.
//
// - int: The size of the current page of results.
//
// - *ErrorResponseDTO: An error response, if any occurred during validation or retrieval.
func (s *Service) Get(jsonMap *JsonMap) ([]map[string]any, int, int, *ErrorResponseDTO) {
	validationErr := validateToggles(jsonMap, s.toggles)
	if validationErr != nil {
		return nil, 0, 0, validationErr
	}
	r := *s.repo
	response, totalCount, size, err := r.repository(jsonMap)
	if err != nil {
		return nil, 0, 0, err
	}
	return response, totalCount, size, nil
}
