package tesoql

import (
	"time"
)

func treatDateTime(k string, dateFields map[string]string, val interface{}) interface{} {
	if isDateTimeFieldKey(k, dateFields) {
		val = parseDateTimeOrReturn(val)
	}
	return val
}

func parseDateTimeOrReturn(val interface{}) interface{} {
	strVal, ok := (val).(string)
	if ok {
		dateFilter, err := time.Parse(time.RFC3339, strVal)
		if err == nil {
			parsedVal := interface{}(dateFilter)
			return &parsedVal
		}
	}
	return val
}

func isDateTimeFieldKey(k string, dateFieldKeys map[string]string) bool {
	if _, exists := dateFieldKeys[k]; exists {
		return true
	}
	return false
}

func newResponse(errType string, msg string, errCode int) *ErrorResponseDTO {
	return &ErrorResponseDTO{
		ErrorType: errType,
		ErrorMsg:  msg,
		ErrorCode: errCode,
	}
}
