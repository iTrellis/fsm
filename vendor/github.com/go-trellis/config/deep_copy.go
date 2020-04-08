// GNU GPL v3 License
// Copyright (c) 2017 github.com:go-trellis

package config

// DeepCopy 深度拷贝
func DeepCopy(value interface{}) interface{} {
	switch valueType := value.(type) {
	case map[string]interface{}, Options:
		newMap := make(map[string]interface{})
		valueMap := valueType.(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}
		return newMap
	case map[interface{}]interface{}:
		newMap := make(map[interface{}]interface{})
		for k, v := range valueType {
			newMap[k] = DeepCopy(v)
		}
		return newMap
	case []interface{}:
		newSlice := make([]interface{}, len(valueType))
		for k, v := range valueType {
			newSlice[k] = DeepCopy(v)
		}
		return newSlice
	}

	return value
}
