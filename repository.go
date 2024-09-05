package tesoql

type iTesoQlRepo interface {
	repository(jsonMap *JsonMap) ([]map[string]interface{}, int, int, *ErrorResponseDTO)
}
