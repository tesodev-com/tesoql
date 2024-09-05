package tesoql

func validateToggles(jsonMap *JsonMap, toggleConfig *ToggleConfig) *ErrorResponseDTO {
	if toggleConfig != nil {
		err := validateUpperToggles(jsonMap, toggleConfig)
		if err != nil {
			return err
		}
		err = validateConditionToggles(jsonMap, toggleConfig)
		if err != nil {
			return err
		}
		err = validateSortingToggles(jsonMap, toggleConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

func validateUpperToggles(jsonMap *JsonMap, t *ToggleConfig) *ErrorResponseDTO {
	if t.DisableTotalCount {
		jsonMap.TotalCount = false
	}

	if t.DisableSearch && len(jsonMap.Search) > 0 {
		return newResponse(TESOQL_TOGGLE_ERROR, "DisableSearch toggle is open.", SEARCHABLE_TOGGLE_ERR_CODE)
	}

	if t.DisableProjection && len(jsonMap.ProjectionFields) > 0 {
		return newResponse(TESOQL_TOGGLE_ERROR, "DisableProjection toggle is open.", PROJECTION_TOGGLE_ERR_CODE)
	}

	if t.DisableSorting && len(jsonMap.SortConditions) > 0 {
		return newResponse(TESOQL_TOGGLE_ERROR, "DisableSorting toggle is open.", SORTABLE_TOGGLE_ERR_CODE)
	}

	if t.DisablePagination && (jsonMap.Pagination.Limit > 0 || jsonMap.Pagination.Offset > 0) {
		return newResponse(TESOQL_TOGGLE_ERROR, "DisablePagination toggle is open.", PAGINATION_TOGGLE_ERR_CODE)
	}

	if t.DisableConditioning && len(jsonMap.Conditions) > 0 {
		return newResponse(TESOQL_TOGGLE_ERROR, "DisableConditioning toggle is open.", CONDITION_TOGGLE_ERR_CODE)
	}

	return nil
}

func validateConditionToggles(jsonMap *JsonMap, toggleConfig *ToggleConfig) *ErrorResponseDTO {
	for _, ops := range jsonMap.Conditions {
		if toggleConfig.ConditioningToggles != nil {
			if toggleConfig != nil && toggleConfig.ConditioningToggles != nil {
				if toggleConfig.ConditioningToggles.DisableGreaterThan && ops.GreaterThan != nil {
					return newResponse(TESOQL_TOGGLE_ERROR, "DisableGreaterThan toggle is open.", GREATERTHAN_CONDITION_TOGGLE_ERR_CODE)
				}
				if toggleConfig.ConditioningToggles.DisableGreaterOrEqual && ops.GreaterOrEqual != nil {
					return newResponse(TESOQL_TOGGLE_ERROR, "DisableGreaterOrEqual toggle is open.", GREATEROREQUAL_CONDITION_TOGGLE_ERR_CODE)
				}
				if toggleConfig.ConditioningToggles.DisableLowerThan && ops.LowerThan != nil {
					return newResponse(TESOQL_TOGGLE_ERROR, "DisableLowerThan toggle is open.", LOWERTHAN_CONDITION_TOGGLE_ERR_CODE)
				}
				if toggleConfig.ConditioningToggles.DisableLowerOrEqual && ops.LowerOrEqual != nil {
					return newResponse(TESOQL_TOGGLE_ERROR, "DisableLowerOrEqual toggle is open.", LOWEROREQUAL_CONDITION_TOGGLE_ERR_CODE)
				}
				if toggleConfig.ConditioningToggles.DisableValuesToExclude && ops.ValuesToExclude != nil {
					return newResponse(TESOQL_TOGGLE_ERROR, "DisableValuesToExclude toggle is open.", VALUESTOEXCLUDE_CONDITION_TOGGLE_ERR_CODE)
				}
				if toggleConfig.ConditioningToggles.DisableValuesToExactMatch && ops.ValuesToExactMatch != nil {
					return newResponse(TESOQL_TOGGLE_ERROR, "DisableValuesToExactMatch toggle is open.", VALUESTOEXACTMATCH_CONDITION_TOGGLE_ERR_CODE)

				}
			}
		}
	}
	return nil
}

func validateSortingToggles(jm *JsonMap, t *ToggleConfig) *ErrorResponseDTO {
	for _, sortInput := range jm.SortConditions {
		if sortInput.SortCondition == "ASC" && t.SortingToggles != nil && t.SortingToggles.DisableLowToHigh {
			return newResponse(TESOQL_TOGGLE_ERROR, "DisableLowToHigh toggle is open.", LOWTOHIGH_CONDITION_TOGGLE_ERR_CODE)
		}
		if sortInput.SortCondition == "DESC" && t.SortingToggles != nil && t.SortingToggles.DisableHighToLow {
			return newResponse(TESOQL_TOGGLE_ERROR, "DisableHighToLow toggle is open.", HIGHTOLOW_CONDITION_TOGGLE_ERR_CODE)
		}
	}
	return nil
}
