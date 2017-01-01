package utils

// Response: 返回响应
func Response(status int, data interface{}, err string) (result map[string]interface{}) {

	result = map[string]interface{}{
		"status": status,
		"data":   data,
		"err":    err,
	}

	return

}
