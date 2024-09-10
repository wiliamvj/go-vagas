package utils

import "fmt"

func ConvertMap(input map[interface{}]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range input {
		keyStr, ok := k.(string)
		if !ok {
			keyStr = fmt.Sprintf("%v", k)
		}
		result[keyStr] = v
	}
	return result
}
