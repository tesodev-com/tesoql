package tesoql

import (
	"fmt"
)

// Validate performs a series of checks on the JsonMap instance to ensure
// that the search, projection, sorting, and pagination settings are valid
// according to the provided FieldsMap and PaginationConfig.
//
// It validates search fields, projection fields, sorting conditions, and
// adjusts pagination settings. If any validation fails, it returns an
// ErrorResponseDTO with the relevant error information.
//
// Example usage:
//
//	jsonMap := tesoql.JsonMap{ /* JSON input */ }
//	err := jsonMap.Validate(fieldsMap, paginationConfig)
//	if err != nil {
//		// Handle validation error
//	}
//
// Returns:
//
// - *ErrorResponseDTO: An error response if validation fails, or nil if all validations pass.
func (jm *JsonMap) Validate(cfg *Config) *ErrorResponseDTO {

	if cfg.FieldsMap != nil {
		err := jm.validateSearchAndProjection(cfg.FieldsMap)
		if err != nil {
			return err
		}
	}

	err := jm.validateConditions(cfg.FieldsMap)
	if err != nil {
		return err
	}

	err = jm.validateSorting(cfg.FieldsMap)
	if err != nil {
		return err
	}

	jm.validatePagination(cfg.Pagination)

	return nil
}

// validateSorting checks if the sorting fields specified in the JsonMap
// exist in the FieldsMap and ensures the sorting condition is either "ASC" or "DESC".
//
// Returns:
//
// - *ErrorResponseDTO: An error response if validation fails, or nil if validation passes.
func (jm *JsonMap) validateSorting(fm *FieldsMap) *ErrorResponseDTO {
	for _, sortInput := range jm.SortConditions {
		if _, exists := fm.SortingFields[sortInput.Field]; !exists {
			return newResponse(
				TESOQL_VALIDATION_ERROR,
				fmt.Sprintf("Field : '%v' is not sortable.", sortInput.Field),
				SORTABLE_ERR_CODE)
		}
		if sortInput.SortCondition != "ASC" && sortInput.SortCondition != "DESC" {
			return newResponse(
				TESOQL_VALIDATION_ERROR,
				"Sort condition operators cannot be different than 'ASC' or 'DESC' (type : string)!",
				SORTABLE_ERR_CODE)
		}

	}
	return nil
}

// validatePagination adjusts the pagination settings in the JsonMap based on the provided
// PaginationConfig. It ensures that the limit and offset are within acceptable bounds.
func (jm *JsonMap) validatePagination(paginationCfg *PaginationConfig) {
	o := jm.Pagination.Offset
	if o < 0 || o > 500 {
		o = 0
	}

	upperBound := int64(50)
	if paginationCfg != nil {
		upperBound = paginationCfg.LimitUpperBound
	}
	l := jm.Pagination.Limit
	if l <= 0 || l > upperBound {
		l = upperBound
	}

	jm.Pagination.Limit = l
	jm.Pagination.Offset = o

}

// validateSearchAndProjection checks if the search and projection fields specified
// in the JsonMap exist in the FieldsMap.
//
// Returns:
//
// - *ErrorResponseDTO: An error response if validation fails, or nil if validation passes.
func (jm *JsonMap) validateSearchAndProjection(fm *FieldsMap) *ErrorResponseDTO {
	if fm.SearchFields != nil && jm.Search != nil {
		for field := range jm.Search {
			if _, exists := fm.SearchFields[field]; !exists {
				return newResponse(
					TESOQL_VALIDATION_ERROR,
					fmt.Sprintf("Field : '%v' is not searchable.", field),
					SEARCHABLE_ERR_CODE)
			}
		}
	}

	if fm.ProjectionFields != nil && jm.ProjectionFields != nil {
		for _, field := range jm.ProjectionFields {
			if _, exists := fm.ProjectionFields[field]; !exists {
				return newResponse(
					TESOQL_VALIDATION_ERROR,
					fmt.Sprintf("Field : '%v' is not compatible for to apply projectioning.", field),
					PROJECTION_ERR_CODE)
			}
		}
	}
	return nil
}

// validateConditions checks if the condition fields specified in the JsonMap
// exist in the FieldsMap and are compatible for applying condition queries.
//
// Returns:
//
// - *ErrorResponseDTO: An error response if validation fails, or nil if validation passes.
func (jm *JsonMap) validateConditions(fm *FieldsMap) *ErrorResponseDTO {
	for field := range jm.Conditions {
		if fm != nil && fm.ConditionFields != nil {
			if _, exists := fm.ConditionFields[field]; !exists {
				return newResponse(
					TESOQL_VALIDATION_ERROR,
					fmt.Sprintf("Field : '%v' is not compatible to apply a condition query.", field),
					CONDITION_ERR_CODE)
			}

		}
	}
	return nil
}
